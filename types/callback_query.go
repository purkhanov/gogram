package types

type CallbackQuery struct {
	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From User `json:"from"`

	// Optional. Message sent by the bot with the
	// callback button that originated the query
	Message *any `json:"message,omitempty"`

	// Optional. Identifier of the message sent via the bot in
	// inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Global identifier, uniquely corresponding to the chat
	// to which the message with the callback button was sent.
	// Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`

	// Optional. Data associated with the callback button.
	// Be aware that the message originated the query can
	// contain no callback buttons with this data.
	Data string `json:"data,omitempty"`

	// Optional. Short name of a Game to be returned, serves as
	// the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}
