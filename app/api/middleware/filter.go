package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"regexp"
)

func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() != "/api/v1/upload/image" {
			var bodyBytes []byte
			if c.Request.Body != nil {
				//读取参数
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			//重新赋值用于参数绑定
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			src := string(bodyBytes)
			//HTML过滤
			re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
			if re.MatchString(src) {
				c.JSON(403, "what are you doing")
				c.Abort()
			}
		}
		//STYLE过滤
		//re, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
		//if re.MatchString(src){
		//	c.JSON(403,"what are you doing")
		//	c.Abort()
		//}
		////SCRIPT过滤
		//re, _ = regexp.Compile("\\<(?i)script[\\S\\s]+?\\</?(?i)script\\>")
		//if re.MatchString(src){
		//	c.JSON(403,"what are you doing")
		//	c.Abort()
		//}
	}
}
