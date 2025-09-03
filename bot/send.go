package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/purkhanov/gogram/types"
	"github.com/purkhanov/gogram/utils"
)

const (
	sendMessageUrl      = "/sendMessage"
	sendAudioUrl        = "/sendAudio"
	answerCallbackQuery = "/answerCallbackQuery"
)

type ReplyMarkup interface {
	ValidateReplyMarkup() error
}

type SendMessageOptions struct {
	// Unique identifier of the business connection on
	// behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Unique identifier for the target chat or username of the
	// target channel (in the format @channelusername)
	ChatID uint `json:"chat_id" validate:"required"`

	// Unique identifier for the target message thread (topic)
	// of the forum; for forum supergroups only
	MessageThreadID uint `json:"message_thread_id,omitempty"`

	// dentifier of the direct messages topic to which
	// the message will be sent; required if the message
	// is sent to a direct messages chat
	DirectMessagesTopicID uint `json:"direct_messages_topic_id,omitempty"`

	// Text of the message to be sent, 1-4096
	// characters after entities parsing
	Text string `json:"text" validate:"required"`

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
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send text messages.
// On success, the sent Message is returned.
func (b *Bot) SendMessage(params SendMessageOptions) error {
	if err := utils.ValidateStruct(params); err != nil {
		return err
	}

	data, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal params: %w", err)
	}

	c, cancel := context.WithTimeout(b.Ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+sendMessageUrl, data,
	)
	if err != nil {
		return err
	}

	var result types.APIResponse[types.Message]

	if err := json.Unmarshal(resp, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !result.Ok {
		return fmt.Errorf(
			"telegram API error: code %d - %s", result.ErrorCode, result.Description,
		)
	}

	return nil
}

type SendVoiceOptions struct {
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

	// Audio file to send. Pass a file_id as String to send a file
	// that exists on the Telegram servers (recommended), pass an
	// HTTP URL as a String for Telegram to get a file from the
	// Internet, or upload a new one using multipart/form-data.
	// More information on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Voice string `json:"voice"`

	// Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the message text.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the
	// caption, which can be specified instead of parse_mode
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`

	// Duration of the voice message in seconds
	Duration uint `json:"duration,omitempty"`

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
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send audio files, if you want Telegram
// clients to display them in the music player. Your audio must
// be in the .MP3 or .M4A format. On success, the sent Message is
// returned. Bots can currently send audio files of up to 50 MB in
// size, this limit may be changed in the future.
func (b *Bot) SendAudio(params SendVoiceOptions) (types.Message, error) {
	var response types.APIResponse[types.Message]

	return response.Result, nil
}

type AnswerCallbackQueryOptions struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id" validate:"required"`

	// Text of the notification. If not specified, nothing
	// will be shown to the user, 0-200 characters
	Text string `json:"text,omitempty"`

	// If True, an alert will be shown by the client
	// instead of a notification at the top of the
	// chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// URL that will be opened by the user's client. If you have
	// created a Game and accepted the conditions via @BotFather,
	// specify the URL that opens your game - note that this will
	// only work if the query comes from a callback_game button.
	Url string `json:"url,omitempty"`

	// The maximum amount of time in seconds that the result of the
	// callback query may be cached client-side. Telegram apps will
	// support caching starting in version 3.14. Defaults to 0.
	CacheTime uint `json:"cache_time,omitempty"`
}

func (b *Bot) AnswerCallbackQuery(params AnswerCallbackQueryOptions) error {
	if err := utils.ValidateStruct(params); err != nil {
		return err
	}

	data, err := json.Marshal(params)
	if err != nil {
		return err
	}

	c, cancel := context.WithTimeout(b.Ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+answerCallbackQuery, data,
	)
	if err != nil {
		return err
	}

	var result types.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return fmt.Errorf("failed to unmarshal result: %w", err)
	}

	if !result.Ok {
		return fmt.Errorf(
			"telegram API error: code %d - %s",
			result.ErrorCode, result.Description,
		)
	}

	return nil
}
