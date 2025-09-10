package types

// This object contains information about a message that is being
// replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	// Origin of the message replied to by the given message
	Origin *MessageOrigin `json:"origin"`

	// Optional. Chat the original message belongs to.
	// Available only if the chat is a supergroup or a channel.
	Chat *Chat `json:"chat,omitempty"`

	// Optional. Unique message identifier inside the original chat.
	// Available only if the original chat is a supergroup or a channel.
	MessageID int `json:"message_id,omitempty"`

	// Optional. Options used for link preview generation
	// for the original message, if it is a text message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a checklist
	Checklist *Checklist `json:"checklist,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Optional. Message is a game, information about the game.
	// More about games » (https://core.telegram.org/bots/api#games)
	Game *Game `json:"game,omitempty"`

	// Optional. Message is a scheduled giveaway, information about the giveaway
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Optional. Message is an invoice for a payment, information about the
	// invoice. More about payments » (https://core.telegram.org/bots/api#payments)
	Invoice *Invoice `json:"invoice,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
}
