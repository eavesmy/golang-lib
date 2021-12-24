package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var KEY = "zhongnanhai"

/*
rc4(key,user_id + "|" + timestamp + "|" + sign)
key: "zhongnanhai"
sign: rc4(user_id,uid=user_id&t=timestamp)

能做到基础的防止伪造token。每次请求都会更换token，重复的token 不能被使用。

*/

type Endata struct {
	Uid        string
	timed      string
	RandString string
}

func GenToken(uid string) string {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	token, err := at.SignedString([]byte(KEY))
	if err != nil {
		return ""
	}
	return token
}

func ParseToken(token string) (uid, t, sign string) {

	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		fmt.Println("token->err", err)
		return
	}
	uid = claim.Claims.(jwt.MapClaims)["uid"].(string)
	t = ""
	sign = GenToken(uid)
	return
}

func VeriSign(uid, t, sign string) bool {
	uid_temp, _, _ := ParseToken(sign)
	if uid_temp != uid {
		return false
	}
	return true
}
