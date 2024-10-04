package v1

import (
	"china-russia/app/admin/swag/response"
	"china-russia/app/api/controller"
	"china-russia/common"
	"china-russia/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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
	r := model.Recharge{}
	w := model.Withdraw{}
	m := model.Member{}
	t := model.Trade{}
	res := response.ReportData{
		RechargeAmount:          decimal.NewFromFloat(r.Sum("create_time >= ? and status = ?", []interface{}{zero, model.StatusAccept}, "amount")),
		RechargeAmountTotal:     decimal.NewFromFloat(r.Sum("status = ?", []interface{}{model.StatusAccept}, "amount")),
		UsdtRechargeAmount:      decimal.NewFromFloat(r.Sum("create_time >= ? and status = ? and type = 4", []interface{}{zero, model.StatusAccept}, "amount")),
		UsdtRechargeAmountTotal: decimal.NewFromFloat(r.Sum("status = ? and type = 4", []interface{}{model.StatusAccept}, "amount")),
		WithdrawAmount:          decimal.NewFromFloat(w.Sum("create_time >= ? and status = ?", []interface{}{zero, model.StatusAccept}, "total_amount")),
		WithdrawAmountTotal:     decimal.NewFromFloat(w.Sum("status = ?", []interface{}{model.StatusAccept}, "total_amount")),
		RegCount:                m.Count("reg_time >= ? and status = ?", []interface{}{zero, model.StatusOk}),
		RegCountTotal:           m.Count("status = ?", []interface{}{model.StatusOk}),
		RegBuyCount:             m.Count("reg_time >= ? and status = ? and is_buy = ?", []interface{}{zero, model.StatusOk, model.StatusOk}),
		RegBuyCountTotal:        m.Count("status = ? and is_buy = ? ", []interface{}{model.StatusOk, model.StatusOk}),
		SendMoney:               decimal.NewFromFloat(t.Sum("create_time >= ? and trade_type in (?)", []interface{}{zero, []int{7, 8, 16, 17}}, "amount")),
		SendMoneyTotal:          decimal.NewFromFloat(t.Sum("trade_type in (?)", []interface{}{[]int{7, 8, 16, 17}}, "amount")),
	}
	this.Json(c, 0, "ok", res)
	return
}
