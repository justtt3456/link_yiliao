package middleware

import (
	"china-russia/extends"
	"china-russia/global"
	"china-russia/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		fmt.Println("token:", token)
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  "账号未登录",
			})
			c.Abort()
			return
		}
		jwtService := extends.JwtUtils{}
		parseToken := jwtService.ParseToken(token)
		if parseToken == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  "账号未登录",
			})
			c.Abort()
			return
		}
		res := global.REDIS.Get(fmt.Sprintf(model.StringKeyAdmin, parseToken.Id)).Val()
		u := model.Admin{}
		err := json.Unmarshal([]byte(res), &u)
		if err != nil {
			u.Id = parseToken.Id
			if !u.Get() {
				c.JSON(http.StatusOK, gin.H{
					"code": 10000,
					"msg":  "账号未登录",
				})
				c.Abort()
				return
			}
			if u.Token != parseToken.Key {
				c.JSON(http.StatusOK, gin.H{
					"code": 10000,
					"msg":  "账号未登录",
				})
				c.Abort()
				return
			}
			bytes, _ := json.Marshal(u)
			global.REDIS.Set(fmt.Sprintf(model.StringKeyAdmin, u.Id), string(bytes), u.ExpireTime())
		}
		if u.Token != parseToken.Key {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"msg":  "账号未登录",
			})
			c.Abort()
			return
		}
		//权限验证
		if !Permission(c, u) {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "当前账号无权限",
			})
			c.Abort()
			return
		}

		if c.Request.Method == http.MethodPost {
			key := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(u.Id)
			value := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(u.Id)
			if global.REDIS.Get(key).Val() == value {
				c.JSON(http.StatusOK, gin.H{
					"code": 10001,
					"msg":  "请勿重复提交表单",
				})
				c.Abort()
				return
			}
			global.REDIS.Set(key, value, time.Second*1)
		}
		c.Set("admin", u)
	}
}

func Permission(c *gin.Context, admin model.Admin) bool {
	//过滤权限
	ignorePath := map[string]bool{
		"/admin/api/product/remote_list":          true,
		"/admin/api/product_category/remote_list": true,
		"/admin/api/logout":                       true,
		"/admin/api/index":                        true,
		"/admin/api/config/base":                  true,
		"/admin/api/agent/list":                   true,
		"/admin/api/bank/list":                    true,
		"/admin/api/member_level/list":            true,
		"/admin/api/payment/list":                 true,
	}
	if !ignorePath[c.Request.URL.Path] {
		if admin.Role != 1 { //非超级管理员
			p := model.PermissionMap{RoleId: admin.Role}
			return p.ValidPermission(c.Request.URL.Path)
		}
	}
	return true
}
