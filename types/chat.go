package types

type Chat struct {
	// Unique identifier for this chat. This number may
	// have more than 32 significant bits and some
	// programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	ID uint `json:"id"`

	// Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// Optional. Last name of the other party in a private chat
	LastMame string `json:"last_name,omitempty"`

	// Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum string `json:"is_forum,omitempty"`
}

type ChatMemberUpdated struct {
	// Chat the user belongs to
	Chat Chat `json:"chat"`

	// Performer of the action, which resulted in the change
	From User `json:"from"`

	// Date the change was done in Unix time
	Date int `json:"date"`

	// Previous information about the chat member
	OldChatMember ChatMember `json:"old_chat_member"`

	// New information about the chat member
	NewChatMember ChatMember `json:"new_chat_member"`

	// Optional. Chat invite link, which was used by the user to
	// join the chat; for joining by invite link events only.
	InviteLink ChatInviteLink `json:"invite_link"`

	// Optional. True, if the user joined the chat after
	// sending a direct join request without using an invite
	// link and being approved by an administrator
	ViaJoinRequest bool `json:"via_join_request,omitempty"`

	// Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatMember struct{}

type ChatInviteLink struct {
	// The invite link. If the link was created by another chat
	// administrator, then the second part of the link will be
	// replaced with “…”.
	InviteLink ChatMember `json:"invite_link"`

	// Creator of the link
	Creator User `json:"creator"`

	// True, if users joining the chat via the link
	// need to be approved by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request"`

	// True, if the link is primary
	IsPrimary bool `json:"is_primary"`

	// True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`

	// Optional. Invite link name
	Name string `json:"name,omitempty"`

	// Optional. Point in time (Unix timestamp) when
	// the link will expire or has been expired
	ExpireDate int `json:"expire_date,omitempty"`

	// Optional. The maximum number of users that can be
	// members of the chat simultaneously after joining
	// the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`

	// Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`

	// Optional. The number of seconds the subscription
	// will be active for before the next payment
	SubscriptionPeriod int `json:"subscription_period,omitempty"`

	// Optional. The amount of Telegram Stars a user must pay initially
	// and after each subsequent subscription period to be a member of
	// the chat using the link
	SubscriptionPrice int `json:"subscription_price,omitempty"`
}

// This object contains information about a chat that was
// shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	// Identifier of the request
	RequestID int `json:"request_id"`

	// Identifier of the shared chat. This number may
	// have more than 32 significant bits and some
	// programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	// The bot may not have access to the chat and could
	// be unable to use this identifier, unless the chat
	// is already known to the bot by some other means.
	ChatID int `json:"chat_id"`

	// Optional. Title of the chat, if the title was requested by the bot.
	Title string `json:"title,omitempty"`

	// Optional. Username of the chat, if the
	// username was requested by the bot and available.
	Username string `json:"username,omitempty"`

	// Optional. Available sizes of the chat photo,
	// if the photo was requested by the bot
	Photo []PhotoSize `json:"photo,omitempty"`
}

// This object represents a service message about a
// user allowing a bot to write messages after adding
// it to the attachment menu, launching a Web App
// from a link, or accepting an explicit request from
// a Web App sent by the method requestWriteAccess.
// (https://core.telegram.org/bots/webapps#initializing-mini-apps)
type WriteAccessAllowed struct {
	// Optional. True, if the access was granted
	// after the user accepted an explicit request
	// from a Web App sent by the method requestWriteAccess
	// (https://core.telegram.org/bots/webapps#initializing-mini-apps)
	FromRequest bool `json:"from_request,omitempty"`

	// Optional. Name of the Web App, if the access was
	// granted when the Web App was launched from a link
	WebAppName string `json:"web_app_name,omitempty"`

	// Optional. True, if the access was granted when
	// the bot was added to the attachment or side menu
	FromAttachmentMenu bool `json:"from_attachment_menu,omitempty"`
}

type ChatJoinRequest struct {
	// Chat to which the request was sent
	Chat Chat `json:"chat"`

	// User that sent the join request
	From User `json:"from"`

	// Identifier of a private chat with the user who sent the
	// join request. This number may have more than 32 significant
	// bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant
	// bits, so a 64-bit integer or double-precision float type are
	// safe for storing this identifier. The bot can use this identifier
	// for 5 minutes to send messages until the join request is processed,
	// assuming no other administrator contacted the user.
	UserChatID int `json:"user_chat_id"`

	// Date the request was sent in Unix time
	Date int `json:"date"`

	// Optional. Bio of the user.
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Information about the chat boost
}

type ChatBoostRemoved struct {
	Chat    Chat `json:"chat"`     // Chat which was boosted
	BoostID Chat `json:"boost_id"` // Unique identifier of the boost

	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int `json:"remove_date"`

	Source ChatBoostSource `json:"source"` // Source of the removed boost
}

type ChatBoost struct {
	BoostID string `json:"boost_id"` // Unique identifier of the boost

	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int `json:"add_date"`

	// Point in time (Unix timestamp) when the boost
	// will automatically expire, unless the booster's
	// Telegram Premium subscription is prolonged
	ExpirationDate int `json:"expiration_date"`

	// Source of the added boost
	Source ChatBoostSource `json:"source"`
}

type ChatBoostSource struct {
	// ChatBoostSourcePremium
	// ChatBoostSourceGiftCode
	// ChatBoostSourceGiveaway
}
