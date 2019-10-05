package protocol

// UserFacingMessage error that gets sent back to user
type UserFacingMessage struct {
	Message    string
	StatusCode int
	SubError   error
}

func (e *UserFacingMessage) Error() string {
	return e.Message
}
