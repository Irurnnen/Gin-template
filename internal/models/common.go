package models

type Message struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
