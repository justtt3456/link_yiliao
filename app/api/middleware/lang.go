package middleware

import (
	"china-russia/global"
	"github.com/gin-gonic/gin"
)

var langMap map[string]bool = map[string]bool{
	"zh_cn": true,
	"zh_hk": true,
	"en":    true,
	"ja_JP": true,
	"vi_VN": true,
}

func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.Header.Get("lang")
		if lang == "" {
			global.Language = "zh_cn"
			return
		}
		if _, ok := langMap[lang]; !ok {
			global.Language = "zh_cn"
			return
		}
		global.Language = lang
	}
}
