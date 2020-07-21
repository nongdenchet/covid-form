package service

import (
	"log"
	"time"

	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/repository"
	"github.com/nongdenchet/covidform/utils"
)

type VisitService interface {
	SubmitForm(venueID string, r SubmitFormRequest) (*SubmitFormResponse, error)
	GetByVenue(*model.Venue, time.Time) (*GetVisitsResponse, error)
}

type visitServiceImpl struct {
	Repo repository.VisitRepo
}

func NewVisitService(repo repository.VisitRepo) VisitService {
	return visitServiceImpl{Repo: repo}
}

type VisitDTO struct {
	Name  string  `json:"name"`
	Phone string  `json:"phone"`
	Email *string `json:"email"`
	Date  string  `json:"date"`
}

// region submit

type SubmitFormRequest struct {
	Name  string  `json:"name"`
	Phone string  `json:"phone"`
	Email *string `json:"email"`
}

type SubmitFormResponse struct {
	Status string   `json:"status"`
	Data   VisitDTO `json:"data"`
}

func (s visitServiceImpl) SubmitForm(venueID string, r SubmitFormRequest) (*SubmitFormResponse, error) {
	if len(r.Name) == 0 {
		return nil, utils.NewUserError("Username must be at least 3 characters")
	}
	if len(r.Phone) == 0 {
		return nil, utils.NewUserError("Password must be at least 6 characters")
	}

	v, err := s.Repo.Create(venueID, r.Name, r.Phone, r.Email)
	if err != nil {
		return nil, err
	}

	return &SubmitFormResponse{
		Status: "successful",
		Data: VisitDTO{
			Name:  v.Name,
			Phone: v.Phone,
			Email: v.Email,
			Date:  v.CreatedAt.Format(utils.DateFormat),
		},
	}, nil
}

// endregion

// region get by venue

type GetVisitsResponse struct {
	Results []VisitDTO `json:"results"`
}

func (s visitServiceImpl) GetByVenue(venue *model.Venue, date time.Time) (*GetVisitsResponse, error) {
	log.Printf("date: %+v", date)
	venues, err := s.Repo.GetByVenue(venue.ID.String(), date)
	if err != nil {
		return nil, err
	}

	dto := make([]VisitDTO, 0)
	for _, v := range venues {
		dto = append(dto, VisitDTO{
			Name:  v.Name,
			Phone: v.Phone,
			Email: v.Email,
			Date:  v.CreatedAt.Format(utils.DateFormat),
		})
	}

	return &GetVisitsResponse{Results: dto}, nil
}

// endregion
