package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
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
