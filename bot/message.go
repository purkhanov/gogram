package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/purkhanov/gogram/types"
)

const (
	sendMessageUrl = "/sendMessage"
)

type SendMessageParams struct {
	// Unique identifier of the business connection on
	// behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Unique identifier for the target chat or username of the
	// target channel (in the format @channelusername)
	ChatID uint `json:"chat_id"`

	// Unique identifier for the target message thread (topic)
	// of the forum; for forum supergroups only
	MessageThreadID uint `json:"message_thread_id,omitempty"`

	// dentifier of the direct messages topic to which
	// the message will be sent; required if the message
	// is sent to a direct messages chat
	DirectMessagesTopicID uint `json:"direct_messages_topic_id,omitempty"`

	// Text of the message to be sent, 1-4096
	// characters after entities parsing
	Text string `json:"text"`

	// Mode for parsing entities in the message text.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in
	// message text, which can be specified instead of parse_mode
	Entities []types.MessageEntity `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions any `json:"link_preview_options,omitempty"`

	// Sends the message silently. Users will
	// receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Pass True to allow up to 1000 messages per second, ignoring
	// broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// Unique identifier of the message effect to be added
	// to the message; for private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// A JSON-serialized object containing the parameters of the
	// suggested post to send; for direct messages chats only.
	// If the message is sent as a reply to another suggested post,
	// then that suggested post is automatically declined.
	SuggestedPostParameters any `json:"suggested_post_parameters,omitempty"`

	// Description of the message to reply to
	ReplyParameters any `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for
	// an inline keyboard, custom reply keyboard, instructions to
	// remove a reply keyboard or to force a reply from the user
	ReplyMarkup any `json:"reply_markup,omitempty"`
}

func (sm *SendMessageParams) validate() error {
	if sm.ChatID == 0 {
		return fmt.Errorf("chat id is required")
	}

	if sm.Text == "" {
		return fmt.Errorf("text is required")
	}

	return nil
}

// Use this method to send text messages.
// On success, the sent Message is returned.
func (b *Bot) SendMessage(params SendMessageParams) (types.Message, error) {
	var response types.APIResponse[types.Message]

	if err := params.validate(); err != nil {
		return response.Result, err
	}

	data, err := json.Marshal(params)
	if err != nil {
		return response.Result, fmt.Errorf("failed to marshal params: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost, b.urlWithToken+sendMessageUrl, bytes.NewBuffer(data),
	)
	if err != nil {
		return response.Result, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := b.api.DoRequest(req)
	if err != nil {
		return response.Result, fmt.Errorf("failed to send request: %w", err)
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response.Result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Result, nil
}
