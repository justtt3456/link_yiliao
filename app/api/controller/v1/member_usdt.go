package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type MemberUsdtController struct {
	controller.AuthController
}

// Summary usdt列表
// Tags 用户usdt
// Param token header string false "用户令牌"
// Param object query request.Request false "查询参数"
// Success 200 {object} response.MemberBankResponse
// Router /member_bank/list [get]
func (this MemberUsdtController) List(c *gin.Context) {
	s := service.MemberUsdtList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.List(*member))
	return
}

// Summary 绑定usdt
// Tags 用户usdt
// Param token header string false "用户令牌"
// Param object body request.MemberBankCreate false "查询参数"
// Success 200 {object} response.Response
// Router /member_bank/create [post]
func (this MemberUsdtController) Create(c *gin.Context) {
	s := service.MemberUsdtCreate{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Create(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
