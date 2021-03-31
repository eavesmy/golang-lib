package veri

import "regexp"

func IsEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func IsMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func IsYYYYMMDD(layout string) bool {
	pattern := `\d{4}-\d{2}-\d{2}`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(layout)
}

func IsIdCard(idCard string) bool {
	res, _ := regexp.Match("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", []byte(idCard))
	return res
}

func IsChinaName(name string) bool {
	res, _ := regexp.Match("^[\u4E00-\u9FA5]{2,4}$", []byte(name))
	return res
}

// 支持中英数字及下划线和小数点
func IsNickname(name string) bool {
	res, _ := regexp.Match("^[\\.a-zA-Z0-9_\u4e00-\u9fa5]+$", []byte(name))
	return res
}
