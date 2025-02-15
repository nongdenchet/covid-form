package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/service"
	"github.com/nongdenchet/covidform/utils"
)

func (h handlerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request service.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't parse request")
		return
	}

	res, err := h.venueService.Login(request)
	if err != nil {
		log.Printf("%+v", err)
		switch v := err.(type) {
		default:
			respondWithError(w, http.StatusInternalServerError, "internal error")
		case utils.UserError:
			respondWithError(w, http.StatusBadRequest, v.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (h handlerImpl) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request service.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't parse request")
		return
	}

	res, err := h.venueService.Register(request)
	if err != nil {
		log.Printf("%+v", err)
		switch v := err.(type) {
		default:
			respondWithError(w, http.StatusInternalServerError, "internal error")
		case utils.UserError:
			respondWithError(w, http.StatusBadRequest, v.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (h handlerImpl) GetVenueHandler(w http.ResponseWriter, r *http.Request) {
	venueID := mux.Vars(r)["id"]
	res, err := h.venueService.GetVenue(venueID)
	if err != nil {
		log.Printf("%+v", err)
		switch v := err.(type) {
		default:
			respondWithError(w, http.StatusInternalServerError, "internal error")
		case utils.NotFoundError:
			respondWithError(w, http.StatusNotFound, v.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (h handlerImpl) UpdateVenueHandler(w http.ResponseWriter, r *http.Request) {
	v := context.Get(r, utils.Venue).(*model.Venue)

	var request service.UpdateVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't parse request")
		return
	}

	res, err := h.venueService.UpdateVenue(v, request)
	if err != nil {
		log.Printf("%+v", err)
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}
