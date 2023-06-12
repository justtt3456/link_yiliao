package v1

import (
	"china-russia/app/agent/service"
	"github.com/gin-gonic/gin"
)

type RechargeController struct {
	AuthController
}

// @Summary 充值列表
// @Tags 充值
// @Param token header string false "用户令牌"
// @Param object query request.RechargeListRequest false "查询参数"
// @Success 200 {object} response.RechargeResponse
// @Router /recharge/page_list [get]
func (this RechargeController) PageList(c *gin.Context) {
	s := service.RechargeService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	agent := this.AgentInfo(c)
	this.Json(c, 0, "ok", s.PageList(agent))
	return
}
