package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	AuthController
}

// @Summary 权限列表
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.PermissionListResponse
// @Router /permission/list [get]
func (this PermissionController) List(c *gin.Context) {
	s := service.PermissionService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 权限添加
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object body request.PermissionCreateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /permission/create [post]
func (this PermissionController) Create(c *gin.Context) {
	s := service.PermissionCreateService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err := s.Create(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 权限修改
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object body request.PermissionUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /permission/update [post]
func (this PermissionController) Update(c *gin.Context) {
	s := service.PermissionUpdateService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err := s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 权限删除
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object body request.PermissionRemoveRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /permission/remove [post]
func (this PermissionController) Remove(c *gin.Context) {
	s := service.PermissionRemoveService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err := s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
