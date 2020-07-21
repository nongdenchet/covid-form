package model

type Visit struct {
	Base
	VenueID string
	Venue   Venue `gorm:"foreignkey:VenueID"`
	Name    string
	Phone   string
	Email   *string
}
