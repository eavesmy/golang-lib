package main

import (
	"fmt"
	"github.com/eavesmy/golang-lib/auth"
)

func main ()  {


	token := auth.GenToken("u9822")
	uid,t,sign := auth.ParseToken(token)

	fmt.Println(uid,t,sign )

}