package auth

import (
	"strings"
	"time"

	"github.com/eavesmy/golang-lib/crypto"
	gtype "github.com/eavesmy/golang-lib/type"
)

var KEY = "zhongnanhai"

/*
rc4(key,user_id + "|" + timestamp + "|" + sign)
key: "zhongnanhai"
sign: rc4(user_id,uid=user_id&t=timestamp)

能做到基础的防止伪造token。每次请求都会更换token，重复的token 不能被使用。

*/



func GenToken(uid string) string {

	t := gtype.Int642String(time.Now().Unix())
	sign := uid + "|" + t + "|" + "data_zhanwei"
	token := crypto.Base64EnCrypetcode(KEY, sign)

	return token
	// return crypto.Rc4(token, KEY)
}

func ParseToken(token string) (uid, t, sign string) {

	str,err := crypto.Base64DeCryptcode(KEY, token)
	if err!=nil {
		return
	}
	arr := strings.Split(str, "|")

	if len(arr) < 3 {
		return
	}

	uid = arr[0]
	t = arr[1]
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
