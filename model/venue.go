package model

type Venue struct {
	Base
	Username     string
	PasswordHash string
	Logo         *string
}
