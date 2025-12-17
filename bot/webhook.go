package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/purkhanov/gogram/api"
	"github.com/purkhanov/gogram/utils"
)

const (
	setWebhookUrl     = "/setWebhook"
	deleteWebhookUrl  = "/deleteWebhook"
	getWebhookInfoUrl = "/getWebhookInfo"
)

// type webhookResponse[T webhookInfo | string | bool] struct {
// 	Ok          bool   `json:"ok"`
// 	ErrorCode   int    `json:"error_code"`
// 	Description string `json:"description"`
// 	Result      T      `json:"result"`
// }

type webhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`

	// True, if a custom certificate was provided
	// for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// Optional. Currently used webhook IP address
	IpAddress string `json:"ip_address,omitempty"`

	// Optional. Unix time for the most recent error that
	// happened when trying to deliver an update via webhook
	LastErrorDate int `json:"last_error_date,omitempty"`

	// Optional. Error message in human-readable format
	// for the most recent error that happened when
	// trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// Optional. Unix time of the most recent error that
	// happened when trying to synchronize available
	// updates with Telegram datacenters
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`

	// Optional. The maximum allowed number of simultaneous HTTPS
	// connections to the webhook for update delivery
	MaxConnections int `json:"max_connections,omitempty"`

	// Optional. A list of update types the bot is subscribed to.
	// Defaults to all update types except chat_member
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type WebhookOptions struct {
	// HTTPS URL to send updates to. Use an empty
	// string to remove webhook integration
	URL string `json:"url" validate:"required"`

	// Upload your public key certificate so that the
	// root certificate in use can be checked. See our
	// self-signed guide for details.
	Certificate string `json:"certificate,omitempty"`

	// The fixed IP address which will be used to send webhook
	// requests instead of the IP address resolved through DNS
	IPAddress string `json:"ip_address,omitempty"`

	// The maximum allowed number of simultaneous HTTPS connections
	// to the webhook for update delivery, 1-100. Defaults to 40.
	// Use lower values to limit the load on your bot's server, and
	// higher values to increase your bot's throughput.
	MaxConnections uint8 `json:"max_connections,omitempty" validate:"min=1,max=100"`

	// A JSON-serialized list of the update types you want your
	// bot to receive. For example, specify ["message",
	// "edited_channel_post", "callback_query"] to only receive updates
	// of these types. See Update for a complete list of available update
	// types. Specify an empty list to receive all update types except
	// chat_member, message_reaction, and message_reaction_count (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created
	// before the call to the setWebhook, so unwanted updates may be
	// received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`

	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token”
	// in every webhook request, 1-256 characters. Only characters A-Z, a-z,
	// 0-9, _ and - are allowed. The header is useful to ensure that the
	// request comes from a webhook set by you.
	SecretToken string `json:"secret_token,omitempty" validate:"min=1,max=256"`
}

// Use this method to specify a URL and receive incoming updates via an
// outgoing webhook. Whenever there is an update for the bot, we will send
// an HTTPS POST request to the specified URL, containing a JSON-serialized
// Update. In case of an unsuccessful request (a request with response HTTP
// status code different from 2XY), we will repeat the request and give up
// after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the webhook was set by you, you can
// specify secret data in the parameter secret_token. If specified, the
// request will contain a header “X-Telegram-Bot-Api-Secret-Token” with
// the secret token as content.
func (b *Bot) SetWebhook(options WebhookOptions) (string, error) {
	if err := utils.ValidateStruct(options); err != nil {
		return "", fmt.Errorf("invalid webhook parameters: %w", err)
	}

	if !strings.HasPrefix(options.URL, "https://") {
		return "", errors.New("webhook URL must use HTTPS")
	}

	fullUrl := b.urlWithToken + setWebhookUrl

	c, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	var resp []byte
	var err error

	if options.Certificate != "" {
		resp, err = b.setWebhookWithCertificate(c, fullUrl, options)
	} else {
		resp, err = b.setWebhookWithoutCertificate(c, fullUrl, options)
	}

	if err != nil {
		return "", fmt.Errorf("failed to set webhook: %w", err)
	}

	return b.parseWebhookResponse(resp)
}

func (b *Bot) setWebhookWithoutCertificate(
	ctx context.Context, fullURL string, params WebhookOptions,
) ([]byte, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal parameters: %w", err)
	}

	return b.api.DoRequestWithContextAndData(
		ctx, http.MethodPost, fullURL, data,
	)
}

func (b *Bot) setWebhookWithCertificate(
	c context.Context, fullURL string, params WebhookOptions,
) ([]byte, error) {
	// Validate certificate file exists and is readable
	fileInfo, err := os.Stat(params.Certificate)
	if err != nil {
		return nil, fmt.Errorf("certificate file error: %w", err)
	}

	// Check file size limit (e.g., 5MB)
	const maxCertificateSize = 5 << 20 // 5MB

	if fileInfo.Size() > maxCertificateSize {
		return nil, fmt.Errorf(
			"certificate file too large: %d bytes (max: %d)",
			fileInfo.Size(), maxCertificateSize,
		)
	}

	file, err := os.Open(params.Certificate)
	if err != nil {
		return nil, fmt.Errorf("failed to open certificate file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writeField := func(key, value string) error {
		return writer.WriteField(key, value)
	}

	if err := writeField("url", params.URL); err != nil {
		return nil, fmt.Errorf("failed to write URL field: %w", err)
	}

	if params.IPAddress != "" {
		err := writeField("url", params.IPAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to write URL field: %w", err)
		}
	}

	if params.MaxConnections != 0 {
		err := writeField("max_connections", fmt.Sprintf("%d", params.MaxConnections))
		if err != nil {
			return nil, fmt.Errorf("failed to write max_connections field: %w", err)
		}
	}

	if len(params.AllowedUpdates) > 0 {
		allowedUpdatesJSON, err := json.Marshal(params.AllowedUpdates)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal allowed_updates: %w", err)
		}
		err = writeField("allowed_updates", string(allowedUpdatesJSON))
		if err != nil {
			return nil, fmt.Errorf("failed to write allowed_updates field: %w", err)
		}
	}

	err = writeField("drop_pending_updates", fmt.Sprintf("%t", params.DropPendingUpdates))
	if err != nil {
		return nil, fmt.Errorf("failed to write drop_pending_updates field: %w", err)
	}

	if params.SecretToken != "" {
		err := writeField("secret_token", params.SecretToken)
		if err != nil {
			return nil, fmt.Errorf("failed to write secret_token field: %w", err)
		}
	}

	// Add certificate file
	part, err := writer.CreateFormFile("certificate", filepath.Base(params.Certificate))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to copy certificate file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(c, http.MethodPost, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return b.api.DoRequest(req)
}

func (b *Bot) parseWebhookResponse(resp []byte) (string, error) {
	var result api.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if !result.Ok {
		return "", fmt.Errorf("telegram API error: %s", result.Description)
	}

	return result.Description, nil
}

// Pass True to drop all pending updates
func (b *Bot) DeleteWebhook(dropPendingUpdates bool) (string, error) {
	fullURL := b.urlWithToken + deleteWebhookUrl
	param := map[string]bool{"drop_pending_updates": dropPendingUpdates}

	data, err := json.Marshal(param)
	if err != nil {
		return "", fmt.Errorf("cannot to marshal: %w", err)
	}

	ctx, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(ctx, http.MethodPost, fullURL, data)

	if err != nil {
		return "", fmt.Errorf("failed to delete webhook: %w", err)
	}

	return b.parseWebhookResponse(resp)
}

// func (b *Bot) GetWebhookInfo() (webhookInfo, error) {
// 	c, cancel := context.WithTimeout(b.Ctx, httpRequestTimeout)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(
// 		c, http.MethodGet, b.urlWithToken+getWebhookInfoUrl, nil,
// 	)
// 	if err != nil {
// 		return webhookInfo{}, fmt.Errorf("failed to create request: %w", err)
// 	}

// 	resp, err := b.api.DoRequest(req)
// 	if err != nil {
// 		return webhookInfo{}, err
// 	}

// 	var result types.APIResponse[webhookInfo]

// 	if err := json.Unmarshal(resp, &result); err != nil {
// 		return webhookInfo{}, fmt.Errorf("failed to parse response: %w", err)
// 	}

// 	if !result.Ok {
// 		return webhookInfo{}, fmt.Errorf(
// 			"telegram API error: code %d - %s",
// 			result.ErrorCode, result.Description,
// 		)
// 	}

// 	return result.Result, nil
// }
