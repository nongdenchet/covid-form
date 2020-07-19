package utils

func NewUserError(message string) error {
	return UserError{message}
}

type UserError struct {
	message string
}

func (e UserError) Error() string {
	return e.message
}
