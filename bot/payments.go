package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/purkhanov/gogram/api"
	"github.com/purkhanov/gogram/types"
	"github.com/purkhanov/gogram/utils"
)

const (
	sendInvoiceUrl            = "/sendInvoice"
	answerPreCheckoutQueryUrl = "/answerPreCheckoutQuery"
	createInvoiceLinkUrl      = "/createInvoiceLink"
	answerShippingQueryUrl    = "/answerShippingQuery"
)

type SendInvoiceOptions struct {
	// Unique identifier for the target chat or username
	// of the target channel (in the format @channelusername)
	ChatID uint `json:"chat_id" validate:"required"`

	// Product name, 1-32 characters
	Title string `json:"title" validate:"required"`

	// Product description, 1-255 characters
	Description string `json:"description" validate:"required"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be
	// displayed to the user, use it for your internal processes.
	Payload string `json:"payload" validate:"required"`

	// Three-letter ISO 4217 currency code, see more on
	// currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency" validate:"required"`

	// Price breakdown, a JSON-serialized list of components
	// (e.g. product price, tax, discount, delivery cost,
	// delivery tax, bonus, etc.). Must contain exactly one
	// item for payments in Telegram Stars.
	Prices []types.LabeledPrice `json:"prices" validate:"required"`

	// Unique identifier for the target message thread (topic)
	// of the forum; for forum supergroups only
	MessageThreadID uint `json:"message_thread_id,omitempty"`

	// Identifier of the direct messages topic to which the message will
	// be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicID uint `json:"direct_messages_topic_id,omitempty"`

	// Payment provider token, obtained via @BotFather. Pass
	// an empty string for payments in Telegram Stars.
	ProviderToken string `json:"provider_token,omitempty"`

	// The maximum accepted amount for tips in the smallest
	// units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.45 pass
	// max_tip_amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies). Defaults to 0.
	// Not supported for payments in Telegram Stars.
	MaxTipAmount uint `json:"max_tip_amount,omitempty"`

	// A JSON-serialized array of suggested amounts of tips in the
	// smallest units of the currency (integer, not float/double).
	// \At most 4 suggested tip amounts can be specified. The suggested
	// tip amounts must be positive, passed in a strictly increased
	// order and must not exceed max_tip_amount.
	SuggestedTipAmounts []uint `json:"suggested_tip_amounts,omitempty"`

	// Unique deep-linking parameter. If left empty, forwarded
	// copies of the sent message will have a Pay button,
	// allowing multiple users to pay directly from the
	// forwarded message, using the same invoice. If non-empty,
	// forwarded copies of the sent message will have a URL
	// button with a deep link to the bot (instead of a Pay button),
	// with the value used as the start parameter
	StartParameter string `json:"start_parameter,omitempty"`

	// JSON-serialized data about the invoice, which will be
	// shared with the payment provider. A detailed description
	// of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// URL of the product photo for the invoice. Can be a
	// photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	PhotoUrl string `json:"photo_url,omitempty"`

	// Photo size in bytes
	PhotoSize   uint16 `json:"photo_size,omitempty"`
	PhotoWidth  uint16 `json:"photo_width,omitempty"`
	PhotoHeight uint16 `json:"photo_height,omitempty"`

	// Pass True if you require the user's full name to complete
	// the order. Ignored for payments in Telegram Stars.
	NeedName bool `json:"need_name,omitempty"`

	// Pass True if you require the user's phone number to complete
	// the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Pass True if you require the user's email address to
	// complete the order. Ignored for payments in Telegram Stars.
	NeedEmail bool `json:"need_email,omitempty"`

	// Pass True if you require the user's shipping address to
	// complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Pass True if the user's phone number should be sent to
	// the provider. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// Pass True if the user's email address should be sent to
	// the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// Pass True if the final price depends on the shipping
	// method. Ignored for payments in Telegram Stars.
	IsFlexible bool `json:"is_flexible,omitempty"`

	// Sends the message silently. Users will
	// receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Pass True to allow up to 1000 messages per second, ignoring
	// broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// Unique identifier of the message effect to be
	// added to the message; for private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// A JSON-serialized object containing the parameters of the
	// suggested post to send; for direct messages chats only.
	// If the message is sent as a reply to another suggested post,
	// then that suggested post is automatically declined.
	SuggestedPostParameters any `json:"suggested_post_parameters,omitempty"`

	// Description of the message to reply to
	ReplyParameters any `json:"reply_parameters,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty,
	// one 'Pay total price' button will be shown. If not empty,
	// the first button must be a Pay button.
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send invoices. On success,
// the sent Message is returned.
func (b *Bot) SendInvoice(param SendInvoiceOptions) error {
	if err := utils.ValidateStruct(param); err != nil {
		return err
	}

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}

	c, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+sendInvoiceUrl, data,
	)

	if err != nil {
		return err
	}

	var result api.APIResponse[types.Message]

	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("telegram API error: code %d - %s", result.ErrorCode, result.Description)
	}

	return nil
}

type AnswerPreCheckoutQueryOptions struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id" validate:"required"`

	// Specify True if everything is alright (goods are
	// available, etc.) and the bot is ready to proceed
	// with the order. Use False if there are any problems.
	Ok bool `json:"ok"`

	// Required if ok is False. Error message in human readable
	// form that explains the reason for failure to proceed with
	// the checkout (e.g. "Sorry, somebody just bought the last
	// of our amazing black T-shirts while you were busy filling
	// out your payment details. Please choose a different color
	// or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

func (b *Bot) AnswerPreCheckoutQuery(params AnswerPreCheckoutQueryOptions) error {
	if err := utils.ValidateStruct(params); err != nil {
		return err
	}

	data, err := json.Marshal(params)
	if err != nil {
		return err
	}

	c, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+answerPreCheckoutQueryUrl, data,
	)
	if err != nil {
		return err
	}

	var result api.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("telegram API error: code %d - %s", result.ErrorCode, result.Description)
	}

	return nil
}

type CreateInvoiceLinkOptions struct {
	// Unique identifier of the business connection on behalf of which
	// the link will be created. For payments in Telegram Stars only.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Product name, 1-32 characters
	Title string `json:"title" validate:"required"`

	// Product description, 1-255 characters
	Description string `json:"description" validate:"required"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be
	// displayed to the user, use it for your internal processes.
	Payload string `json:"payload" validate:"required"`

	// Payment provider token, obtained via @BotFather.
	// Pass an empty string for payments in Telegram Stars.
	ProviderToken string `json:"provider_token,omitempty"`

	// Three-letter ISO 4217 currency code, see more on currencies
	// (https://core.telegram.org/bots/payments#supported-currencies)
	// Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency" validate:"required"`

	// rice breakdown, a JSON-serialized list of components
	//  (e.g. product price, tax, discount, delivery cost,
	// delivery tax, bonus, etc.). Must contain exactly one
	// item for payments in Telegram Stars.
	Prices []types.LabeledPrice `json:"prices" validate:"required"`

	// The number of seconds the subscription will be active for
	// before the next payment. The currency must be set to “XTR”
	// (Telegram Stars) if the parameter is used. Currently, it must
	// always be 2592000 (30 days) if specified. Any number of
	// subscriptions can be active for a given bot at the same time,
	// including multiple concurrent subscriptions from the same user.
	// Subscription price must no exceed 10000 Telegram Stars.
	SubscriptionPeriod uint32 `json:"subscription_period,omitempty"`

	// The maximum accepted amount for tips in the smallest units
	// of the currency (integer, not float/double). For example,
	// for a maximum tip of US$ 1.45 pass max_tip_amount = 145.
	// See the exp parameter in currencies.json, it shows the
	// number of digits past the decimal point for each currency
	// (2 for the majority of currencies). Defaults to 0.
	// Not supported for payments in Telegram Stars.
	MaxTipAmount uint32 `json:"max_tip_amount,omitempty"`

	// A JSON-serialized array of suggested amounts of tips in the smallest
	// units of the currency (integer, not float/double). At most 4 suggested
	// tip amounts can be specified. The suggested tip amounts must be positive,
	// passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []uint32 `json:"suggested_tip_amounts,omitempty"`

	// JSON-serialized data about the invoice, which will be shared
	// with the payment provider. A detailed description of required
	// fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// URL of the product photo for the invoice. Can be a
	// photo of the goods or a marketing image for a service.
	PhotoUrl string `json:"photo_url,omitempty"`

	// Photo size in bytes
	PhotoSize uint16 `json:"photo_size,omitempty"`

	//  	Photo width
	Photo_width uint16 `json:"photo_width,omitempty"`

	// Photo height
	PhotoHeight uint16 `json:"photo_height,omitempty"`

	// Pass True if you require the user's full name to complete
	// the order. Ignored for payments in Telegram Stars.
	NeedName bool `json:"need_name,omitempty"`

	// Pass True if you require the user's phone number to complete
	// the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Pass True if you require the user's email address to complete
	// the order. Ignored for payments in Telegram Stars.
	NeedEmail bool `json:"need_email,omitempty"`

	// Pass True if you require the user's shipping address to
	// complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Pass True if the user's phone number should be sent to
	// the provider. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// Pass True if the user's email address should be sent to
	// the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// Pass True if the final price depends on the shipping
	// method. Ignored for payments in Telegram Stars.
	IsFlexible bool `json:"is_flexible,omitempty"`
}

func (b *Bot) CreateInvoiceLink(params CreateInvoiceLinkOptions) (string, error) {
	if err := utils.ValidateStruct(params); err != nil {
		return "", err
	}

	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	c, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+createInvoiceLinkUrl, data,
	)
	if err != nil {
		return "", err
	}

	var result api.APIResponse[string]

	if err := json.Unmarshal(resp, &result); err != nil {
		return "", err
	}

	if !result.Ok {
		return "", fmt.Errorf(
			"telegram API error: code %d - %s", result.ErrorCode, result.Description,
		)
	}

	return result.Result, nil
}

type AnswerShippingQueryOptions struct {
	// Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id" validate:"required"`

	// Pass True if delivery to the specified address is possible
	// and False if there are any problems (for example, if
	// delivery to the specified address is not possible)
	Ok bool `json:"ok" validate:"required"`

	// Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []types.ShippingOption `json:"shipping_options,omitempty"`

	// Required if ok is False. Error message in human readable
	// form that explains why it is impossible to complete the
	// order (e.g. “Sorry, delivery to your desired address is
	// unavailable”). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

func (b *Bot) AnswerShippingQuery(params AnswerShippingQueryOptions) error {
	if len(params.ShippingQueryID) == 0 {
		return fmt.Errorf("shipping_query_id is required")
	}

	if params.Ok && len(params.ShippingOptions) == 0 {
		return fmt.Errorf("shipping_options are required when ok is true")
	}

	if !params.Ok && len(params.ErrorMessage) == 0 {
		return fmt.Errorf("error_message is required when ok is false")
	}

	data, err := json.Marshal(params)
	if err != nil {
		return err
	}

	c, cancel := context.WithTimeout(b.ctx, httpRequestTimeout)
	defer cancel()

	resp, err := b.api.DoRequestWithContextAndData(
		c, http.MethodPost, b.urlWithToken+answerShippingQueryUrl, data,
	)
	if err != nil {
		return err
	}

	var result api.APIResponse[bool]

	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf(
			"telegram API error: code %d - %s", result.ErrorCode, result.Description,
		)
	}
	return nil
}
