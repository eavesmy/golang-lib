package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/fwhezfwhez/errorx"
)

func Sha256WithRsa(privateRaw string, msg string) (string, error) {
	privateRaw = strings.Trim(privateRaw, "\n")
	if !strings.HasPrefix(privateRaw, "-----BEGIN RSA PRIVATE KEY-----") {
		privateRaw = fmt.Sprintf("%s\n%s\n%s", "-----BEGIN RSA PRIVATE KEY-----", privateRaw, "-----END RSA PRIVATE KEY-----")
	}

	blockPri, _ := pem.Decode([]byte(privateRaw))
	if blockPri == nil {
		return "", fmt.Errorf("blockPri is nil")
	}

	rsaPri, e := genPriKey(blockPri.Bytes, PKCS8)
	if e != nil {
		panic(e)
	}

	h := sha256.New()
	h.Write([]byte(msg))
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPri, crypto.SHA256, d)
	if err != nil {
		return "", errorx.Wrap(err)
	}
	encodedSig := base64.StdEncoding.EncodeToString(signature)
	return encodedSig, nil
}

const (
	PKCS1 int64 = iota
	PKCS8
)

func genPriKey(privateKey []byte, privateKeyType int64) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, fmt.Errorf("unsupport private key type")
		}
	}
	return priKey, nil
}
