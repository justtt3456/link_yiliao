package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	AuthController
}

// @Summary 角色列表
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object query request.RoleListRequest false "查询参数"
// @Success 200 {object} response.RoleListResponse
// @Router /role/list [get]
func (this RoleController) List(c *gin.Context) {
	s := service.RoleListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list, err := s.PageList()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", list)
	return
}

// @Summary 添加角色
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.RoleCreateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /role/create [post]
func (this RoleController) Create(c *gin.Context) {
	s := service.RoleCreateService{}
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

// @Summary 修改角色
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.RoleUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /role/update [post]
func (this RoleController) Update(c *gin.Context) {
	s := service.RoleUpdateService{}
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

// @Summary 删除角色
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.RoleRemoveRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /role/remove [post]
func (this RoleController) Remove(c *gin.Context) {
	s := service.RoleRemoveService{}
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
