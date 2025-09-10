package types

// This object represents a gift that can be sent by the bot.
type Gift struct {
	// Unique identifier of the gift
	ID string `json:"id"`

	// The sticker that represents the gift
	Sticker Sticker `json:"sticker"`

	// The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`

	// Optional. The number of Telegram Stars that
	// must be paid to upgrade the gift to a unique one
	UpgradeStarCount int `json:"upgrade_star_count,omitempty"`

	// Optional. The total number of the gifts of this
	// type that can be sent; for limited gifts only
	TotalCount int `json:"total_count,omitempty"`

	// Optional. The number of remaining gifts of this
	// type that can be sent; for limited gifts only
	RemainingCount int `json:"remaining_count,omitempty"`

	// Optional. Information about the chat that published the gift
	PublisherChat *Chat `json:"publisher_chat,omitempty"`
}

// This object represent a list of gifts.
type Gifts struct {
	// The list of gifts
	Gifts []Gift `json:"gifts"`
}

// This object describes the model of a unique gift.
type UniqueGiftModel struct {
	// Name of the model
	Name string `json:"name"`

	// The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`

	// The number of unique gifts that receive
	// this model for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

// This object describes the symbol shown on the pattern of a unique gift.
type UniqueGiftSymbol struct {
	// Name of the symbol
	Name string `json:"name"`

	// The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`

	// The number of unique gifts that receive
	// this model for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

// This object describes the colors of the backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	// The color in the center of the backdrop in RGB format
	CenterColor int `json:"center_color"`

	// The color on the edges of the backdrop in RGB format
	EdgeColor int `json:"edge_color"`

	// The color to be applied to the symbol in RGB format
	SymbolColor int `json:"symbol_color"`

	// The color for the text on the backdrop in RGB format
	TextColor int `json:"text_color"`
}

// This object describes the backdrop of a unique gift.
type UniqueGiftBackdrop struct {
	// Name of the backdrop
	Name string `json:"name"`

	// Colors of the backdrop
	Colors UniqueGiftBackdropColors `json:"colors"`

	// The number of unique gifts that receive
	// this backdrop for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

// This object describes a unique gift
// that was upgraded from a regular gift.
type UniqueGift struct {
	// Human-readable name of the regular gift
	// from which this unique gift was upgraded
	BaseName string `json:"base_name"`

	// Unique name of the gift. This name can be used
	// in https://t.me/nft/... links and story areas
	Name string `json:"name"`

	// Unique number of the upgraded gift among
	// gifts upgraded from the same regular gift
	Number int `json:"number"`

	// Model of the gift
	Model UniqueGiftModel `json:"model"`

	// Symbol of the gift
	Symbol UniqueGiftSymbol `json:"symbol"`

	// Backdrop of the gift
	Backdrop UniqueGiftBackdrop `json:"backdrop"`

	// Optional. Information about the chat that published the gift
	PublisherChat *Chat `json:"publisher_chat"`
}

// Describes a service message about a
// regular gift that was sent or received.
type GiftInfo struct {
	// Information about the gift
	Gift *Gift `json:"gift"`

	// Optional. Unique identifier of the received gift for the bot;
	// only present for gifts received on behalf of business accounts
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// Optional. Number of Telegram Stars that can be claimed by the receiver
	// by converting the gift; omitted if conversion to Telegram Stars is impossible
	ConvertStarCount int `json:"convert_star_count,omitempty"`

	// Optional. Number of Telegram Stars that were prepaid
	// by the sender for the ability to upgrade the gift
	PrepaidUpgradeStarCount int `json:"prepaid_upgrade_star_count,omitempty"`

	// Optional. True, if the gift can be upgraded to a unique gift
	CanBeUpgraded bool `json:"can_be_upgraded,omitempty"`

	// Optional. Text of the message that was added to the gift
	Text string `json:"text,omitempty"`

	// Optional. Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// Optional. True, if the sender and gift text are shown only to
	// the gift receiver; otherwise, everyone will be able to see them
	IsPrivate bool `json:"is_private,omitempty"`
}

// Describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	// Information about the gift
	Gift UniqueGift `json:"gift"`

	// Origin of the gift. Currently, either “upgrade” for
	// gifts upgraded from regular gifts, “transfer” for
	// gifts transferred from other users or channels,
	// or “resale” for gifts bought from other users
	Origin string `json:"origin"`

	// Optional. For gifts bought from other users, the price paid for the gift
	LastResaleStarCount int `json:"last_resale_star_count,omitempty"`

	// Optional. Unique identifier of the received gift for the bot;
	// only present for gifts received on behalf of business accounts
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// Optional. Number of Telegram Stars that must be paid to
	// transfer the gift; omitted if the bot cannot transfer the gift
	TransferStarCount int `json:"transfer_star_count,omitempty"`

	// Optional. Point in time (Unix timestamp) when the gift can be transferred.
	// If it is in the past, then the gift can be transferred now
	NextTransferDate int `json:"next_transfer_date,omitempty"`
}
