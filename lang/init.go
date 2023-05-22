package lang

import (
	"china-russia/global"
)

func Lang(msg string) string {
	switch global.Language {
	case "zh_cn":
		v, ok := zh_cn[msg]
		if !ok {
			return msg
		}
		return v
	case "zh_hk":
		v, ok := zh_hk[msg]
		if !ok {
			return msg
		}
		return v
	case "ja_JP":
		v, ok := ja_JP[msg]
		if !ok {
			return msg
		}
		return v
	case "vi_VN":
		v, ok := vi_VN[msg]
		if !ok {
			return msg
		}
		return v
	}
	return msg
}
