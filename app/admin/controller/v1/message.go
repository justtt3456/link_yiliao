package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	AuthController
}

// @Summary 站内信列表
// @Tags 站内信
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.MessageListResponse
// @Router /message/page_list [get]
func (this MessageController) PageList(c *gin.Context) {
	s := service.MessageList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加站内信
// @Tags 站内信
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MessageCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /message/create [post]
func (this MessageController) Create(c *gin.Context) {
	s := service.MessageCreate{}
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

// @Summary 修改站内信
// @Tags 站内信
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MessageUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /message/update [post]
func (this MessageController) Update(c *gin.Context) {
	s := service.MessageUpdate{}
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

// @Summary 修改站内信状态
// @Tags 站内信
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MessageUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /message/update_status [post]
func (this MessageController) UpdateStatus(c *gin.Context) {
	s := service.MessageUpdateStatus{}
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

// @Summary 删除站内信
// @Tags 站内信
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MessageRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /message/remove [post]
func (this MessageController) Remove(c *gin.Context) {
	s := service.MessageRemove{}
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
