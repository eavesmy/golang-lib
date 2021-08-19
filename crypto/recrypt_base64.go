/*
 * @Author: jjc
 * @Date: 2021-06-01 12:53:39
 * @LastEditTime: 2021-06-01 12:59:01
 * @LastEditors: Please set LastEditors
 * @Description:
 * @FilePath: /tianshen/root/go/pkg/mod/github.com/eavesmy/golang-lib@v0.3.5/crypto/recrypt_base64.go
 */
package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
)

const (
	//BASE64字符表,不要有重复
	base64Table = "<>:;9,7234!@#$CDVWX%5&*ABYZabcghijklmnopqrstuvwxyz01EFGHIJKLMNO="
	//	hashFunctionHeader = "tianshen.eavesemy.hhh"
	hashFunctionFooter = "09.O25.O20.78"
)

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s + "zhifeiya")) //使用zhifeiya名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 获取一个Guid值
 */
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

var coder = base64.NewEncoding(base64Table)

/**
 * base64加密
 */
func Base64EnCrypetcode(hashFunctionHeader string, str string) string {
	var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

/**
 * base64解密
 */
func Base64DeCryptcode(hashFunctionHeader string, str string) (string, error) {
	var src []byte = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}
