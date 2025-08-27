package types

type TextQuote struct {
	// Text of the quoted part of a message that
	// is replied to by the given message
	Text string `json:"text"`

	// Optional. Special entities that appear in the quote.
	// Currently, only bold, italic, underline, strikethrough,
	// spoiler, and custom_emoji entities are kept in quotes.
	Entities []MessageEntity `json:"entities,omitempty"`

	// Approximate quote position in the original message in
	// UTF-16 code units as specified by the sender
	Position int `json:"position,omitempty"`

	// Optional. True, if the quote was chosen manually
	// by the message sender. Otherwise, the quote was
	// added automatically by the server.
	IsManual bool `json:"is_manual,omitempty"`
}
