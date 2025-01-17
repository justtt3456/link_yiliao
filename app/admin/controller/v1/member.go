package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	AuthController
}

// @Summary 用户列表
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object query request.MemberList false "查询参数"
// @Success 200 {object} response.MemberListResponse
// @Router /member/page_list [get]
func (this MemberController) PageList(c *gin.Context) {
	s := service.MemberList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list, err := s.PageList()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", list)
	return
}

// Summary 添加用户
// Tags 用户
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.MemberCreate false "查询参数"
// Success 200 {object} response.Response
// Router /member/create [post]
func (this MemberController) Create(c *gin.Context) {
	s := service.MemberCreate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Create(c); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改用户备注
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/update [post]
func (this MemberController) Update(c *gin.Context) {
	s := service.MemberUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改用户密码
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberUpdatePassword false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/update_password [post]
func (this MemberController) UpdatePassword(c *gin.Context) {
	s := service.MemberUpdatePassword{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.UpdatePassword(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 更新用户状态
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/update_status [post]
func (this MemberController) UpdateStatus(c *gin.Context) {
	s := service.MemberUpdateStatus{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}

	if err = s.UpdateStatus(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
func (this MemberController) Remove(c *gin.Context) {
	s := service.MemberRemove{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}

	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 我的团队
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object body request.MemberTeamReq false "查询参数"
// @Success 200 {object} response.MemberListData
// @Router /member/team [post]
func (this MemberController) Team(c *gin.Context) {
	s := service.MemberTeam{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.GetTeam()

	this.Json(c, 0, "ok", list)
	return
}

// @Summary 用户银行卡列表
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object query request.MemberBankList false "查询参数"
// @Success 200 {object} response.MemberBankListResponse
// @Router /member/bankcard/list [get]
func (this MemberController) BankCardList(c *gin.Context) {
	s := service.MemberBankList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.List()
	this.Json(c, 0, "ok", list)
	return
}

// @Summary 修改用户银行卡
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberBankUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/bankcard/update [post]
func (this MemberController) UpdateBankCard(c *gin.Context) {
	s := service.MemberBankUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除用户银行卡
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberBankRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/bankcard/remove [post]
func (this MemberController) RemoveBankCard(c *gin.Context) {
	s := service.MemberBankRemove{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 用户usdt列表
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object query request.MemberBankList false "查询参数"
// @Success 200 {object} response.MemberBankListResponse
// @Router /member_usdt/list [get]
func (this MemberController) UsdtList(c *gin.Context) {
	s := service.MemberUsdtList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.List()
	this.Json(c, 0, "ok", list)
	return
}

// @Summary 修改用户usdt
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberBankUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /member_usdt/update [post]
func (this MemberController) UpdateUsdt(c *gin.Context) {
	s := service.MemberUsdtUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除用户usdt
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberBankRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /member_usdt/remove [post]
func (this MemberController) RemoveUsdt(c *gin.Context) {
	s := service.MemberUsdtRemove{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 用户实名认证列表
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object query request.MemberVerifiedList false "查询参数"
// @Success 200 {object} response.MemberVerifiedListResponse
// @Router /member/verified/page_list [get]
func (this MemberController) VerifiedPageList(c *gin.Context) {
	s := service.MemberVerifiedList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.PageList()

	this.Json(c, 0, "ok", list)
	return
}

// @Summary 修改用户实名认证
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberVerifiedUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/verified/update [post]
func (this MemberController) UpdateVerified(c *gin.Context) {
	s := service.MemberVerifiedUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改用户实名认证信息
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberVerifiedInfoUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/verified_info/update [post]
func (this MemberController) UpdateVerifiedInfo(c *gin.Context) {
	s := service.MemberVerifiedInfoUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.UpdateInfo(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除用户实名认证
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MemberVerifiedRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/verified/remove [post]
func (this MemberController) RemoveVerified(c *gin.Context) {
	s := service.MemberVerifiedRemove{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 赠送优惠券
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object body request.SendCouponReq false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/sendCoupon [post]
func (this MemberController) SendCoupon(c *gin.Context) {
	s := service.SendCoupon{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}

	this.Json(c, 0, "ok", s.Send())
	return
}

// @Summary 赠送优惠券
// @Tags 用户
// @Param token header string false "用户令牌"
// @Param object body request.GetCodeReq false "查询参数"
// @Success 200 {object} response.Response
// @Router /member/getCode [post]
func (this MemberController) GetCode(c *gin.Context) {
	s := service.GetCode{}
	if err := c.ShouldBindJSON(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}

	this.Json(c, 0, "ok", s.GetCode())
	return
}
