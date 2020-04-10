package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type Rsa struct {
	pubKey []byte
	priKey []byte
}

func (rsa *Rsa) SetKey(pubKey, priKey []byte) {
	rsa.pubKey = pubKey
	rsa.priKey = priKey
}

// 加密
func (rsa *Rsa) RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(rsa.pubKey)
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
func (rsa *Rsa) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(rsa.priKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
