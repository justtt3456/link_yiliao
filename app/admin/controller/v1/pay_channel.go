package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type PayChannelController struct {
	AuthController
}

// @Summary 支付通道列表
// @Tags 三方支付
// @Param token header string false "用户令牌"
// @Param object query request.PayChannelListRequest false "查询参数"
// @Success 200 {object} response.PayChannelListResponse
// @Router /pay_channel/page_list [get]
func (this PayChannelController) PageList(c *gin.Context) {
	s := service.PayChannelListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加支付通道
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PayChannelCreateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /pay_channel/create [post]
func (this PayChannelController) Create(c *gin.Context) {
	s := service.PayChannelCreateService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Create(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改支付通道
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PayChannelUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /pay_channel/update [post]
func (this PayChannelController) Update(c *gin.Context) {
	s := service.PayChannelUpdateService{}
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

// @Summary 修改支付通道状态
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PayChannelUpdateStatusRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /pay_channel/update_status [post]
func (this PayChannelController) UpdateStatus(c *gin.Context) {
	s := service.PayChannelUpdateStatusService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.UpdateStatus(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除支付通道
// @Tags 三方支付
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.PayChannelRemoveRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /pay_channel/remove [post]
func (this PayChannelController) Remove(c *gin.Context) {
	s := service.PayChannelRemoveService{}
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
