package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(data string) string {
	return _sha1([]byte(data))
}

func Sha1Bytes(data []byte) string {
	return _sha1(data)
}

func _sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte("")))
}
