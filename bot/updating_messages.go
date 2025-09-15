package bot

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/purkhanov/gogram/types"
)

// Updating messages
// The following methods allow you to change an existing message
// in the message history instead of sending a new one with a
// result of an action. This is most useful for messages with
// inline keyboards using callback queries, but can also help
// reduce clutter in conversations with regular chat bots.

const (
	deleteMessageUrl  = "/deleteMessage"
	deleteMessagesUrl = "/deleteMessages"
)

type EditMessageTextOptions struct {
	// Unique identifier of the business connection on
	// behalf of which the message to be edited was sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username
	// of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified.
	// Identifier of the message to edit
	MessageID uint `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New text of the message, 1-4096
	// characters after entities parsing
	Text string `json:"text" validate:"required"`

	// Mode for parsing entities in the message text.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in
	// message text, which can be specified instead of parse_mode
	Entities []types.MessageEntity `json:"entities,omitempty"`

	// link_preview_options
	LinkPreviewOptions *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit text and game messages. On success, if the
// edited message is not an inline message, the edited Message is
// returned, otherwise True is returned. Note that business messages
// that were not sent by the bot and do not contain an inline keyboard
// can only be edited within 48 hours from the time they were sent.
func (b *Bot) EditMessageText(options EditMessageTextOptions) (bool, error) {
	if options.InlineMessageID != "" {
		if options.ChatID != "" || options.MessageID != 0 {
			return false, fmt.Errorf(
				"ChatID and MessageID should not be specified for inline messages",
			)
		}
	}

	if options.ChatID == "" {
		return false, fmt.Errorf("ChatID is required for non-inline messages")
	}

	if options.MessageID == 0 {
		return false, fmt.Errorf("MessageID is required for non-inline messages")
	}

	if options.Text == "" {
		return false, fmt.Errorf("text is required")
	}

	data, err := json.Marshal(options)
	if err != nil {
		return false, err
	}

	requestUrl := b.urlWithToken + "/editMessageTextUrl"

	resp, err := b.api.DoRequestWithTimeout(http.MethodPost, requestUrl, data)
	if err != nil {
		return false, err
	}

	var result types.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return false, err
	}

	if !result.Ok {
		return false, fmt.Errorf(
			"telegram API error: code %d - %s",
			result.ErrorCode, result.Description,
		)
	}

	return true, nil
}

// Please note, that it is currently only possible to edit
// messages without reply_markup or with inline keyboards.

// Use this method to delete a message, including service
// messages, with the following limitations:
//
// - A message can only be deleted if it was sent less than 48 hours ago.
//
// - Service messages about a supergroup, channel,
// or forum topic creation can't be deleted.
//
// - A dice message in a private chat can only be
// deleted if it was sent more than 24 hours ago.
//
// - Bots can delete outgoing messages in
// private chats, groups, and supergroups.
//
// - Bots can delete incoming messages in private chats.
//
// - Bots granted can_post_messages permissions
// can delete outgoing messages in channels.
//
// - If the bot is an administrator of a
// group, it can delete any message there.
//
// - If the bot has can_delete_messages administrator right
// in a supergroup or a channel, it can delete any message there.
//
// - If the bot has can_manage_direct_messages administrator right in a
// channel, it can delete any message in the corresponding direct messages chat.
//
// Returns True on success.
func (b *Bot) DeleteMessage(chatID, messageID uint) error {
	data := map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	return b.deleteMsgs(data, b.urlWithToken+deleteMessageUrl)
}

func (b *Bot) DeleteMessages(chatID uint, messageIDs []uint) error {
	data := map[string]any{
		"chat_id":     chatID,
		"message_ids": messageIDs,
	}

	return b.deleteMsgs(data, b.urlWithToken+deleteMessagesUrl)
}

func (b *Bot) deleteMsgs(data map[string]any, url string) error {
	c, cancel := context.WithTimeout(b.Ctx, httpRequestTimeout)
	defer cancel()

	dataByte, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	resp, err := b.api.DoRequestWithContextAndData(c, http.MethodPost, url, dataByte)
	if err != nil {
		return err
	}

	var result types.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf(
			"telegram API error: code %d - %s",
			result.ErrorCode, result.Description,
		)
	}

	return nil
}
