# Golang 开发常用函数库

# 使用
```golang
package main

import "github.com/eavesmy/golang-lib/crypto"
import "fmt"

func main(){
	word := crypto.Md5("Hello world")
	fmt.Println(word) // 32 length string.
}
```
