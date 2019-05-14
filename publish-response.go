package model

// PublishResponse response from the rpm server for a publish event
type PublishResponse struct {
	Message string `json:"message,omitempty"`
}
