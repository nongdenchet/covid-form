package handler

import (
	"net/http"

	"github.com/nongdenchet/covidform/utils"
)

func AuthenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(utils.AuthenticationHeader)

		_, valid := utils.ValidateToken(token)
		if valid {
			next.ServeHTTP(w, r)
		} else {
			respondWithError(w, http.StatusUnauthorized, "unauthorized")
		}
	})
}
