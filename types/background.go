package types

// This object describes the way a background
// is filled based on the selected colors.
type BackgroundFill struct {
	// BackgroundFillSolid
	// BackgroundFillGradient
	// BackgroundFillFreeformGradient
}

// The background is filled using the selected color.
type BackgroundFillSolid struct {
	// Type of the background fill, always “solid”
	Type string `json:"type"`

	// The color of the background fill in the RGB24 format
	Color int `json:"color"`
}

// The background is a gradient fill.
type BackgroundFillGradient struct {
	// Type of the background fill, always “gradient”
	Type string `json:"type"`

	// Top color of the gradient in the RGB24 format
	TopColor int `json:"top_color"`

	// Bottom color of the gradient in the RGB24 format
	BottomColor int `json:"bottom_color"`

	// Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int `json:"rotation_angle"`
}

// The background is a freeform gradient that
// rotates after every message in the chat.
type BackgroundFillFreeformGradient struct {
	// Type of the background fill, always “freeform_gradient”
	Type string `json:"type"`

	// A list of the 3 or 4 base colors that are used to
	// generate the freeform gradient in the RGB24 format
	Colors []int `json:"colors"`
}

// This object describes the type of a
// background. Currently, it can be one of
type BackgroundType struct {
	// BackgroundTypeFill
	// BackgroundTypeWallpaper
	// BackgroundTypePattern
	// BackgroundTypeChatTheme
}

// The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	// Type of the background, always “fill”
	Type string `json:"type"`

	// The background fill
	Fill BackgroundFill `json:"fill"`

	// Dimming of the background in dark
	// themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
}

// The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	// Type of the background, always “wallpaper”
	Type string `json:"type"`

	// Document with the wallpaper
	Document Document `json:"document"`

	// Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`

	// Optional. True, if the wallpaper is downscaled to fit
	// in a 450x450 square and then box-blurred with radius 12
	IsBlurred bool `json:"is_blurred,omitempty"`

	// Optional. True, if the background moves
	// slightly when the device is tilted
	IsMoving bool `json:"is_moving,omitempty"`
}

// The background is a .PNG or .TGV (gzipped subset of SVG
// with MIME type “application/x-tgwallpattern”) pattern to
// be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	// Type of the background, always “pattern”
	Type string `json:"type"`

	// Document with the pattern
	Document Document `json:"document"`

	// The background fill that is combined with the pattern
	Fill BackgroundFill `json:"fill"`

	// Intensity of the pattern when it is shown
	// above the filled background; 0-100
	Intensity int `json:"intensity"`

	// Optional. True, if the background fill must be
	// applied only to the pattern itself. All other
	// pixels are black in this case. For dark themes only
	IsInverted bool `json:"is_inverted,omitempty"`

	// Optional. True, if the background moves
	// slightly when the device is tilted
	IsMoving bool `json:"is_moving,omitempty"`
}

// The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	// Type of the background, always “chat_theme”
	Type string `json:"type"`

	// Name of the chat theme, which is usually an emoji
	ThemeName string `json:"theme_name"`
}
