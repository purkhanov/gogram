package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

const (
	setWebhookUrl    = "/setWebhook"
	deleteWebhookUrl = "/deleteWebhook"
)

type webhookResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type SetWebhookParameters struct {
	// HTTPS URL to send updates to. Use an empty
	// string to remove webhook integration
	URL string

	// Upload your public key certificate so that the
	// root certificate in use can be checked. See our
	// self-signed guide for details.
	Certificate string

	// The fixed IP address which will be used to send webhook
	// requests instead of the IP address resolved through DNS
	IPAddress string

	// The maximum allowed number of simultaneous HTTPS connections
	// to the webhook for update delivery, 1-100. Defaults to 40.
	// Use lower values to limit the load on your bot's server, and
	// higher values to increase your bot's throughput.
	MaxConnections uint8

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
	AllowedUpdates []string

	// Pass True to drop all pending updates
	DropPendingUpdates bool

	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token”
	// in every webhook request, 1-256 characters. Only characters A-Z, a-z,
	// 0-9, _ and - are allowed. The header is useful to ensure that the
	// request comes from a webhook set by you.
	SecretToken string

	formData url.Values
}

func (p *SetWebhookParameters) validateAndGetParams() error {
	p.formData = url.Values{}

	if p.URL == "" {
		return fmt.Errorf("URL is required")
	}
	p.formData.Add("url", p.URL)

	if p.IPAddress != "" {
		if net.ParseIP(p.IPAddress) == nil {
			return errors.New("invalid IP address")
		}
		p.formData.Add("ip_address", p.IPAddress)
	}

	if p.MaxConnections != 0 {
		if p.MaxConnections > 100 {
			return fmt.Errorf("max connections must be 1-100got %d", p.MaxConnections)
		}
		p.formData.Add("max_connections", fmt.Sprint(p.MaxConnections))
	}

	if len(p.AllowedUpdates) > 0 {
		updateJSON, err := json.Marshal(p.AllowedUpdates)
		if err != nil {
			return fmt.Errorf("failed to marshal allowed_updates: %w", err)
		}
		p.formData.Add("allowed_updates", string(updateJSON))
	}

	if p.DropPendingUpdates {
		p.formData.Add("drop_pending_updates", "true")
	}

	if p.SecretToken != "" {
		if len(p.SecretToken) > 256 {
			return fmt.Errorf("secret token must be 1-256 characters, got %d", len(p.SecretToken))
		}

		// Validate allowed characters: A-Z, a-z, 0-9, _, -
		validPattern := regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
		if !validPattern.MatchString(p.SecretToken) {
			return fmt.Errorf("secret token contains invalid characters. Only A-Z, a-z, 0-9, _, - are allowed")
		}

		p.formData.Add("secret_token", p.SecretToken)
	}

	return nil
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
func (b *Bot) SetWebhook(params SetWebhookParameters) (string, error) {
	if err := params.validateAndGetParams(); err != nil {
		return "", err
	}

	fullUrl := b.urlWithToken + setWebhookUrl

	var resp *http.Response
	var err error

	if params.Certificate != "" {
		resp, err = b.setWebhookWithCertificate(fullUrl, params)
	} else {
		resp, err = http.PostForm(fullUrl, params.formData)
	}

	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("telegram API returned non-200 status: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result webhookResponse

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if !result.Ok {
		return "", fmt.Errorf("telegram API error: %s", result.Description)
	}

	return result.Description, nil
}

func (b *Bot) setWebhookWithCertificate(fullUrl string, params SetWebhookParameters) (*http.Response, error) {
	file, err := os.Open(params.Certificate)
	if err != nil {
		return nil, fmt.Errorf("failed to open certificate file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add other parameters
	for key, values := range params.formData {
		for _, vavalue := range values {
			writer.WriteField(key, vavalue)
		}
	}

	// Add certificate file
	part, err := writer.CreateFormFile("certificate", filepath.Base(params.Certificate))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	io.Copy(part, file)
	writer.Close()

	return http.Post(fullUrl, writer.FormDataContentType(), body)
}

// Pass True to drop all pending updates
func (b *Bot) DeleteWebhook(dropPendingUpdates bool) (string, error) {
	req, err := http.NewRequest(http.MethodGet, b.urlWithToken+deleteWebhookUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", contentTypeJSON)

	if dropPendingUpdates {
		params := url.Values{}
		params.Add("drop_pending_updates", "True")
		req.URL.RawQuery = params.Encode()
	}

	resp, err := b.api.DoRequest(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}

	result, ok := resp.Result.(webhookResponse)
	if !ok {
		return "", fmt.Errorf("invalid webhook response type: %v", result)
	}

	return result.Description, nil
}
