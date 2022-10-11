package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type TradeController struct {
	controller.AuthController
}

// @Summary 账单列表
// @Tags 账单
// @Param token header string false "用户令牌"
// @Param object query request.Trade false "查询参数"
// @Success 200 {object} response.TradeList
// @Router /trade/page_list [get]
func (this TradeController) PageList(c *gin.Context) {
	s := service.TradeService{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(*member))
	return
}

// @Summary 收益记录
// @Tags 账单
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.TradeList
// @Router /trade/income_list [get]
func (this TradeController) IncomeList(c *gin.Context) {
	s := service.Tradev2Service{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(*member))
	return
}
