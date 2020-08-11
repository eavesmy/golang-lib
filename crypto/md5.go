package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(txt string, crypted ...string) string {
	var c []byte
	if len(crypted) != 0 {
		c = []byte(crypted[0])
	}
	return md5_32(txt, c)
}

func Md5_16(txt string, crypted ...string) string {
	var c []byte
	if len(crypted) != 0 {
		c = []byte(crypted[0])
	}
	txt = md5_32(txt, c)
	ret := txt[0:16] + txt[17:32]
	return ret
}

func md5_32(txt string, crypted []byte) string {

	h := md5.New()
	h.Write([]byte(txt))
	cipherStr := h.Sum(crypted)

	return hex.EncodeToString(cipherStr)
}

func md5_32_bytes(txt []byte, crypted []byte) string {
	h := md5.New()
	h.Write(txt)
	cipherStr := h.Sum(crypted)

	return hex.EncodeToString(cipherStr)
}

func Md5Bytes(txt []byte, crypted ...[]byte) string {
	var c []byte
	if len(crypted) != 0 {
		c = crypted[0]
	}
	return md5_32_bytes(txt, c)
}
