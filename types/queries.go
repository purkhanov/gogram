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

