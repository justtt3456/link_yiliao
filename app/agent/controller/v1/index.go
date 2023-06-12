package v1

import (
	"china-russia/app/agent/swag/response"
	"china-russia/app/api/controller"
	"china-russia/common"
	"china-russia/model"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	controller.Controller
}

// @Summary 首页报表
// @Tags 首页
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ReportResponse
// @Router /index [get]
func (this IndexController) Report(c *gin.Context) {
	zero := common.GetTodayZero()
	//r := model.Recharge{}
	//w := model.Withdraw{}
	m := model.Member{}
	//t := model.Trade{}
	res := response.ReportData{
		//RechargeAmount:      float64(r.Sum("create_time >= ? and status = ?", []interface{}{zero, 2}, "amount")),
		//RechargeAmountTotal: float64(r.Sum("status = ?", []interface{}{2}, "amount")),
		//WithdrawAmount:      float64(w.Sum("create_time >= ? and status = ?", []interface{}{zero, 2}, "amount")),
		//WithdrawAmountTotal: float64(w.Sum("status = ?", []interface{}{2}, "amount")),
		RegCount:         m.Count("reg_time >= ? and status = ?", []interface{}{zero, 1}),
		RegCountTotal:    m.Count("status = ?", []interface{}{1}),
		RegBuyCount:      m.Count("reg_time >= ? and status = ? and is_buy = ?", []interface{}{zero, 1, 1}),
		RegBuyCountTotal: m.Count("status = ? and is_buy = ? ", []interface{}{1, 1}),
		//SendMoney:           float64(t.Sum("create_time >= ? and trade_type in (?)", []interface{}{zero, []int{7, 8, 10, 13, 16, 17, 18, 19, 20, 21}}, "amount")),
		//SendMoneyTotal:      float64(t.Sum("trade_type in (?)", []interface{}{[]int{7, 8, 10, 13, 16, 17, 18, 19, 20, 21}}, "amount")),
	}
	this.Json(c, 0, "ok", res)
	return
}
