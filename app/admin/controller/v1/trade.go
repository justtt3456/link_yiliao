package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type TradeController struct {
	AuthController
}

// @Summary 账单列表
// @Tags 账单
// @Param token header string false "用户令牌"
// @Param object query request.TradeRequest false "查询参数"
// @Success 200 {object} response.TradeResponse
// @Router /trade/page_list [get]
func (this TradeController) PageList(c *gin.Context) {
	s := service.TradeService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}
