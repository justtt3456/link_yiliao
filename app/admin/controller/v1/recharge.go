package v1

import (
	"china-russia/app/admin/service"
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
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 审核充值
// @Tags 充值
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.RechargeUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /recharge/update [post]
func (this RechargeController) Update(c *gin.Context) {
	s := service.RechargeUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	admin := this.AdminInfo(c)
	if err = s.Update(*admin); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
