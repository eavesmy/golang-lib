package auth
import (
	"encoding/json"
	"fmt"
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

type Endata struct {
	Uid        string
	timed      string
	RandString string
}

func GenToken(uid string) string {

	t := gtype.Int642String(time.Now().Unix())

	sign_data := Endata{uid, t, "rand_data"}
	en_data, _ := json.Marshal(sign_data)
	token := crypto.Base64EnCrypetcode(KEY, string(en_data))
	return token
	// return crypto.Rc4(token, KEY)
}

func ParseToken(token string) (uid, t, sign string) {

	str, err := crypto.Base64DeCryptcode(KEY, token)
	if err != nil {
		return
	}
	data := new(Endata)
	err = json.Unmarshal([]byte(str), &data)
	if err != nil {
		fmt.Println("token Error ", err)
		return
	}
	if data.Uid == "" {
		return
	}

	uid = data.Uid
	t = data.timed
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
