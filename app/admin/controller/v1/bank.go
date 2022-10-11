package v1

import (
	"finance/app/admin/service"
	"finance/model"
	"github.com/gin-gonic/gin"
)

type BankController struct {
	AuthController
}

// Summary 可用银行列表(不分页)
// Tags 银行
// Param token header string false "用户令牌"
// Param object query request.Request false "查询参数"
// Success 200 {object} response.Response
// Router /bank/list [get]
func (this BankController) List(c *gin.Context) {
	bank := model.Bank{Status: model.StatusOk}
	this.Json(c, 0, "ok", bank.List())
	return
}

//  Summary 银行列表
//  Tags 银行
//  Param token header string false "用户令牌"
//  Param object query request.BankList false "查询参数"
//  Success 200 {object} response.BankListResponse
//  Router /bank/page_list [get]
func (this BankController) PageList(c *gin.Context) {
	s := service.BankList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

//  Summary 添加银行
//  Tags 银行
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.BankCreate false "查询参数"
//  Success 200 {object} response.Response
//  Router /bank/create [post]
func (this BankController) Create(c *gin.Context) {
	s := service.BankCreate{}
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

//  Summary 修改银行
//  Tags 银行
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.BankUpdate false "查询参数"
//  Success 200 {object} response.Response
//  Router /bank/update [post]
func (this BankController) Update(c *gin.Context) {
	s := service.BankUpdate{}
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

//  Summary 修改银行状态
//  Tags 银行
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.BankUpdateStatus false "查询参数"
//  Success 200 {object} response.Response
//  Router /bank/update_status [post]
func (this BankController) UpdateStatus(c *gin.Context) {
	s := service.BankUpdateStatus{}
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

//  Summary 删除银行
//  Tags 银行
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.BankRemove false "查询参数"
//  Success 200 {object} response.Response
//  Router /bank/remove [post]
func (this BankController) Remove(c *gin.Context) {
	s := service.BankRemove{}
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
