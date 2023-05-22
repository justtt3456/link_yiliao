package middleware

import (
	"china-russia/extends"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  lang.Lang("User is not logged in"),
			})
			c.Abort()
			return
		}
		jwtService := extends.JwtUtils{}
		parseToken := jwtService.ParseToken(token)
		if parseToken == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  lang.Lang("User is not logged in"),
			})
			c.Abort()
			return
		}
		m := model.Member{Id: parseToken.Id}
		if !m.Get() {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  lang.Lang("User is not logged in"),
			})
			c.Abort()
			return
		}
		if m.Token != parseToken.Key {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  lang.Lang("User is not logged in"),
			})
			c.Abort()
			return
		}

		if c.Request.Method == http.MethodPost {
			key := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(m.Id)
			value := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(m.Id)
			if global.REDIS.Get(key).Val() == value {
				c.JSON(http.StatusOK, gin.H{
					"code": 10001,
					"msg":  "请勿重复提交表单",
				})
				c.Abort()
				return
			}
			global.REDIS.Set(key, value, time.Second*3)
		}

		c.Set("member", m)
	}
}
