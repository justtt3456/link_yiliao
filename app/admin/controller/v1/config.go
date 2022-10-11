package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	AuthController
}

// @Summary 基础配置
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigBaseResponse
// @Router /config/base [get]
func (this ConfigController) Base(c *gin.Context) {
	s := service.ConfigBase{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res, err := s.Get()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 修改基础配置
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigBaseUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/base/update [post]
func (this ConfigController) BaseUpdate(c *gin.Context) {
	s := service.ConfigBaseUpdate{}
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

// @Summary 资金配置
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigFundsResponse
// @Router /config/funds [get]
func (this ConfigController) Funds(c *gin.Context) {
	s := service.ConfigFunds{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res, err := s.Get()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// @Summary 修改资金配置
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigFundsUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/funds/update [post]
func (this ConfigController) FundsUpdate(c *gin.Context) {
	s := service.ConfigFundsUpdate{}
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

// @Summary 收款银行卡列表
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigBankResponse
// @Router /config/bank/list [get]
func (this ConfigController) BankList(c *gin.Context) {
	s := service.ConfigBankList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 收款银行卡添加
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigBankCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/bank/create [post]
func (this ConfigController) BankCreate(c *gin.Context) {
	s := service.ConfigBankCreate{}
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

// @Summary 收款银行卡修改
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigBankUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/bank/update [post]
func (this ConfigController) BankUpdate(c *gin.Context) {
	s := service.ConfigBankUpdate{}
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

// @Summary 收款银行卡状态
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigBankUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/bank/update_status [post]
func (this ConfigController) BankUpdateStatus(c *gin.Context) {
	s := service.ConfigBankUpdateStatus{}
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

// @Summary 收款银行卡删除
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigBankRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/bank/remove [post]
func (this ConfigController) BankRemove(c *gin.Context) {
	s := service.ConfigBankRemove{}
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

//   Summary 收款支付宝列表
//   Tags 配置
//   Param token header string false "用户令牌"
//   Param object query request.Request false "查询参数"
//   Success 200 {object} response.ConfigAlipayResponse
//   Router /config/alipay/list [get]
func (this ConfigController) AlipayList(c *gin.Context) {
	s := service.ConfigAlipayList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

//   Summary 收款支付宝添加
//   Tags 配置
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.ConfigAlipayCreate false "查询参数"
//   Success 200 {object} response.Response
//   Router /config/alipay/create [post]
func (this ConfigController) AlipayCreate(c *gin.Context) {
	s := service.ConfigAlipayCreate{}
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

//   Summary 收款支付宝修改
//   Tags 配置
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.ConfigAlipayUpdate false "查询参数"
//   Success 200 {object} response.Response
//   Router /config/alipay/update [post]
func (this ConfigController) AlipayUpdate(c *gin.Context) {
	s := service.ConfigAlipayUpdate{}
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

//   Summary 收款支付宝状态
//   Tags 配置
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.ConfigAlipayUpdateStatus false "查询参数"
//   Success 200 {object} response.Response
//   Router /config/alipay/update_status [post]
func (this ConfigController) AlipayUpdateStatus(c *gin.Context) {
	s := service.ConfigAlipayUpdateStatus{}
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

//   Summary 收款支付宝删除
//   Tags 配置
//   Accept application/json
//   Produce application/json
//   Param token header string false "用户令牌"
//   Param object body request.ConfigAlipayRemove false "查询参数"
//   Success 200 {object} response.Response
//   Router /config/alipay/remove [post]
func (this ConfigController) AlipayRemove(c *gin.Context) {
	s := service.ConfigAlipayRemove{}
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

// Summary 收款usdt列表
// Tags 配置
// Param token header string false "用户令牌"
// Param object query request.Request false "查询参数"
// Success 200 {object} response.ConfigUsdtResponse
// Router /config/usdt/list [get]
func (this ConfigController) UsdtList(c *gin.Context) {
	s := service.ConfigUsdtList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// Summary 收款usdt添加
// Tags 配置
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.ConfigUsdtCreate false "查询参数"
// Success 200 {object} response.Response
// Router /config/usdt/create [post]
func (this ConfigController) UsdtCreate(c *gin.Context) {
	s := service.ConfigUsdtCreate{}
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

// Summary 收款usdt修改
// Tags 配置
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.ConfigUsdtUpdate false "查询参数"
// Success 200 {object} response.Response
// Router /config/usdt/update [post]
func (this ConfigController) UsdtUpdate(c *gin.Context) {
	s := service.ConfigUsdtUpdate{}
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

// Summary 收款usdt状态
// Tags 配置
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.ConfigUsdtUpdateStatus false "查询参数"
// Success 200 {object} response.Response
// Router /config/usdt/update_status [post]
func (this ConfigController) UsdtUpdateStatus(c *gin.Context) {
	s := service.ConfigUsdtUpdateStatus{}
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

// Summary 收款usdt删除
//  Tags 配置
//  Accept application/json
//  Produce application/json
//  Param token header string false "用户令牌"
//  Param object body request.ConfigUsdtRemove false "查询参数"
//  Success 200 {object} response.Response
//  Router /config/usdt/remove [post]
func (this ConfigController) UsdtRemove(c *gin.Context) {
	s := service.ConfigUsdtRemove{}
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

// @Summary 客服列表
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigKfResponse
// @Router /config/kf/list [get]
func (this ConfigController) KfList(c *gin.Context) {
	s := service.ConfigKfList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 客服修改
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigKfUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/kf/update [post]
func (this ConfigController) KfUpdate(c *gin.Context) {
	s := service.ConfigKfUpdate{}
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

// @Summary 客服状态
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigKfUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/kf/update_status [post]
func (this ConfigController) KfUpdateStatus(c *gin.Context) {
	s := service.ConfigKfUpdateStatus{}
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

//@Summary 语言列表
//@Tags 配置
//@Param token header string false "用户令牌"
//@Param object query request.Request false "查询参数"
//@Success 200 {object} response.ConfigLangResponse
//@Router /config/lang/list [get]
func (this ConfigController) LangList(c *gin.Context) {
	s := service.ConfigLangList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

//Summary 语言状态
//Tags 配置
//Accept application/json
// Produce application/json
//Param token header string false "用户令牌"
//Param object body request.ConfigLangUpdateStatus false "查询参数"
// Success 200 {object} response.Response
// Router /config/lang/update_status [post]
func (this ConfigController) LangUpdateStatus(c *gin.Context) {
	s := service.ConfigLangUpdateStatus{}
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

// @Summary 充值方式列表
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigWithdrawMethodResponse
// @Router /config/recharge_method/list [get]
func (this ConfigController) RechargeMethodList(c *gin.Context) {
	s := service.ConfigRechargeMethodList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 修改充值方式
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigRechargeMethodUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/recharge_method/update [post]
func (this ConfigController) RechargeMethodUpdate(c *gin.Context) {
	s := service.ConfigRechargeMethodUpdate{}
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

// @Summary 修改充值方式状态
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigRechargeMethodUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/recharge_method/update_status [post]
func (this ConfigController) RechargeMethodUpdateStatus(c *gin.Context) {
	s := service.ConfigRechargeMethodUpdateStatus{}
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

// @Summary 提现方式列表
// @Tags 配置
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ConfigWithdrawMethodResponse
// @Router /config/withdraw_method/list [get]
func (this ConfigController) WithdrawMethodList(c *gin.Context) {
	s := service.ConfigWithdrawMethodList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 修改提现方式
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigWithdrawMethodUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/withdraw_method/update [post]
func (this ConfigController) WithdrawMethodUpdate(c *gin.Context) {
	s := service.ConfigWithdrawMethodUpdate{}
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

// @Summary 修改提现方式状态
// @Tags 配置
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ConfigWithdrawMethodUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /config/withdraw_method/update_status [post]
func (this ConfigController) WithdrawMethodUpdateStatus(c *gin.Context) {
	s := service.ConfigWithdrawMethodUpdateStatus{}
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
