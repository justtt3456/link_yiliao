package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type MemberAddressController struct {
	controller.AuthController
}

// @Summary	收货地址列表
// @Tags		用户收货地址
// @Param		token	header		string			false	"用户令牌"
// @Param		object	query		request.Request	false	"查询参数"
// @Success	200		{object}	response.MemberAddressResponse
// @Router		/member_address/list [get]
func (this MemberAddressController) List(c *gin.Context) {
	s := service.MemberAddressList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.List(*member))
	return
}

// @Summary	绑定收货地址
// @Tags		用户收货地址
// @Param		token	header		string						false	"用户令牌"
// @Param		object	body		request.MemberAddressCreate	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member_address/create [post]
func (this MemberAddressController) Create(c *gin.Context) {
	s := service.MemberAddressCreate{}
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

// @Summary	修改收货地址
// @Tags		用户收货地址
// @Param		token	header		string						false	"用户令牌"
// @Param		object	body		request.MemberAddressUpdate	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member_address/update [post]
func (this MemberAddressController) Update(c *gin.Context) {
	s := service.MemberAddressUpdate{}
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

// @Summary	删除收货地址
// @Tags		用户收货地址
// @Param		token	header		string						false	"用户令牌"
// @Param		object	body		request.MemberAddressRemove	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member_address/remove [post]
func (this MemberAddressController) Remove(c *gin.Context) {
	s := service.MemberAddressRemove{}
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
