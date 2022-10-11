package middleware

import (
	"encoding/json"
	"finance/global"
	"finance/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func WhiteIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ips := map[string]bool{}
		ip := c.ClientIP()
		im := model.IP{}
		str := global.REDIS.Get(model.StringKeyWhiteIP).Val()
		if str == "" {
			if im.Get() {
				sp := strings.Split(im.IP, "/")
				for _, v := range sp {
					ips[v] = true
				}
				bytes, _ := json.Marshal(im)
				global.REDIS.Set(model.StringKeyWhiteIP, string(bytes), -1)
			}
		} else {
			err := json.Unmarshal([]byte(str), &im)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 10001,
					"msg":  fmt.Sprintf("IP%s不在白名单", ip),
				})
				c.Abort()
				return
			}
			sp := strings.Split(im.IP, "/")
			for _, v := range sp {
				ips[v] = true
			}
		}
		if !ips[ip] {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  fmt.Sprintf("IP%s不在白名单", ip),
			})
			c.Abort()
			return
		}
	}
}
