package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type UpgradeController struct {
	AuthController
}

// Summary 版本升级列表
// Tags 版本升级
// Param token header string false "用户令牌"
// Param object query request.UpgradeList false "查询参数"
// Success 200 {object} response.UpgradeListResponse
// Router /version/page_list [get]
func (this UpgradeController) PageList(c *gin.Context) {
	s := service.UpgradeList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// Summary 添加版本升级
// Tags 版本升级
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.UpgradeCreate false "查询参数"
// Success 200 {object} response.Response
// Router /version/create [post]
func (this UpgradeController) Create(c *gin.Context) {
	s := service.UpgradeCreate{}
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

// Summary 修改版本升级
// Tags 版本升级
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.UpgradeUpdate false "查询参数"
// Success 200 {object} response.Response
// Router /version/update [post]
func (this UpgradeController) Update(c *gin.Context) {
	s := service.UpgradeUpdate{}
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

// Summary 修改版本升级状态
// Tags 版本升级
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.UpgradeUpdateStatus false "查询参数"
// Success 200 {object} response.Response
// Router /version/update_status [post]
func (this UpgradeController) UpdateStatus(c *gin.Context) {
	s := service.UpgradeUpdateStatus{}
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

// Summary 删除版本升级
// Tags 版本升级
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.UpgradeRemove false "查询参数"
// Success 200 {object} response.Response
// Router /version/remove [post]
func (this UpgradeController) Remove(c *gin.Context) {
	s := service.UpgradeRemove{}
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
