package types

// Describes the paid media added to a message.
type PaidMediaInfo struct {
	// The number of Telegram Stars that must be paid to buy access to the media
	StarCount int `json:"star_count"`

	// Information about the paid media
	PaidMedia []PaidMedia `json:"paid_media"`
}

// This object describes paid media. Currently, it can be one of
type PaidMedia interface {
	// PaidMediaPreview
	// PaidMediaPhoto
	// PaidMediaVideo
	getPaidMediaType() string
}

// The paid media isn't available before the payment.
type PaidMediaPreview struct {
	// Type of the paid media, always “preview”
	Type string `json:"type"`

	// Optional. Media width as defined by the sender
	Width int `json:"width,omitempty"`

	// Optional. Media height as defined by the sender
	Height int `json:"height,omitempty"`

	// Optional. Duration of the media in seconds as defined by the sender
	Duration int `json:"duration,omitempty"`
}

// The paid media is a photo.
type PaidMediaPhoto struct {
	// Type of the paid media, always “photo”
	Type string `json:"type"`

	// The photo
	Photo []PhotoSize `json:"photo"`
}

// The paid media is a video.
type PaidMediaVideo struct {
	// Type of the paid media, always “video”
	Type string `json:"type"`

	// The video
	Video *Video `json:"video"`
}
