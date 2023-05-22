package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Controller
}

// @Summary 管理员登录
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param object body request.LoginRequest false "查询参数"
// @Success 200 {object} response.AdminItemResponse
// @Router /login [post]
func (this LoginController) Login(c *gin.Context) {
	s := service.LoginService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	info, err := s.Login(c)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	admin := service.AdminService{}
	this.Json(c, 0, "ok", admin.Info(*info))
	return
}
