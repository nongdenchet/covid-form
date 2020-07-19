package handler

import (
	"net/http"

	"github.com/gorilla/context"

	"github.com/nongdenchet/covidform/repository"
	"github.com/nongdenchet/covidform/utils"
)

type AuthenMiddleware struct {
	Repo repository.VenueRepo
}

func (am AuthenMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(utils.AuthenticationHeader)
		claims, valid := utils.ValidateToken(token)
		if !valid {
			respondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		venue, err := am.Repo.GetByUsername(claims.Username)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "internal error")
			return
		}
		if venue == nil {
			respondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		context.Set(r, utils.Venue, venue)
		next.ServeHTTP(w, r)
	})
}
