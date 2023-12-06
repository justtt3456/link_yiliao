package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type MedicineController struct {
	controller.AuthController
}

// @Summary	药品列表
// @Tags		药品
// @Param		token	header		string				false	"用户令牌"
// @Param		object	query		request.MedicineList	false	"查询参数"
// @Success	200		{object}	response.MedicineListResponse
// @Router		/medicine/page_list [get]
func (this MedicineController) PageList(c *gin.Context) {
	s := service.MedicineList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary	购买药品
// @Tags		药品
// @Param		token	header		string			false	"用户令牌"
// @Param		object	body		request.MedicineBuyReq	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/medicine/buy [post]
func (this MedicineController) Buy(c *gin.Context) {
	s := service.MedicineBuy{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err = s.Buy(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	投资记录
// @Tags		药品
// @Param		token	header		string					false	"用户令牌"
// @Param		object	query		request.MedicineOrder	false	"查询参数"
// @Success	200		{object}	response.MedicineBuyListResp
// @Router		/medicine/buy_list [get]
func (this MedicineController) BuyList(c *gin.Context) {
	s := service.BuyMedicineList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.List(member))
	return
}
