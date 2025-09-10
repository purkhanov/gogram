package types

// This object represents a message about a scheduled giveaway.
type Giveaway struct {
	// The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`

	// Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int `json:"winners_selection_date"`

	// The number of users which are supposed to
	// be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`

	// Optional. True, if only users who join the chats
	// after the giveaway started should be eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// Optional. True, if the list of giveaway
	// winners will be visible to everyone
	HasPublicWinners bool `json:"has_public_winners,omitempty"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`

	// Optional. A list of two-letter ISO 3166-1 alpha-2 country
	// codes indicating the countries from which eligible users
	// for the giveaway must come. If empty, then all users can
	// participate in the giveaway. Users with a phone number that
	// was bought on Fragment can always participate in giveaways.
	CountryCodes []string `json:"country_codes,omitempty"`

	// Optional. The number of Telegram Stars to be split between
	// giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// Optional. The number of months the Telegram Premium subscription won
	// from the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`
}

// This object represents a message about the
// completion of a giveaway with public winners.
type GiveawayWinners struct {
	// The chat that created the giveaway
	Chat *Chat `json:"chat"`

	// Identifier of the message with the giveaway in the chat
	GiveawayMessageID int `json:"giveaway_message_id"`

	// Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int `json:"winners_selection_date"`

	// Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`

	// List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`

	// Optional. The number of other chats the user had
	// to join in order to be eligible for the giveaway
	AdditionalChatCount int `json:"additional_chat_count,omitempty"`

	// Optional. The number of Telegram Stars that were split
	// between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// Optional. The number of months the Telegram Premium
	// subscription won from the giveaway will be active for;
	// for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`

	// Optional. Number of undistributed prizes
	UnclaimedPrizeCount int `json:"unclaimed_prize_count,omitempty"`

	// Optional. True, if only users who had joined the chats
	// after the giveaway started were eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// Optional. True, if the giveaway was canceled
	// because the payment for it was refunded
	WasRefunded bool `json:"was_refunded,omitempty"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`
}
