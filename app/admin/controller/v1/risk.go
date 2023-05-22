package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type RiskController struct {
	AuthController
}

// Summary 获取风控设置
// Tags 产品
// Param token header string false "用户令牌"
// Param object query request.Request false "查询参数"
// Success 200 {object} response.RiskResponse
// Router /risk [get]
func (this RiskController) Index(c *gin.Context) {
	s := service.Risk{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res, err := s.Get()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// Summary 修改风控设置
// Tags 产品
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.RiskUpdate false "查询参数"
// Success 200 {object} response.Response
// Router /invest/update [post]
func (this RiskController) Update(c *gin.Context) {
	s := service.RiskUpdate{}
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
