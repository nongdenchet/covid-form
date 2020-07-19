package service

import (
	"github.com/nongdenchet/covidform/repository"
	"github.com/nongdenchet/covidform/utils"
)

type VenueService interface {
	Register(RegisterRequest) (*RegisterResponse, error)
	Login(LoginRequest) (*LoginResponse, error)
}

type venueServiceImpl struct {
	Repo repository.VenueRepo
}

func NewVenueService(repo repository.VenueRepo) VenueService {
	return venueServiceImpl{Repo: repo}
}

// region login

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s venueServiceImpl) Login(r LoginRequest) (*LoginResponse, error) {
	respError := utils.NewUserError("Username or password is wrong")

	current, err := s.Repo.GetByUsername(r.Username)
	if err != nil {
		return nil, err
	}

	// Missing user
	if current == nil {
		return nil, respError
	}

	// Check password
	if !utils.ComparePasswords(current.PasswordHash, r.Password) {
		return nil, respError
	}

	// Generate token
	token, err := utils.GenerateToken(r.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token}, nil
}

// endregion

// region register

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Confirm  string `json:"passwordConfirm"`
}

type RegisterResponse struct {
	Status string `json:"status"`
}

func (s venueServiceImpl) Register(r RegisterRequest) (*RegisterResponse, error) {
	if r.Password != r.Confirm {
		return nil, utils.NewUserError("Password mismatch")
	}
	if len(r.Username) < 3 {
		return nil, utils.NewUserError("Username must be at least 3 characters")
	}
	if len(r.Password) < 6 {
		return nil, utils.NewUserError("Password must be at least 6 characters")
	}

	current, err := s.Repo.GetByUsername(r.Username)
	if err != nil {
		return nil, err
	}
	if current != nil {
		return nil, utils.NewUserError("Username already used")
	}

	ph := utils.HashAndSalt(r.Password)
	_, err = s.Repo.Create(r.Username, ph)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{Status: "success"}, nil
}

// endregion
