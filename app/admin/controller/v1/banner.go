package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type BannerController struct {
	AuthController
}

// @Summary banner列表
// @Tags banner
// @Param token header string false "用户令牌"
// @Param object query request.BannerList false "查询参数"
// @Success 200 {object} response.BannerListResponse
// @Router /banner/page_list [get]
func (this BannerController) PageList(c *gin.Context) {
	s := service.BannerList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加banner
// @Tags banner
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.BannerCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /banner/create [post]
func (this BannerController) Create(c *gin.Context) {
	s := service.BannerCreate{}
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

// @Summary 修改banner
// @Tags banner
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.BannerUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /banner/update [post]
func (this BannerController) Update(c *gin.Context) {
	s := service.BannerUpdate{}
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

// @Summary 修改banner状态
// @Tags banner
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.BannerUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /banner/update_status [post]
func (this BannerController) UpdateStatus(c *gin.Context) {
	s := service.BannerUpdateStatus{}
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

// @Summary 删除banner
// @Tags banner
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.BannerRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /banner/remove [post]
func (this BannerController) Remove(c *gin.Context) {
	s := service.BannerRemove{}
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
