package types

type ChatBoost struct {
	// Unique identifier of the boost
	BoostID string `json:"boost_id"`

	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int `json:"add_date"`

	// Point in time (Unix timestamp) when the boost
	// will automatically expire, unless the booster's
	// Telegram Premium subscription is prolonged
	ExpirationDate int `json:"expiration_date"`

	// Source of the added boost
	Source ChatBoostSource `json:"source"`
}

// This object represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	// Number of boosts added by the user
	BoostCount int `json:"boost_count"`
}

type ChatBoostRemoved struct {
	Chat    Chat `json:"chat"`     // Chat which was boosted
	BoostID Chat `json:"boost_id"` // Unique identifier of the boost

	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int `json:"remove_date"`

	Source ChatBoostSource `json:"source"` // Source of the removed boost
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Information about the chat boost
}

type ChatBoostSource struct {
	// ChatBoostSourcePremium
	// ChatBoostSourceGiftCode
	// ChatBoostSourceGiveaway
}
