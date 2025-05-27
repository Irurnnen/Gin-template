package models

// Message use for notification responses
type Message struct {
	Message string `json:"message"`
}

// HTTPError use for error responses
type HTTPError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
