package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/service"
	"github.com/nongdenchet/covidform/utils"
)

func (h handlerImpl) SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	venueID := mux.Vars(r)["id"]

	var request service.SubmitFormRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "can't parse request")
		return
	}

	res, err := h.visitService.SubmitForm(venueID, request)
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

func (h handlerImpl) GetVisitsByVenueHandler(w http.ResponseWriter, r *http.Request) {
	v := context.Get(r, utils.Venue).(*model.Venue)

	// Parse date
	date, err := time.Parse(utils.DateFormat, r.URL.Query().Get("date"))
	if err != nil {
		log.Printf("%+v", err)
		respondWithError(w, http.StatusBadRequest, "can't parse date")
		return
	}

	res, err := h.visitService.GetByVenue(v, date)
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
