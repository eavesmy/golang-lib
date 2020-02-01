package crypto

import "crypto/rc4"

func Rc4(txt string, key string) string {

	if txt == "" || key == "" {
		return ""
	}

	src := []byte(txt)
	k := []byte(key)

	c, _ := rc4.NewCipher(k)
	r := make([]byte, len(src))

	c.XORKeyStream(r, src)

	defer c.Reset()

	return string(r)
}
