package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/nongdenchet/covidform/model"
)

func DecodeGenerateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(context context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(context, w, response)
}
