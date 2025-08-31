package types

type ShippingQuery struct {
	ID   string `json:"id"`   // Unique query identifier
	From User   `json:"from"` // User who sent the query

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// User specified shipping address
	ShippingAddress string `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID   string `json:"id"`   // Unique query identifier
	From User   `json:"from"` // User who sent the query

	// Three-letter ISO 4217 currency code, or “XTR” for
	// payments in Telegram Stars
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45
	// pass amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies).
	TotaAmount int `json:"total_amount"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionIS string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

type ShippingAddress struct {
	// Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

type ShippingOption struct {
	ID     string         `json:"id"`     // Shipping option identifier
	Title  string         `json:"title"`  // Option title
	Prices []LabeledPrice `json:"prices"` // List of price portions
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label string `json:"label"` // Portion label

	// Price of the product in the smallest units of the currency
	// (integer, not float/double). For example, for a price of
	// US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number,omitempty"`     // Optional. User's phone number
	Email           string           `json:"email,omitempty"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` // Optional. Optional. User shipping address
}

type PaidMediaPurchased struct {
	From             User   `json:"from"`               // User who purchased the media
	PaidMediaPayload string `json:"paid_media_payload"` // Bot-specified paid media payload
}
