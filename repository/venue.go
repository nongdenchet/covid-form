package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nongdenchet/covidform/model"
)

type VenueRepo interface {
	Create(username string, passwordHash string) (*model.Venue, error)
	GetByUsername(username string) (*model.Venue, error)
}

type venueRepoImpl struct {
	db *gorm.DB
}

func (v venueRepoImpl) Create(username string, passwordHash string) (*model.Venue, error) {
	user := &model.Venue{Username: username, PasswordHash: passwordHash}
	err := v.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (v venueRepoImpl) GetByUsername(username string) (*model.Venue, error) {
	var venue model.Venue
	if result := v.db.Where("username = ?", username).First(&venue); result.Error != nil {
		if result.RecordNotFound() {
			return nil, nil
		}

		return nil, result.Error
	}

	return &venue, nil
}

func NewVenueRepo(db *gorm.DB) VenueRepo {
	return venueRepoImpl{db}
}
