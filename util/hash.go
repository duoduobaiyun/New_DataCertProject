package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func Sha256HashString(data string) (string) {
	hashSHA256:=sha256.New()
	hashSHA256.Write([]byte(data))
	bytes := hashSHA256.Sum(nil)
	return hex.EncodeToString(bytes)
}

func Sha256HashReader(reader io.Reader) (string, error) {
	bytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "",err
	}
	sha56Hash := sha256.New()
	sha56Hash.Write(bytes)
	hashBytes := sha56Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}