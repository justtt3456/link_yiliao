package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"finance/app/api/swag/request"
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginController struct {
	controller.Controller
}

// @Summary 用户登录
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body request.Login false "查询参数"
// @Success 200 {object} response.MemberResponse
// @Router /login [post]
func (this LoginController) Login(c *gin.Context) {
	s := service.LoginService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res, err := s.DoLogin(c)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 用户注册
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body request.Register false "查询参数"
// @Success 200 {object} response.MemberResponse
// @Router /register [post]
func (this LoginController) Register(c *gin.Context) {
	s := service.RegisterService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res, err := s.Insert(c)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 验证码
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body request.SendCode false "查询参数"
// @Success 200 {object} response.Response
// @Router /sendCode [post]
func (this LoginController) SendCode(c *gin.Context) {
	s := request.SendCode{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if s.Username == "" {
		this.Json(c, 10001, "手机号必传", nil)
		return
	}

	if !common.IsMobile(s.Username, global.Language) {
		this.Json(c, 10001, "手机号必传", nil)
		return
	}

	var code string
	if global.CONFIG.Phone.Dev {
		code = "1111"
	} else {
		code = common.RandIntRunes(4)
		msg := fmt.Sprintf(global.CONFIG.Phone.Sign, code)
		param := map[string]string{
			"u": global.CONFIG.Phone.Username,
			"p": common.Md5String(global.CONFIG.Phone.Password),
			"m": s.Username,
			"c": msg,
		}
		b, err := common.GetParam(global.CONFIG.Phone.Url, param, nil, nil)
		if err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		if string(b) != "0" {
			this.Json(c, 10001, "发送失败", nil)
			return
		}
	}

	var key string
	if s.Type == 1 {
		key = fmt.Sprintf("reg_%v", s.Username)
	} else {
		key = fmt.Sprintf("forget_%v", s.Username)
	}
	if global.REDIS.Get(key).Val() != "" {
		this.Json(c, 0, "验证码已发送，请间隔5分钟再尝试", nil)
		return
	}
	global.REDIS.Set(key, code, 300*time.Second)
	this.Json(c, 0, "发送成功", nil)
	return
}
