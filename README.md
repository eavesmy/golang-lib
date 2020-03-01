# Golang 开发常用函数库

# 功能(未标注作者即为自己开发)

### csv golang 操作 csv 文件
### auth 一个最小化的登陆、注册、session、cookie 管理组件。
### crypto 包含 rc4、md5 等加解密的封装

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
