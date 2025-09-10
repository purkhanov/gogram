package types

/*
Telegram Passport is a unified authorization method for services
that require personal identification. Users can upload their documents
once, then instantly share their data with services that require
real-world ID (finance, ICOs, etc.). Please see the manual for details.
*/

// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Array with information about documents and other
	// Telegram Passport elements that was shared with the bot
	Data []EncryptedPassportElement `json:"data"`

	// Encrypted credentials required to decrypt the data
	Credentials EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram
// Passport. Currently all Telegram Passport files are
// in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can
	// be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed
	// to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// File size in bytes
	FileSize int `json:"file_size"`

	// Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

// Describes documents or other Telegram Passport
// elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of “personal_details”, “passport”,
	// “driver_license”, “identity_card”, “internal_passport”,
	// “address”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”,
	// “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`

	// Optional. Base64-encoded encrypted Telegram Passport
	// element data provided by the user; available only for
	// “personal_details”, “passport”, “driver_license”,
	// “identity_card”, “internal_passport” and “address”
	// types. Can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	Data string `json:"data"`

	// Optional. User's verified phone number;
	// available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`

	// Optional. User's verified email address;
	// available only for “email” type
	Email string `json:"email,omitempty"`

	// Optional. Array of encrypted files with documents
	// provided by the user; available only for “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”
	// and “temporary_registration” types. Files can be decrypted
	// and verified using the accompanying EncryptedCredentials.
	Files []PassportFile `json:"files,omitempty"`

	// Optional. Encrypted file with the front side of the document,
	// provided by the user; available only for “passport”,
	// “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	FrontSide PassportFile `json:"front_side,omitempty"`

	// Optional. Encrypted file with the reverse side of the document,
	// provided by the user; available only for “driver_license” and
	// “identity_card”. The file can be decrypted and verified using
	// the accompanying EncryptedCredentials.
	ReverseSide PassportFile `json:"reverse_side,omitempty"`

	// Optional. Encrypted file with the selfie of the user holding a
	// document, provided by the user; available if requested for
	// “passport”, “driver_license”, “identity_card” and
	// “internal_passport”. The file can be decrypted and verified
	// using the accompanying EncryptedCredentials.
	Selfie PassportFile `json:"selfie,omitempty"`

	// Optional. Array of encrypted files with translated versions of
	// documents provided by the user; available if requested for
	// “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and
	// verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation,omitempty"`

	// Base64-encoded element hash for using in
	// PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// Describes data required for decrypting and authenticating
// EncryptedPassportElement. See the Telegram Passport
// Documentation for a complete description of the data
// decryption and authentication processes.
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique
	// user's payload, data hashes and secrets required for
	// EncryptedPassportElement decryption and authentication
	Data string `json:"data"`

	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`

	// Base64-encoded secret, encrypted with the bot's
	// public RSA key, required for data decryption
	Secret string `json:"secret"`
}
