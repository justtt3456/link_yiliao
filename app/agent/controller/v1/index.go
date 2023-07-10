package v1

import (
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type IndexController struct {
	AuthController
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
	agent := this.AgentInfo(c)
	if agent == nil {
		this.Json(c, 10000, "代理未登录", nil)
		return
	}
	ids := make([]int, 0)
	global.DB.Model(model.Member{}).Select("id").Where("agent_id = ?", agent.Id).Scan(&ids)
	if len(ids) == 0 {
		this.Json(c, 0, "ok", response.ReportData{})
		return
	}
	res := response.ReportData{
		RechargeAmount:      decimal.NewFromFloat(r.Sum("create_time >= ? and status = ? and uid in ?", []interface{}{zero, model.StatusAccept, ids}, "amount")),
		RechargeAmountTotal: decimal.NewFromFloat(r.Sum("status = ? and uid in ?", []interface{}{model.StatusAccept, ids}, "amount")),
		WithdrawAmount:      decimal.NewFromFloat(w.Sum("create_time >= ? and status = ? and uid in ?", []interface{}{zero, model.StatusAccept, ids}, "amount")),
		WithdrawAmountTotal: decimal.NewFromFloat(w.Sum("status = ? and uid in ?", []interface{}{model.StatusAccept, ids}, "amount")),
		RegCount:            m.Count("reg_time >= ? and status = ? and id in ?", []interface{}{zero, model.StatusOk, ids}),
		RegCountTotal:       m.Count("status = ? and id in ?", []interface{}{model.StatusOk, ids}),
		RegBuyCount:         m.Count("reg_time >= ? and status = ? and is_buy = ? and id in ?", []interface{}{zero, model.StatusOk, model.StatusOk, ids}),
		RegBuyCountTotal:    m.Count("status = ? and is_buy = ? and id in ?", []interface{}{model.StatusOk, model.StatusOk, ids}),
		SendMoney:           decimal.NewFromFloat(t.Sum("create_time >= ? and trade_type in (?) and uid in ?", []interface{}{zero, []int{7, 8, 16, 17}, ids}, "amount")),
		SendMoneyTotal:      decimal.NewFromFloat(t.Sum("trade_type in (?) and uid in ?", []interface{}{[]int{7, 8, 16, 17}, ids}, "amount")),
	}
	this.Json(c, 0, "ok", res)
	return
}
