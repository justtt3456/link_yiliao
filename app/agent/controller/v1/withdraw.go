package v1

import (
	"china-russia/app/agent/service"
	"github.com/gin-gonic/gin"
)

type WithdrawController struct {
	AuthController
}

// @Summary 提现列表
// @Tags 提现
// @Param token header string false "用户令牌"
// @Param object query request.WithdrawListRequest false "查询参数"
// @Success 200 {object} response.WithdrawResponse
// @Router /withdraw/page_list [get]
func (this WithdrawController) PageList(c *gin.Context) {
	s := service.WithdrawListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	agent := this.AgentInfo(c)
	this.Json(c, 0, "ok", s.PageList(agent))
	return
}
