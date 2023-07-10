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
	regular := "^(13|14|15|16|17|18|19)\\d{9}$"
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
	regRuler := "(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
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
