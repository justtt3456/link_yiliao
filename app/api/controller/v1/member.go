package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	controller.AuthController
}

// @Summary	用户信息
// @Tags		用户
// @Param		token	header		string			false	"用户令牌"
// @Param		object	query		request.Request	false	"查询参数"
// @Success	200		{object}	response.MemberResponse
// @Router		/member/info [get]
func (this MemberController) Info(c *gin.Context) {
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", member.Info())
	return
}

// @Summary	修改用户信息
// @Tags		用户
// @Param		token	header		string				false	"用户令牌"
// @Param		object	body		request.MemberInfo	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member/update [post]
//func (this MemberController) UpdateInfo(c *gin.Context) {
//	s := service.Member{}
//	if err := c.ShouldBindJSON(&s); err != nil {
//		this.Json(c, 10001, err.Error(), nil)
//		return
//	}
//	member := this.MemberInfo(c)
//	err := s.UpdateInfo(member)
//	if err != nil {
//		this.Json(c, 10001, err.Error(), nil)
//		return
//	}
//	this.Json(c, 0, "ok", nil)
//	return
//}

// @Summary	修改登录密码
// @Tags		用户
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.MemberPassword	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member/password [post]
func (this MemberController) UpdatePassword(c *gin.Context) {
	s := service.MemberPassword{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.UpdatePassword(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	修改支付密码
// @Tags		用户
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.MemberPassword	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member/pay_password [post]
func (this MemberController) UpdatePayPassword(c *gin.Context) {
	s := service.MemberPassword{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.UpdatePayPassword(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	退出登录
// @Tags		用户
// @Param		token	header		string			false	"用户令牌"
// @Param		object	body		request.Request	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/logout [post]
func (this MemberController) Logout(c *gin.Context) {
	s := service.Member{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Logout(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	实名认证
// @Tags		用户
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.MemberVerified	false	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/member/verified [post]
func (this MemberController) Verified(c *gin.Context) {
	s := service.MemberVerified{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Verified(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	我的团队
// @Tags		用户
// @Param		token	header		string				false	"用户令牌"
// @Param		object	query		request.Pagination	false	"查询参数"
// @Success	200		{object}	response.MyTeamList
// @Router		/member/team [get]
func (this MemberController) Team(c *gin.Context) {
	s := service.MemberTeam{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	list, err := s.GetTeam(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", list)
	return
}

// @Summary	可用和可提  余额互转
// @Tags		用户
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.MemberTransfer	false	"查询参数"
// @Success	200		{object}	response.MyTeamList
// @Router		/member/transfer [post]
func (this MemberController) Transfer(c *gin.Context) {
	s := service.MemberTransfer{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Transfer(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary	可用和可提  余额互转
// @Tags		用户
// @Param		token	header		string					false	"用户令牌"
// @Param		object	body		request.MemberTransfer	false	"查询参数"
// @Success	200		{object}	response.MyTeamList
// @Router		/member/transfer [post]
func (this MemberController) MemberCoupon(c *gin.Context) {
	s := service.MemberTransfer{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err := s.Transfer(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
