package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
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
	admin := this.AdminInfo(c)
	switch s.Handle {
	//上分
	case 1:
		if err = s.Recharge(*admin, s.Handle, s.IsFrontend); err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		this.Json(c, 0, "ok", nil)
		return
	//下分,冻结
	case 2, 3:
		if err = s.Withdraw(*admin, s.IsFrontend); err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		this.Json(c, 0, "ok", nil)
		return

	//解冻(添加可提现余额)
	case 4:
		if err = s.TopupUseBalance(*admin, s.Handle, s.IsFrontend); err != nil {
			this.Json(c, 10001, err.Error(), nil)
			return
		}
		this.Json(c, 0, "ok", nil)
		return
	}

	this.Json(c, 10001, "参数错误", nil)
	return
}
