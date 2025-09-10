package types

// This object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”.
	// The type of the sticker is independent from its format, which is
	// determined by the fields is_animated and is_video.
	Type string `json:"type"`

	// Sticker width
	Width int `json:"width"`

	// Sticker height
	Height int `json:"height"`

	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated"`

	// True, if the sticker is a video sticker
	IsVideo bool `json:"is_video"`

	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// Optional. For premium regular stickers, premium animation for the sticker
	PremiumAnimation *File `json:"premium_animation,omitempty"`

	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Optional. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	// Optional. True, if the sticker must be repainted to a text color in
	// messages, the color of the Telegram Premium badge in emoji status,
	// white color on chat photos, or another appropriate color in other places
	NeedsRepainting bool `json:"needs_repainting,omitempty"`

	// Optional. File size in bytes`
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`

	// Sticker set title
	Title string `json:"title"`

	// Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`

	// List of all set stickers
	Stickers []Sticker `json:"stickers"`

	// Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumbnail PhotoSize `json:"thumbnail,omitempty"`
}

// This object describes the position on faces where a mask should be placed by default
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`

	// Shift by X-axis measured in widths of the mask scaled to the face
	// size, from left to right. For example, choosing -1.0 will place
	// mask just to the left of the default mask position.
	XShift float32 `json:"x_shift"`

	// Shift by Y-axis measured in heights of the mask scaled to the
	// face size, from top to bottom. For example, 1.0 will place the
	// mask just below the default mask position.
	YShift float32 `json:"y_shift"`

	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}

// This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5
	// for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int `json:"value"`
}
