package types

type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Photo width
	Width int `json:"width"`

	// Photo height
	Height int `json:"height"`

	// Optional. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents an animation file
// (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width as defined by the sender
	Width int `json:"width"`

	// Video height as defined by the sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`

	// Optional. Animation thumbnail as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Original animation filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and
	// some programming languages may have difficulty/silent defects
	// in interpreting it. But it has at most 52 significant bits,
	// so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed
	// to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`

	// Optional. Performer of the audio as defined by the sender or by audio tags
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio as defined by the sender or by audio tags
	Title string `json:"title,omitempty"`

	// Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some
	// programming languages may have difficulty/silent defects in interpreting
	// it. But it has at most 52 significant bits, so a signed 64-bit integer
	// or double-precision float type are safe for storing this value.
	FileSize int `json:"file_size,omitempty"`

	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// This object represents a general file (as opposed
// to photos, voice messages and audio files).
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed
	// to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Optional. Document thumbnail as defined by the sender
	Thumbnail PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31
	// and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float
	// type are safe for storing this value.
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a story.
type Story struct {
	// Chat that posted the story
	Chat *Chat `json:"chat"`

	// Unique identifier for the story in the chat
	ID int `json:"id"`
}

// This object represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width as defined by the sender
	Width int `json:"width"`

	// Video height as defined by the sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Available sizes of the cover of the video in the message
	Cover []PhotoSize `json:"cover,omitempty"`

	// Optional. Timestamp in seconds from which the video will play in the message
	StartTimestamp int `json:"start_timestamp,omitempty"`

	// Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some
	// programming languages may have difficulty/silent defects in interpreting
	// it. But it has at most 52 significant bits, so a signed 64-bit integer
	// or double-precision float type are safe for storing this value.
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by the sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`

	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some
	// programming languages may have difficulty/silent defects in interpreting
	// it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this value.
	FileSize int `json:"file_size,omitempty"`
}
