package v1

import (
	"finance/model"
	"github.com/gin-gonic/gin"
)

type SoundController struct {
	Controller
}

func (this SoundController) Data(c *gin.Context) {
	//获取充值
	r := model.Recharge{}
	//获取提现
	w := model.Withdraw{}
	res := map[string]int64{
		"recharge": r.Count(r.TableName()+".status = ? and "+r.TableName()+".payment_id=0", []interface{}{model.StatusReview}),
		"withdraw": w.Count(w.TableName()+".status = ?", []interface{}{model.StatusReview}),
	}
	this.Json(c, 0, "ok", res)
	return
}
