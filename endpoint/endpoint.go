package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nongdenchet/covidform/model"
)

type IdenticontService interface {
	Generate(context.Context, string, int) (string, error)
}

func MakeGenerateEndpoint(svc IdenticontService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GenerateRequest)
		url, err := svc.Generate(ctx, req.Text, req.Size)
		if err != nil {
			return model.GenerateResponse{Err: err.Error()}, nil
		}

		return model.GenerateResponse{Url: url}, nil
	}
}
