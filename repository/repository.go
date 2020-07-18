package repository

import (
	"image/png"
	"os"
	"path/filepath"

	"github.com/nongdenchet/covidform/generator"
	"github.com/nongdenchet/covidform/helpers"
)

type IdenticonRepo interface {
	GetIdenticon([]byte, int) (string, error)
	CreateIdenticon([]byte, int) (string, error)
}

type IdenticonRepoImpl struct{}

func (IdenticonRepoImpl) GetIdenticon(hash []byte, size int) (string, error) {
	filePath := helpers.GetIdenticonFilePath(hash, size)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", nil
	}

	return filepath.Abs(filePath)
}

func (IdenticonRepoImpl) CreateIdenticon(hash []byte, size int) (string, error) {
	filePath := helpers.GetIdenticonFilePath(hash, size)
	file, err := helpers.PrepareFile(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	icon := generator.GenerateIcon(hash, size)
	err = png.Encode(file, icon)
	if err != nil {
		return "", err
	}

	return filepath.Abs(filePath)
}
