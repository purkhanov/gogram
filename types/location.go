package types

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
