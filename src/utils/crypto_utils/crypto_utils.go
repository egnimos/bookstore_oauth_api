package crypto_utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetSha512(input string) string {
	hash := sha512.New512_256()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
