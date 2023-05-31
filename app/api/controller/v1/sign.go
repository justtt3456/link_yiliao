package v1

import (
	"china-russia/app/api/controller"
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
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
	key := fmt.Sprintf("member_sign_%v_%v", today, member.Id)
	v := global.REDIS.Get(key).Val()
	if v != "" {
		this.Json(c, 10001, "今日已签到", nil)
		return
	}
	global.REDIS.Set(key, 123, -1)

	//基础配置表
	config := model.SetBase{}
	config.Get()

	//签到奖励
	member.WithdrawBalance = member.WithdrawBalance.Add(config.SignRewards)
	member.Update("withdraw_balance")

	//记录奖励
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  22,
		ItemId:     0,
		Amount:     config.SignRewards,
		Before:     member.WithdrawBalance.Sub(config.SignRewards),
		After:      member.WithdrawBalance,
		Desc:       "签到奖励",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	trade.Insert()

	this.Json(c, 0, "签到成功", nil)
	return
}
