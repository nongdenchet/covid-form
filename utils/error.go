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

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) error {
	return NotFoundError{message}
}

func (e NotFoundError) Error() string {
	return e.message
}
