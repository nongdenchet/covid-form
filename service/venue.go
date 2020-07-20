package service

import (
	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/repository"
	"github.com/nongdenchet/covidform/utils"
)

// region dto

type VenueData struct {
	Token string   `json:"token"`
	Venue VenueDTO `json:"venue"`
}

type VenueDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"addeess"`
	Logo    *string `json:"logo"`
}

func venueToDTO(v *model.Venue) VenueDTO {
	return VenueDTO{
		ID:      v.ID.String(),
		Name:    v.Name,
		Address: v.Address,
		Logo:    v.Logo,
	}
}

// endregion

type VenueService interface {
	Register(RegisterRequest) (*RegisterResponse, error)
	Login(LoginRequest) (*LoginResponse, error)
	GetVenue(string) (*GetVenueResponse, error)
	UpdateVenue(*model.Venue, UpdateVenueRequest) (*VenueDTO, error)
}

type venueServiceImpl struct {
	Repo repository.VenueRepo
}

func NewVenueService(repo repository.VenueRepo) VenueService {
	return venueServiceImpl{Repo: repo}
}

// region get venue

type GetVenueResponse struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Logo    *string `json:"logo"`
}

func (s venueServiceImpl) GetVenue(id string) (*GetVenueResponse, error) {
	v, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, utils.NewNotFoundError("venue is not found")
	}

	return &GetVenueResponse{
		Name:    v.Name,
		Address: v.Address,
		Logo:    v.Logo,
	}, nil
}

// endregion

// region login

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string    `json:"status"`
	Data   VenueData `json:"data"`
}

func (s venueServiceImpl) Login(r LoginRequest) (*LoginResponse, error) {
	respError := utils.NewUserError("Username or password is wrong")

	// User
	v, err := s.Repo.GetByUsername(r.Username)
	if err != nil {
		return nil, err
	}

	// Missing user
	if v == nil {
		return nil, respError
	}

	// Check password
	if !utils.ComparePasswords(v.PasswordHash, r.Password) {
		return nil, respError
	}

	// Generate token
	token, err := utils.GenerateToken(r.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Status: "successful",
		Data: VenueData{
			Token: token,
			Venue: venueToDTO(v),
		},
	}, nil
}

// endregion

// region register

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Confirm  string `json:"passwordConfirm"`
}

type RegisterResponse struct {
	Status string    `json:"status"`
	Data   VenueData `json:"data"`
}

func (s venueServiceImpl) Register(r RegisterRequest) (*RegisterResponse, error) {
	if len(r.Username) < 3 {
		return nil, utils.NewUserError("Username must be at least 3 characters")
	}
	if len(r.Password) < 6 {
		return nil, utils.NewUserError("Password must be at least 6 characters")
	}
	if r.Password != r.Confirm {
		return nil, utils.NewUserError("Password mismatch")
	}

	current, err := s.Repo.GetByUsername(r.Username)
	if err != nil {
		return nil, err
	}
	if current != nil {
		return nil, utils.NewUserError("Username already used")
	}

	ph := utils.HashAndSalt(r.Password)
	v, err := s.Repo.Create(r.Username, ph)
	if err != nil {
		return nil, err
	}

	// Generate token
	token, err := utils.GenerateToken(v.Username)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		Status: "successful",
		Data: VenueData{
			Token: token,
			Venue: venueToDTO(v),
		},
	}, nil
}

// endregion

// region update venue info

type UpdateVenueRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (s venueServiceImpl) UpdateVenue(
	v *model.Venue,
	r UpdateVenueRequest,
) (*VenueDTO, error) {
	result, err := s.Repo.Update(v.ID.String(), r.Name, r.Address)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, utils.NewNotFoundError("venue not found")
	}

	dto := venueToDTO(result)
	return &dto, nil
}

// end region
