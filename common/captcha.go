package common

import (
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
	"log"
)

var cap *captcha.Captcha

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "topgoer"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func Captcha(c *gin.Context, length ...int) {
	cap = captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}
	cap.SetSize(110, 38)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{162, 173, 211, 1})
	img, str := cap.Create(4, captcha.NUM)
	session := sessions.Default(c)
	session.Set("captcha", str)
	_ = session.Save()
	captchaId := session.Get("captcha")
	log.Println(captchaId)
	png.Encode(c.Writer, img)
}
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	captchaId := session.Get("captcha")
	log.Println(captchaId, code)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captchaId.(string) == code {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
