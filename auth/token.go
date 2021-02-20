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

	sign := ""
	sign += "uid=" + uid
	sign += "&"
	sign += "t=" + t

	sign = crypto.Rc4(sign, uid)

	token := uid + "|" + t + "|" + sign

	return crypto.Rc4(token, KEY)
}

func ParseToken(token string) (uid, t, sign string) {

	token = crypto.Rc4(token, KEY)

	arr := strings.Split(token, "|")

	if len(arr) < 3 {
		return
	}

	uid = arr[0]
	t = arr[1]
	sign = arr[2]

	return
}

func VeriSign(uid, t, sign string) bool {

	str := crypto.Rc4(sign, uid)

	arr := strings.Split(str, "&")
	if len(arr) < 2 {
		return false
	}

	for _, item := range arr {
		param := strings.Split(item, "=")
		k := param[0]
		v := param[1]

		if k == "uid" && v != uid {
			return false
		}

		if k == "t" && v != t {
			return false
		}
	}

	return true
}
