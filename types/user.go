package types

type User struct {
	// Unique identifier for this user or bot. This number
	// may have more than 32 significant bits and some
	// programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	ID int `json:"id"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`

	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`

	// Optional. True, if the bot can be connected to a Telegram Business
	// account to receive its messages. Returned only in getMe.
	CanConnectToBusiness bool `json:"can_connect_to_business,omitempty"`

	// Optional. True, if the bot has a main Web App. Returned only in getMe.
	HasMainWebApp bool `json:"has_main_web_app,omitempty"`
}
