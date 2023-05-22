package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type MemberBankController struct {
	controller.AuthController
}

//	@Summary	银行卡列表
//	@Tags		用户银行卡
//	@Param		token	header		string			false	"用户令牌"
//	@Param		object	query		request.Request	false	"查询参数"
//	@Success	200		{object}	response.MemberBankResponse
//	@Router		/member_bank/list [get]
func (this MemberBankController) List(c *gin.Context) {
	s := service.MemberBankList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.List(*member))
	return
}

//	@Summary	绑定银行卡
//	@Tags		用户银行卡
//	@Param		token	header		string						false	"用户令牌"
//	@Param		object	body		request.MemberBankCreate	false	"查询参数"
//	@Success	200		{object}	response.Response
//	@Router		/member_bank/create [post]
func (this MemberBankController) Create(c *gin.Context) {
	s := service.MemberBankCreate{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Create(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

//	@Summary	修改银行卡
//	@Tags		用户银行卡
//	@Param		token	header		string						false	"用户令牌"
//	@Param		object	body		request.MemberBankUpdate	false	"查询参数"
//	@Success	200		{object}	response.Response
//	@Router		/member_bank/update [post]
func (this MemberBankController) Update(c *gin.Context) {
	s := service.MemberBankUpdate{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Update(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

//	@Summary	删除银行卡
//	@Tags		用户银行卡
//	@Param		token	header		string						false	"用户令牌"
//	@Param		object	body		request.MemberBankRemove	false	"查询参数"
//	@Success	200		{object}	response.Response
//	@Router		/member_bank/remove [post]
func (this MemberBankController) Remove(c *gin.Context) {
	s := service.MemberBankRemove{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Remove(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
