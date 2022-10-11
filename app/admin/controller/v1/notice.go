package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type NoticeController struct {
	AuthController
}

// @Summary 公告消息列表
// @Tags 公告消息
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.NoticeResponse
// @Router /notice/page_list [get]
func (this NoticeController) PageList(c *gin.Context) {
	s := service.NoticeList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加公告消息
// @Tags 公告消息
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.NoticeCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /notice/create [post]
func (this NoticeController) Create(c *gin.Context) {
	s := service.NoticeCreate{}
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

// @Summary 修改公告消息
// @Tags 公告消息
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.NoticeUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /notice/update [post]
func (this NoticeController) Update(c *gin.Context) {
	s := service.NoticeUpdate{}
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

// @Summary 修改公告消息状态
// @Tags 公告消息
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.NoticeUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /notice/update_status [post]
func (this NoticeController) UpdateStatus(c *gin.Context) {
	s := service.NoticeUpdateStatus{}
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

// @Summary 删除公告消息
// @Tags 公告消息
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.NoticeRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /notice/remove [post]
func (this NoticeController) Remove(c *gin.Context) {
	s := service.NoticeRemove{}
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
