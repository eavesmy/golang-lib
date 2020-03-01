package crypto

import (
	"crypto/rc4"
	"encoding/hex"
)

func Rc4(txt string, key string) string {

	isEncrypt := false

	if txt == "" || key == "" {
		return ""
	}
	src, err := hex.DecodeString(txt)
	if err != nil {
		isEncrypt = true
		src = []byte(txt)
	}

	k := []byte(key)

	c, _ := rc4.NewCipher(k)
	r := make([]byte, len(src))

	c.XORKeyStream(r, src)

	defer c.Reset()

	if isEncrypt {
		return hex.EncodeToString(r)
	} else {
		return string(r)
	}

}
