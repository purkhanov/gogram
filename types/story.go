package types

type Story struct {
	// Chat that posted the story
	Chat *Chat `json:"chat"`

	// Unique identifier for the story in the chat
	ID int `json:"id"`
}
