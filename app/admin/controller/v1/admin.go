package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	AuthController
}

// @Summary 管理员列表
// @Tags 管理员
// @Param token header string false "用户令牌"
// @Param object query request.AdminListRequest false "查询参数"
// @Success 200 {object} response.AdminListResponse
// @Router /admin/list [get]
func (this AdminController) List(c *gin.Context) {
	s := service.AdminListService{}
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

// @Summary 添加管理员
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.AdminInsertRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /admin/create [post]
func (this AdminController) Create(c *gin.Context) {
	s := service.AdminInsertService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	admin := this.AdminInfo(c)
	res, err := s.Insert(c, *admin)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 修改管理员
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.AdminUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /admin/update [post]
func (this AdminController) Update(c *gin.Context) {
	s := service.AdminUpdateService{}
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

// @Summary 重置谷歌验证码
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.AdminUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /admin/google [post]
func (this AdminController) Google(c *gin.Context) {
	s := service.AdminGoogleService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	admin := this.AdminInfo(c)
	res, err := s.Google(*admin)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 删除管理员
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.AdminRemoveRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /admin/remove [post]
func (this AdminController) Remove(c *gin.Context) {
	s := service.AdminRemoveService{}
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

// @Summary 退出登录
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.Request false "查询参数"
// @Success 200 {object} response.Response
// @Router /logout [post]
func (this AdminController) Logout(c *gin.Context) {
	admin := this.AdminInfo(c)
	admin.Token = ""
	admin.Update("token")
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改管理员密码
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.AdminPasswordRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /admin/password [post]
func (this AdminController) Password(c *gin.Context) {
	s := service.AdminPasswordService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	admin := this.AdminInfo(c)
	if err = s.Password(*admin); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
