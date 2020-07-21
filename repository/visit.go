package repository

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nongdenchet/covidform/model"
)

type VisitRepo interface {
	Create(venueID string, name string, phone string, email *string) (*model.Visit, error)
	GetByVenue(venueID string, date time.Time) ([]model.Visit, error)
}

type visitRepoImpl struct {
	db *gorm.DB
}

func NewVisitRepo(db *gorm.DB) VisitRepo {
	return visitRepoImpl{db}
}

func (r visitRepoImpl) Create(
	venueID string,
	name string,
	phone string,
	email *string,
) (*model.Visit, error) {
	v := &model.Visit{VenueID: venueID, Name: name, Phone: phone, Email: email}
	err := r.db.Create(v).Error
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r visitRepoImpl) GetByVenue(venueID string, date time.Time) ([]model.Visit, error) {
	v := make([]model.Visit, 0)

	today := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	tomorrow := today.AddDate(0, 0, 1)
	log.Printf("today: %+v", today)
	log.Printf("tomorrow: %+v", tomorrow)

	err := r.db.Where("venue_id = ?", venueID).
		Where("created_at between ? and ?", today.UTC(), tomorrow.UTC()).
		Find(&v).
		Error

	return v, err
}
