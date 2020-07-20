package model

type Venue struct {
	Base
	Username     string
	PasswordHash string
	Name         string
	Address      string
	Logo         *string
}
