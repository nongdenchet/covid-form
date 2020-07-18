package encode

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) ([]byte, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return nil, err
	}

	hash := hasher.Sum(nil)
	size := hex.EncodedLen(len(hash))
	result := make([]byte, size)
	hex.Encode(result, hash)
	return result, nil
}
