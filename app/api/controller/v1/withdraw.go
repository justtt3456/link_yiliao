package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type WithdrawController struct {
	controller.AuthController
}

// @Summary	提现列表
// @Tags		提现
// @Param		token	header		string				false	"用户令牌"
// @Param		object	query		request.Pagination	false	"查询参数"
// @Success	200		{object}	response.WithdrawListResponse
// @Router		/withdraw/page_list [get]
func (this WithdrawController) PageList(c *gin.Context) {
	s := service.WithdrawList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(*member))
	return
}

// @Summary	提现提交
// @Tags		提现
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.WithdrawCreate	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/withdraw/create [post]
func (this WithdrawController) Create(c *gin.Context) {
	s := service.WithdrawCreate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err = s.Create(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "", nil)
	return
}

// @Summary	提现方式
// @Tags		提现
// @Param		token	header		string			false	"用户令牌"
// @Param		object	query		request.Request	false	"查询参数"
// @Success	200		{object}	response.RechargeMethodResponse
// @Router		/withdraw/method [get]
func (this WithdrawController) Method(c *gin.Context) {
	s := service.WithdrawMethod{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}
