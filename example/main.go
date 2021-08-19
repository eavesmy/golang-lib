package main

import (
	"fmt"
	"github.com/eavesmy/golang-lib/auth"
)

func main ()  {

	data:= auth.GenToken("1111")

	aa,_,_ := auth.ParseToken(data)

	fmt.Println(data)
	fmt.Println(aa)


}