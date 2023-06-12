package v1

import (
	"china-russia/app/agent/service"
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
	agent := this.AgentInfo(c)
	list, err := s.PageList(agent)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", list)
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
	agent := this.AgentInfo(c)
	list := s.PageList(agent)
	this.Json(c, 0, "ok", list)
	return
}
