package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nongdenchet/covidform/repository"

	"github.com/jinzhu/gorm"
	"github.com/nongdenchet/covidform/service"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"status": "failed", "error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type Handler interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	GetVenueHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	UpdateVenueHandler(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	venueService service.VenueService
}

func NewHandler(db *gorm.DB) Handler {
	venueRepo := repository.NewVenueRepo(db)

	return handlerImpl{
		venueService: service.NewVenueService(venueRepo),
	}
}
