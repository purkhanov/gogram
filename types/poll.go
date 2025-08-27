package types

type Poll struct {
	ID string `json:"id"` // Unique poll identifier

	// Poll question, 1-300 characters
	Question string `json:"question"`

	// Optional. Special entities that appear in the question. Currently,
	// only custom emoji entities are allowed in poll questions
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`

	// List of poll options
	Options []PollOption `json:"options"`

	// Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`

	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`

	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`

	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// Optional. 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode, which are closed,
	// or was sent (not forwarded) by the bot or to the private
	// chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Optional. Text that is shown when a user chooses an incorrect
	// answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`

	// Optional. Special entities like usernames, URLs, bot
	// commands, etc. that appear in the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period,omitempty"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int `json:"close_date,omitempty"`
}

type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`

	// Optional. Special entities that appear in the option text.
	// Currently, only custom emoji entities are allowed in poll option texts
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

type PollAnswer struct {
	// Unique poll identifier
	PollID string `json:"poll_id"`

	// Optional. The chat that changed the answer to the poll,
	// if the voter is anonymous
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// Optional. The user that changed the answer to the poll,
	// if the voter isn't anonymous
	User *User `json:"user,omitempty"`

	// 0-based identifiers of chosen answer options.
	// May be empty if the vote was retracted.
	OptionIDs []int `json:"option_ids"`
}
