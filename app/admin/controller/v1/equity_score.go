package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type EquityScoreController struct {
	AuthController
}

// @Summary 查询股权信息
// @Tags 股权
// @Param token header string false "用户令牌"
// @Param object query request.EquityScorePageList false "查询参数"
// @Success 200 {object} response.EquityScorePageListResponse
// @Router /equity_score/page_list [get]
func (this EquityScoreController) PageList(c *gin.Context) {
	s := service.EquityScoreService{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}
