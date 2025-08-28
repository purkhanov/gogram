package types

// Represents the rights of a business bot.
type BusinessBotRights struct {
	// Optional. True, if the bot can send and edit messages in
	// the private chats that had incoming messages in the last 24 hours
	CanReply bool `json:"can_reply,omitempty"`

	// Optional. True, if the bot can mark incoming private messages as read
	CanReadMessages bool `json:"can_read_messages,omitempty"`

	// Optional. True, if the bot can delete messages sent by the bot
	CanDeleteSentMessages bool `json:"can_delete_sent_messages,omitempty"`

	// Optional. True, if the bot can delete all private messages in managed chats
	CanDeleteAllMessages bool `json:"can_delete_all_messages,omitempty"`

	// Optional. True, if the bot can edit the first and last name of the business account
	CanEditName bool `json:"can_edit_name,omitempty"`

	// Optional. True, if the bot can edit the bio of the business account
	CanEditBio bool `json:"can_edit_bio,omitempty"`

	// Optional. True, if the bot can edit the profile photo of the business account
	CanEditProfilePhoto bool `json:"can_edit_profile_photo,omitempty"`

	// Optional. True, if the bot can edit the username of the business account
	CanEditUsername bool `json:"can_edit_username,omitempty"`

	// Optional. True, if the bot can change the privacy settings
	// pertaining to gifts for the business account
	CanChangeGiftSettings bool `json:"can_change_gift_settings,omitempty"`

	// Optional. True, if the bot can view gifts and the amount of
	// Telegram Stars owned by the business account
	CanViewGiftsAndStars bool `json:"can_view_gifts_and_stars,omitempty"`

	// Optional. True, if the bot can convert regular gifts owned
	// by the business account to Telegram Stars
	CanConvertGiftsToStars bool `json:"can_convert_gifts_to_stars,omitempty"`

	// Optional. True, if the bot can transfer and upgrade gifts owned by the business account
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`

	// Optional. True, if the bot can transfer Telegram Stars received by the business
	// account to its own account, or use them to upgrade and transfer gifts
	CanTransferStars bool `json:"can_transfer_stars,omitempty"`

	// Optional. True, if the bot can post, edit and delete
	// stories on behalf of the business account
	CanManageStories bool `json:"can_manage_stories,omitempty"`
}

// Describes the connection of the bot with a business account.
type BusinessConnection struct {
	// Unique identifier of the business connection
	ID string `json:"id"`

	// Business account user that created the business connection
	User *User `json:"user"`

	// Identifier of a private chat with the user who created the
	// business connection. This number may have more than 32 significant
	// bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits,
	// so a 64-bit integer or double-precision float type are safe for
	// storing this identifier.
	UserChatID int `json:"user_chat_id"`

	// Date the connection was established in Unix time
	Date int `json:"date"`

	// Optional. Rights of the business bot
	Rights *BusinessBotRights `json:"rights,omitempty"`

	// True, if the connection is active
	IsEnabled bool `json:"is_enabled"`
}
