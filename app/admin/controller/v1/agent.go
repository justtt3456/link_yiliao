package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type AgentController struct {
	AuthController
}

//   Summary 代理列表(一级代理不分页)
//   Tags 代理
//   Param token header string false "用户令牌"
//   Param object query request.AgentList false "查询参数"
//   Success 200 {object} response.AgentListResponse
//   Router /agent/list [get]
func (this AgentController) List(c *gin.Context) {
	s := service.AgentList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

//   Summary 代理列表
//   Tags 代理
//   Param token header string false "用户令牌"
//   Param object query request.AgentList false "查询参数"
//   Success 200 {object} response.AgentListResponse
//   Router /agent/page_list [get]
func (this AgentController) PageList(c *gin.Context) {
	s := service.AgentList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

//   Summary 添加代理
//   Tags 代理
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.AgentCreate false "查询参数"
//   Success 200 {object} response.Response
//   Router /agent/create [post]
func (this AgentController) Create(c *gin.Context) {
	s := service.AgentCreate{}
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

//   Summary 修改代理信息
//   Tags 代理
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.AgentUpdate false "查询参数"
//   Success 200 {object} response.Response
//   Router /agent/update [post]
func (this AgentController) Update(c *gin.Context) {
	s := service.AgentUpdate{}
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

//   Summary 更新代理状态
//   Tags 代理
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.AgentUpdateStatus false "查询参数"
//   Success 200 {object} response.Response
//   Router /agent/update_status [post]
func (this AgentController) UpdateStatus(c *gin.Context) {
	s := service.AgentUpdateStatus{}
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
