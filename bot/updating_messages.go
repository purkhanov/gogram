package bot

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/purkhanov/gogram/types"
)

const (
	deleteMessageUrl  = "/deleteMessage"
	deleteMessagesUrl = "/deleteMessages"
)

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
		return fmt.Errorf("telegram API error: code %d - %s", result.ErrorCode, result.Description)
	}

	return nil
}
