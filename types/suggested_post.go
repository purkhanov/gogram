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

// Describes a service message about the approval of a suggested post.
type SuggestedPostApproved struct {
	// Optional. Message containing the suggested post. Note that
	// the Message object in this field will not contain the
	// reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Optional. Amount paid for the post
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// Date when the post will be published
	SendDate int `json:"send_date"`
}

// Describes a service message about the
// failed approval of a suggested post.
// Currently, only caused by insufficient
// user funds at the time of approval.
type SuggestedPostApprovalFailed struct {

	// Optional. Message containing the suggested post whose approval
	// has failed. Note that the Message object in this field will not
	// contain the reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Expected price of the post
	Price *SuggestedPostPrice `json:"price"`
}

// Describes a service message about the rejection of a suggested post.
type SuggestedPostDeclined struct {
	// Optional. Message containing the suggested post. Note that
	// the Message object in this field will not contain the
	// reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Optional. Comment with which the post was declined
	Comment string `json:"comment,omitempty"`
}

// Describes a service message about a
// successful payment for a suggested post.
type SuggestedPostPaid struct {
	// Optional. Message containing the suggested post. Note
	// that the Message object in this field will not contain
	// the reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Currency in which the payment was made. Currently,
	// one of “XTR” for Telegram Stars or “TON” for toncoins
	Currency string `json:"currency"`

	// Optional. The amount of the currency
	// that was received by the channel in
	// nanotoncoins; for payments in toncoins only
	Amount int `json:"amount,omitempty"`

	// Optional. The amount of Telegram Stars that was received
	// by the channel; for payments in Telegram Stars only
	StarAmount *StarAmount `json:"star_amount,omitempty"`
}

// Describes a service message about a payment refund for a suggested post.
type SuggestedPostRefunded struct {
	// Optional. Message containing the suggested post. Note that
	// the Message object in this field will not contain the
	// reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Reason for the refund. Currently, one of “post_deleted”
	// if the post was deleted within 24 hours of being posted
	// or removed from scheduled messages without being posted,
	// or “payment_refunded” if the payer refunded their payment.
	Reason string `json:"reason"`
}
