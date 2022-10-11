package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type InviteCodeController struct {
	AuthController
}

// Summary 邀请码列表
// Tags 代理
// Param token header string false "用户令牌"
// Param object query request.InviteCodeList false "查询参数"
// Success 200 {object} response.InviteCodeListResponse
//  Router /invite_code/page_list [get]
func (this InviteCodeController) PageList(c *gin.Context) {
	s := service.InviteCodeList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

//  Summary 添加邀请码
//  Tags 代理
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.InviteCodeCreate false "查询参数"
//  Success 200 {object} response.Response
//  Router /invite_code/create [post]
func (this InviteCodeController) Create(c *gin.Context) {
	s := service.InviteCodeCreate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Create(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

//  Summary 修改邀请码信息
//  Tags 代理
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.InviteCodeUpdate false "查询参数"
//  Success 200 {object} response.Response
//  Router /invite_code/update [post]
func (this InviteCodeController) Update(c *gin.Context) {
	s := service.InviteCodeUpdate{}
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

//  Summary 删除邀请码
//  Tags 代理
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.InviteCodeRemove false "查询参数"
//  Success 200 {object} response.Response
//  Router /invite_code/remove [post]
func (this InviteCodeController) Remove(c *gin.Context) {
	s := service.InviteCodeRemove{}
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
