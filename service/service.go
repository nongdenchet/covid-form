package service

import (
	"context"
	"errors"

	"github.com/nongdenchet/covidform/encode"
	"github.com/nongdenchet/covidform/repository"
)

type IdenticonServiceImpl struct {
	Repo repository.IdenticonRepo
}

func (s IdenticonServiceImpl) Generate(_ context.Context, text string, size int) (string, error) {
	if len(text) == 0 {
		return "", errors.New("text should not empty")
	}

	if size <= 0 {
		return "", errors.New("size should not empty and must be larger than 0")
	}

	hash, err := encode.GetMD5Hash(text)
	if err != nil {
		return "", err
	}

	filePath, err := s.Repo.GetIdenticon(hash, size)
	if err != nil {
		return "", err
	}

	if len(filePath) > 0 {
		return filePath, nil
	}

	return s.Repo.CreateIdenticon(hash, size)
}
