package types

// This object contains basic information about a
// successful payment. Note that if the buyer initiates
// a chargeback with the relevant payment provider
// following this transaction, the funds may be debited
// from your balance. This is outside of Telegram's control.
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code, or
	// “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45 pass
	// amount = 145. See the exp parameter in currencies.json, it
	// shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Expiration date of the subscription, in
	// Unix time; for recurring payments only
	SubscriptionExpirationDate int `json:"subscription_expiration_date,omitempty"`

	// Optional. True, if the payment is a recurring payment for a subscription
	IsRecurring bool `json:"is_recurring,omitempty"`

	// Optional. True, if the payment is the first payment for a subscription
	IsFirstRecurring bool `json:"is_first_recurring,omitempty"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id,omitempty"`

	// Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

// This object contains basic information about a refunded payment.
type RefundedPayment struct {
	// Three-letter ISO 4217 currency code, or “XTR” for
	// payments in Telegram Stars. Currently, always “XTR”
	Currency string `json:"currency"`

	// Total refunded price in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45,
	// total_amount = 145. See the exp parameter in currencies.json, it
	// shows the number of digits past the decimal point for each currency
	// (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// Optional. Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// Unique query identifier
	ID string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	// Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// Unique query identifier
	ID string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Three-letter ISO 4217 currency code, or “XTR” for
	// payments in Telegram Stars
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45
	// pass amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionIS string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label string `json:"label"`

	// Price of the product in the smallest units of the currency
	// (integer, not float/double). For example, for a price of
	// US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PaidMediaPurchased struct {
	// User who purchased the media
	From User `json:"from"`

	// Bot-specified paid media payload
	PaidMediaPayload string `json:"paid_media_payload"`
}
