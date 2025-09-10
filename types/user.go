package types

type User struct {
	// Unique identifier for this user or bot. This number
	// may have more than 32 significant bits and some
	// programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	ID int `json:"id"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`

	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`

	// Optional. True, if the bot can be connected to a Telegram Business
	// account to receive its messages. Returned only in getMe.
	CanConnectToBusiness bool `json:"can_connect_to_business,omitempty"`

	// Optional. True, if the bot has a main Web App. Returned only in getMe.
	HasMainWebApp bool `json:"has_main_web_app,omitempty"`
}

// This object represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Contact's user identifier in Telegram.
	// This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	UserID int `json:"user_id,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard)
	Vcard string `json:"vcard,omitempty"`
}

type Location struct {
	// Latitude as defined by the sender
	Latitude float64 `json:"latitude"`

	// Longitude as defined by the sender
	Longitude float64 `json:"longitude"`

	// Optional. The radius of uncertainty for the location,
	// measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// Optional. Time relative to the message sending date,
	// during which the location can be updated; in seconds.
	// For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`

	// Optional. The direction in which user is moving,
	// in degrees; 1-360. For active live locations only.
	Heading int `json:"heading,omitempty"`

	// Optional. The maximum distance for proximity alerts
	// about approaching another chat member, in meters.
	// For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// This object represents a venue.
type Venue struct {
	// Venue location. Can't be a live location
	Location *Location `json:"location"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium”
	// or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// Optional. Google Places type of the venue. (See supported types.)
	// (https://developers.google.com/maps/documentation/places/web-service/legacy/supported_types)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}
