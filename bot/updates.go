package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/purkhanov/gogram/types"
)

const (
	getUpdatesUrl  = "/getUpdates"
	updatesTimeout = "3600"
)

// Use this method to receive incoming updates using long polling.
// Returns an Array of Update objects.
type GetUpdateParams struct {
	// Identifier of the first update to be returned. Must be greater
	// by one than the highest among the identifiers of previously
	// received updates. By default, updates starting with the earliest
	// unconfirmed update are returned. An update is considered confirmed
	// as soon as getUpdates is called with an offset higher than its
	// update_id. The negative offset can be specified to retrieve updates
	// starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	Offset int `json:"offset,omitempty"`

	// Limits the number of updates to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit uint `json:"limit,omitempty"`

	// Timeout in seconds for long polling. Defaults to 0, i.e.
	// usual short polling. Should be positive, short polling
	// should be used for testing purposes only.
	Timeout uint `json:"timeout,omitempty"`

	// A JSON-serialized list of the update types you want your bot
	// to receive. For example, specify ["message", "edited_channel_post",
	// "callback_query"] to only receive updates of these types. See
	// Update for a complete list of available update types. Specify an
	// empty list to receive all update types except chat_member,
	// message_reaction, and message_reaction_count (default). If not
	// specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates
	// created before the call to getUpdates, so unwanted updates
	// may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (u *GetUpdateParams) validate() error {
	if u.Limit != 0 && u.Limit > 100 {
		return fmt.Errorf("limit must be between 1 and 100")
	}

	return nil
}

func (b *Bot) GetUpdates(params GetUpdateParams) ([]types.Update, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	formData := url.Values{}
	formData.Add("offset", fmt.Sprintf("%d", params.Offset))

	if params.Limit != 0 {
		formData.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	timeout := fmt.Sprintf("%d", params.Timeout)
	if params.Timeout != 0 {
		timeout = fmt.Sprintf("%d", params.Timeout)
	}
	formData.Add("timeout", timeout)

	if len(params.AllowedUpdates) != 0 {
		allowedUpdatesJSON, _ := json.Marshal(params.AllowedUpdates)
		formData.Add("allowed_updates", string(allowedUpdatesJSON))
	}

	req, err := http.NewRequest(http.MethodGet, b.urlWithToken+getUpdatesUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.URL.RawQuery = formData.Encode()
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := b.api.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var result types.APIResponse[types.Update]

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	if !result.Ok {
		return nil, fmt.Errorf(
			"telegram API returned not ok: error code: %d, description: %s",
			result.ErrorCode, result.Description,
		)
	}

	return result.Result, nil
}
