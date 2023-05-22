package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type GuquanController struct {
	AuthController
}

// @Summary 查询股权信息
// @Tags 股权
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.GuquanResp
// @Router /guquan/list [get]
func (this GuquanController) List(c *gin.Context) {
	s := service.GuquanList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 更改股权信息
// @Tags 股权
// @Param token header string false "用户令牌"
// @Param object body request.GuquanReq false "查询参数"
// @Success 200 {object} response.Response
// @Router /guquan/update [post]
func (this GuquanController) Update(c *gin.Context) {
	s := service.GuquanUpdate{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	err := s.Update()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
