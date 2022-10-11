package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type RechargeController struct {
	controller.AuthController
}

// @Summary 充值列表
// @Tags 充值
// @Param token header string false "用户令牌"
// @Param object query request.RechargeList false "查询参数"
// @Success 200 {object} response.RechargeResponse
// @Router /recharge/page_list [get]
func (this RechargeController) PageList(c *gin.Context) {
	s := service.RechargeList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(*member))
	return
}

// @Summary 充值方式
// @Tags 充值
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.RechargeMethodResponse
// @Router /recharge/method [get]
func (this RechargeController) Method(c *gin.Context) {
	s := service.RechargeMethod{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 收款信息
// @Tags 充值
// @Param token header string false "用户令牌"
// @Param object query request.RechargeMethodInfo false "查询参数"
// @Success 200 {object} response.Response
// @Router /recharge/method_info [get]
func (this RechargeController) MethodInfo(c *gin.Context) {
	s := service.RechargeMethodInfo{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.Info())
	return
}

// @Summary 充值提交
// @Tags 充值
// @Param token header string false "用户令牌"
// @Param object body request.RechargeCreate false "查询参数"
// @Success 200 {object} response.RechargeCreate
// @Router /recharge/create [post]
func (this RechargeController) Create(c *gin.Context) {
	s := service.RechargeCreate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	res, err := s.Create(*member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if res.Data.Url == "" {
		this.Json(c, 0, "", nil)
		return
	}
	this.Json(c, 0, "", res.Data)
	return
}
