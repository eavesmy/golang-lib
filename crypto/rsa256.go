package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Rsa struct {
	pubKey []byte
	priKey []byte
}

func (r *Rsa) SetKey(pubKey, priKey []byte) {
	r.pubKey = pubKey
	r.priKey = priKey
}

// 加密
func (r *Rsa) RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(r.pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func (r *Rsa) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(r.priKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
