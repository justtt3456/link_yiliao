package v1

import (
	"finance/app/admin/service"
	"finance/model"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	AuthController
}

// @Summary 支付列表
// @Tags 三方支付
// @Param token header string false "用户令牌"
// @Param object query request.PaymentListRequest false "查询参数"
// @Success 200 {object} response.PaymentListResponse
// @Router /payment/page_list [get]
func (this PaymentController) PageList(c *gin.Context) {
	s := service.PaymentListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 支付列表(不分页)
// @Tags 三方支付
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.Response
// @Router /payment/list [get]
func (this PaymentController) List(c *gin.Context) {
	s := model.Payment{}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 添加支付
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PaymentAddRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /payment/create [post]
func (this PaymentController) Create(c *gin.Context) {
	s := service.PaymentAddService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Add(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改支付
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PaymentUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /payment/update [post]
func (this PaymentController) Update(c *gin.Context) {
	s := service.PaymentUpdateService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除支付
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PaymentRemoveRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /payment/remove [post]
func (this PaymentController) Remove(c *gin.Context) {
	s := service.PaymentRemoveService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
