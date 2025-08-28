package types

type ResponseType interface {
	Update
}

type APIResponse[T ResponseType] struct {
	Ok          bool               `json:"ok"`
	Result      []T                `json:"result"`
	Description string             `json:"description"`
	ErrorCode   int                `json:"error_code"`
	Parameters  responseParameters `json:"parameters"`
}

type responseParameters struct {
	// Optional. The group has been migrated to a supergroup
	// with the specified identifier. This number may have more
	// than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed
	// 64-bit integer or double-precision float type are safe
	// for storing this identifier.
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`

	// Optional. In case of exceeding flood control, the number of
	// seconds left to wait before the request can be repeated
	Retry_After int `json:"retry_after,omitempty"`
}
