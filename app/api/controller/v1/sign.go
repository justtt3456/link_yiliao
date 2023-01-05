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

//	@Summary	签到
//	@Tags		签到
//	@Param		token	header		string			false	"用户令牌"
//	@Param		object	query		request.Request	false	"查询参数"
//	@Success	200		{object}	response.Response
//	@Router		/sign/sign [get]
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
	//送奖励 1块钱
	member.TotalBalance += int64(model.UNITY)
	member.UseBalance += int64(model.UNITY)
	member.Update("total_balance","use_balance");
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
	this.Json(c, 0, "签到成功", nil)
	return
}
