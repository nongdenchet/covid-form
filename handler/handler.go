package handler

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/nongdenchet/covidform/endpoint"
	"github.com/nongdenchet/covidform/helpers"
	"github.com/nongdenchet/covidform/middleware"
	"github.com/nongdenchet/covidform/repository"
	"github.com/nongdenchet/covidform/service"
)

func NewHandler() *httptransport.Server {
	repo := repository.IdenticonRepoImpl{}

	var s endpoint.IdenticontService
	{
		s = service.IdenticonServiceImpl{Repo: repo}
		s = middleware.NewLoggingMiddleware(s)
		s = middleware.NewIntrumentationMiddleware(s)
	}

	return httptransport.NewServer(
		endpoint.MakeGenerateEndpoint(s),
		helpers.DecodeGenerateRequest,
		helpers.EncodeResponse,
	)
}
