package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ManualController struct {
	AuthController
}

// Summary 手动账变列表
// Tags 账变
// Param token header string false "用户令牌"
// Param object query request.ManualListRequest false "查询参数"
// Success 200 {object} response.ManualListResponse
// Router /manual/page_list [get]
func (this ManualController) PageList(c *gin.Context) {
	s := service.ManualList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 手动账变
// @Tags 账变
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ManualRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /manual/handle [post]
func (this ManualController) Handle(c *gin.Context) {
	s := service.Manual{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if s.Amount.LessThanOrEqual(decimal.Zero) {
		this.Json(c, 10001, "金额错误", nil)
		return
	}
	//1上分可用余额  2下分可用余额  3上分可提现余额 4下分可提现 5上分股权 6下分股权
	switch s.Handle {
	case 1, 2, 3, 4:
		if err = s.Balance(); err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		this.Json(c, 0, "ok", nil)
		return
	case 5, 6:
		if err = s.Equity(); err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		this.Json(c, 0, "ok", nil)
		return
	}
	this.Json(c, 10001, "参数错误", nil)
	return
}
