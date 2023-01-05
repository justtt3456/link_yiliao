package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type InvestController struct {
	controller.AuthController
}

//	@Summary	余额宝首页
//	@Tags		余额宝
//	@Param		token	header		string			false	"用户令牌"
//	@Param		object	query		request.Request	false	"查询参数"
//	@Success	200		{object}	response.InvestIndexResponse
//	@Router		/invest [get]
func (this InvestController) Index(c *gin.Context) {
	s := service.InvestIndex{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.Index(*member))
	return
}

//	@Summary	余额宝转入转出
//	@Tags		余额宝
//	@Param		token	header		string				false	"用户令牌"
//	@Param		object	body		request.InvestOrder	false	"查询参数"
//	@Success	200		{object}	response.Response
//	@Router		/invest/transfer [post]
func (this InvestController) Transfer(c *gin.Context) {
	s := service.InvestOrder{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	if err := s.Insert(*member); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

//	@Summary	余额宝转入转出记录
//	@Tags		余额宝
//	@Param		token	header		string			false	"用户令牌"
//	@Param		object	query		request.Request	false	"查询参数"
//	@Success	200		{object}	response.InvestOrderResponse
//	@Router		/invest/order [get]
func (this InvestController) Order(c *gin.Context) {
	s := service.InvestOrderList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	res := s.PageList(*member)
	this.Json(c, 0, "ok", res)
	return
}
