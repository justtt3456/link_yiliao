package common

import (
	"china-russia/global"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
	"net/http"
	"time"
)

var cap *captcha.Captcha

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	store, _ := redis.NewStore(10, "tcp", global.CONFIG.Redis.Addr, global.CONFIG.Redis.Password, []byte("topgoer"))
	//
	//sessionMaxAge := 3600
	//sessionSecret := "topgoer"
	//var store sessions.Store
	//store = cookie.NewStore([]byte(sessionSecret))
	//store.Options(sessions.Options{
	//	MaxAge: sessionMaxAge, //seconds
	//	Path:   "/",
	//})
	return store
}

func Captcha(c *gin.Context, length ...int) {
	param := struct {
		Time int64 `json:"time" form:"time"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	if param.Time <= 0 {
		c.String(http.StatusOK, "参数错误")
		return
	}
	cap = captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}
	cap.SetSize(110, 38)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{162, 173, 211, 1})
	img, str := cap.Create(4, captcha.NUM)
	global.REDIS.Set(fmt.Sprintf("%v", param.Time), str, time.Minute)
	png.Encode(c.Writer, img)
}
func CaptchaVerify(c *gin.Context, timestamp int64, code string) bool {
	v := global.REDIS.Get(fmt.Sprintf("%v", timestamp)).Val()
	defer global.REDIS.Del(fmt.Sprintf("%v", timestamp))
	if v != code {
		return false
	}
	return true
	//session := sessions.Default(c)
	//captchaId := session.Get("captcha")
	//log.Println(captchaId, code)
	//if captchaId != nil {
	//	session.Delete("captcha")
	//	_ = session.Save()
	//	if captchaId.(string) == code {
	//		return true
	//	} else {
	//		return false
	//	}
	//} else {
	//	return false
	//}
}
