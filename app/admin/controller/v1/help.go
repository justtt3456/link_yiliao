package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type HelpController struct {
	AuthController
}

// @Summary 帮助中心列表
// @Tags 帮助中心
// @Param token header string false "用户令牌"
// @Param object query request.Pagination false "查询参数"
// @Success 200 {object} response.HelpListResponse
// @Router /help/page_list [get]
func (this HelpController) PageList(c *gin.Context) {
	s := service.HelpList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加帮助中心
// @Tags 帮助中心
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.HelpCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /help/create [post]
func (this HelpController) Create(c *gin.Context) {
	s := service.HelpCreate{}
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

// @Summary 修改帮助中心
// @Tags 帮助中心
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.HelpUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /help/update [post]
func (this HelpController) Update(c *gin.Context) {
	s := service.HelpUpdate{}
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

// @Summary 修改帮助中心状态
// @Tags 帮助中心
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.HelpUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /help/update_status [post]
func (this HelpController) UpdateStatus(c *gin.Context) {
	s := service.HelpUpdateStatus{}
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

// @Summary 删除帮助中心
// @Tags 帮助中心
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.HelpRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /help/remove [post]
func (this HelpController) Remove(c *gin.Context) {
	s := service.HelpRemove{}
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
