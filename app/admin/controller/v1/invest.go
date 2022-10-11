package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type InvestController struct {
	AuthController
}

// @Summary 获取投资理财
// @Tags 投资理财
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.InvestResponse
// @Router /invest [get]
func (this InvestController) Index(c *gin.Context) {
	s := service.Invest{}
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

// @Summary 修改基础投资理财
// @Tags 投资理财
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.InvestUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /invest/update [post]
func (this InvestController) Update(c *gin.Context) {
	s := service.InvestUpdate{}
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

// @Summary 投资理财订单列表
// @Tags 投资理财
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.InvestOrderResponse
// @Router /invest/order/page_list [get]
func (this InvestController) OrderPageList(c *gin.Context) {
	s := service.InvestOrderList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 投资理财收益列表
// @Tags 投资理财
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.InvestIncomeResponse
// @Router /invest/income/page_list [get]
func (this InvestController) IncomePageList(c *gin.Context) {
	s := service.InvestIncomeList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}
