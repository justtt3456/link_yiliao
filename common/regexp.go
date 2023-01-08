package common

import (
	"regexp"
	"strconv"
)

func IsMobile(mobile string, lang string) bool {
	if lang == "zh_cn" {
		return cnMobile(mobile)
	}
	return otherMobile(mobile)
}

func cnMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|195|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
func otherMobile(mobile string) bool {
	parseInt, err := strconv.ParseInt(mobile, 10, 64)
	if err != nil {
		return false
	}
	return parseInt > 0
}

func IsIdCard(card string) bool {
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(card)
}

func IsEmail(email string) bool {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	if result {
		return true
	} else {
		return false
	}
}
