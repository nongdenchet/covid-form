package helpers

import (
	"os"
	"strconv"
)

const dirPath = "output"

func GetIdenticonFilePath(hash []byte, size int) string {
	return dirPath + "/" + string(hash) + "-" + strconv.Itoa(size) + ".png"
}

func PrepareFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
}
