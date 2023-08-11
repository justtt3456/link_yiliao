package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"china-russia/app/api/swag/request"
	"china-russia/common"
	"china-russia/extends"
	"china-russia/global"
	"china-russia/model"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	controller.Controller
}

// @Summary	用户登录
// @Tags		用户
// @Accept		application/json
// @Produce	application/json
// @Param		object	body		request.Login	false	"查询参数"
// @Success	200		{object}	response.MemberResponse
// @Router		/login [post]
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

// @Summary	用户注册
// @Tags		用户
// @Accept		application/json
// @Produce	application/json
// @Param		object	body		request.Register	false	"查询参数"
// @Success	200		{object}	response.MemberResponse
// @Router		/register [post]
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

// @Summary	验证码
// @Tags		用户
// @Accept		application/json
// @Produce	application/json
// @Param		object	body		request.SendCode	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/send_sms [post]
func (this LoginController) SendSms(c *gin.Context) {
	s := request.SendCode{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	//验证码
	if !common.CaptchaVerify(c, s.Code) {
		this.Json(c, 10001, "验证码错误", nil)
		return
	}
	if !common.IsMobile(s.Username, global.Language) {
		this.Json(c, 10001, "手机号必传", nil)
		return
	}
	redis := model.Redis{}
	if err := redis.Lock("lock:" + s.Username); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	defer redis.Unlock("lock:" + s.Username)
	sms := extends.SmsBao{}
	err = sms.Send(s.Username)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "发送成功", nil)
	return
}
