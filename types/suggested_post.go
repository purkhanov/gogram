package types

type SuggestedPostInfo struct {
	// State of the suggested post. Currently, it can
	// be one of “pending”, “approved”, “declined”.
	State string `json:"state"`

	// Optional. Proposed price of the post. If the
	// field is omitted, then the post is unpaid.
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// Optional. Proposed send date of the post. If the
	// field is omitted, then the post can be published
	// at any time within 30 days at the sole discretion
	// of the user or administrator who approves it.
	SendDate int `json:"send_date,omitempty"`
}

type SuggestedPostPrice struct {
	// Currency in which the post will be paid. Currently, must
	// be one of “XTR” for Telegram Stars or “TON” for toncoins
	Currency string `json:"currency"`

	// The amount of the currency that will be paid for the post
	// in the smallest units of the currency, i.e. Telegram Stars
	// or nanotoncoins. Currently, price in Telegram Stars must
	// be between 5 and 100000, and price in nanotoncoins must
	// be between 10000000 and 10000000000000.
	Amount int `json:"amount"`
}

type SuggestedPostParameters struct {
	// Optional. Proposed price for the post. If the field is
	// omitted, then the post is unpaid.
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// Optional. Proposed send date of the post. If specified,
	// then the date must be between 300 second and 2678400 seconds
	// (30 days) in the future. If the field is omitted, then the
	// post can be published at any time within 30 days at the sole
	// discretion of the user who approves it.
	SendDate int `json:"send_date,omitempty"`
}
