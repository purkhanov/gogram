package types

type InlineQuery struct {
	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From User `json:"from"`

	// Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`

	// Optional. Type of the chat from which the inline
	// query was sent. Can be either “sender” for a private
	// chat with the inline query sender, “private”, “group”,
	// “supergroup”, or “channel”. The chat type should be
	// always known for requests sent from official clients
	// and most third-party clients, unless the request was
	// sent from a secret chat
	ChatType string `json:"chat_type,omitempty"`

	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// Represents a result of an inline query that was chosen
// by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`

	// The user that chose the result
	From User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// Optional. Identifier of the sent inline message. Available
	// only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to
	// edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// The query that was used to obtain the result
	Query string `json:"query,omitempty"`
}
