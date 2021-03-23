package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(b []byte) (sign string) {
	h := sha256.New()
	h.Write(b)
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
