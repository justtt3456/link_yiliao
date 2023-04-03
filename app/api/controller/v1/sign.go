package v1

import (
	"finance/app/api/controller"
	"finance/common"
	"finance/global"
	"finance/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type SignController struct {
	controller.AuthController
}

// @Summary	签到
// @Tags		签到
// @Param		token	header		string			false	"用户令牌"
// @Param		object	query		request.Request	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/sign/sign [get]
func (this SignController) Sign(c *gin.Context) {
	member := this.MemberInfo(c)
	today := common.GetTodayZero()
	key := fmt.Sprintf("member_sign_%v_%v", today, member.ID)
	v := global.REDIS.Get(key).Val()
	if v != "" {
		this.Json(c, 10001, "今日已签到", nil)
		return
	}
	global.REDIS.Set(key, 123, -1)

	//基础配置表
	config := model.SetBase{}
	config.Get()

	//收盘状态分析
	isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
	if isRetreatStatus == true {
		if config.IncomeBalanceRate == 0 {
			//默认为90%
			config.IncomeBalanceRate = 9000
		}
		//送奖励 1块钱,可用余额,可提现余额分析
		balanceAmount := int64(model.UNITY) * int64(config.IncomeBalanceRate) / int64(model.UNITY)
		useBalanceAmount := int64(model.UNITY) - balanceAmount

		member.TotalBalance += balanceAmount
		member.Balance += balanceAmount
		member.UseBalance += useBalanceAmount
		member.Update("total_balance", "balance", "use_balance")

		//记录奖励
		trade := model.Trade{
			UID:        member.ID,
			TradeType:  22,
			ItemID:     0,
			Amount:     balanceAmount,
			Before:     member.Balance - balanceAmount,
			After:      member.Balance,
			Desc:       "签到奖励(可用余额)",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		trade.Insert()

		trade = model.Trade{
			UID:        member.ID,
			TradeType:  22,
			ItemID:     0,
			Amount:     useBalanceAmount,
			Before:     member.UseBalance - useBalanceAmount,
			After:      member.UseBalance,
			Desc:       "签到奖励(可提现余额)",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		trade.Insert()

	} else {
		//送奖励 1块钱
		member.TotalBalance += int64(model.UNITY)
		member.UseBalance += int64(model.UNITY)
		member.Update("total_balance", "use_balance")

		//记录奖励
		trade := model.Trade{
			UID:        member.ID,
			TradeType:  22,
			ItemID:     0,
			Amount:     int64(model.UNITY),
			Before:     member.UseBalance - int64(model.UNITY),
			After:      member.UseBalance,
			Desc:       "签到奖励",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		trade.Insert()
	}

	this.Json(c, 0, "签到成功", nil)
	return
}
