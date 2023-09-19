package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/model"
	"errors"
)

type ConfigBase struct{}

func (this ConfigBase) Get() (*response.ConfigBase, error) {
	base := model.SetBase{}
	if !base.Get() {
		return nil, errors.New("参数错误")
	}
	return &response.ConfigBase{
		Id:                base.Id,
		AppName:           base.AppName,
		AppLogo:           base.AppLogo,
		VerifiedSend:      base.VerifiedSend,
		RegisterSend:      base.RegisterSend,
		OneSend:           base.OneSend,
		TwoSend:           base.TwoSend,
		ThreeSend:         base.ThreeSend,
		OneSendMoney:      base.OneSendMoney,
		SendDesc:          base.SendDesc,
		RegisterDesc:      base.RegisterDesc,
		TeamDesc:          base.TeamDesc,
		GiftRate:          base.GiftRate,
		SignRewards:       base.SignRewards,
		OneReleaseRate:    base.OneReleaseRate,
		TwoReleaseRate:    base.TwoReleaseRate,
		ThreeReleaseRate:  base.ThreeReleaseRate,
		IncomeBalanceRate: base.IncomeBalanceRate,
	}, nil
}

type ConfigBaseUpdate struct {
	request.ConfigBaseUpdate
}

func (this ConfigBaseUpdate) Update() error {
	conf := model.SetBase{}
	//添加
	if !conf.Get() {
		c := model.SetBase{
			AppName:           this.AppName,
			AppLogo:           this.AppLogo,
			VerifiedSend:      this.VerifiedSend,
			RegisterSend:      this.RegisterSend,
			OneSend:           this.OneSend,
			TwoSend:           this.TwoSend,
			ThreeSend:         this.ThreeSend,
			SendDesc:          this.SendDesc,
			RegisterDesc:      this.RegisterDesc,
			TeamDesc:          this.TeamDesc,
			OneSendMoney:      this.OneSendMoney,
			GiftRate:          this.GiftRate,
			SignRewards:       this.SignRewards,
			OneReleaseRate:    this.OneReleaseRate,
			TwoReleaseRate:    this.TwoReleaseRate,
			ThreeReleaseRate:  this.ThreeReleaseRate,
			IncomeBalanceRate: this.IncomeBalanceRate,
		}
		return c.Insert()
	} else {
		//修改
		conf.Id = this.Id
		conf.AppName = this.AppName
		conf.AppLogo = this.AppLogo
		conf.VerifiedSend = this.VerifiedSend
		conf.RegisterSend = this.RegisterSend
		conf.OneSend = this.OneSend
		conf.TwoSend = this.TwoSend
		conf.ThreeSend = this.ThreeSend
		conf.OneSendMoney = this.OneSendMoney
		conf.SendDesc = this.SendDesc
		conf.RegisterDesc = this.RegisterDesc
		conf.TeamDesc = this.TeamDesc
		conf.GiftRate = this.GiftRate
		conf.SignRewards = this.SignRewards
		conf.OneReleaseRate = this.OneReleaseRate
		conf.TwoReleaseRate = this.TwoReleaseRate
		conf.ThreeReleaseRate = this.ThreeReleaseRate
		conf.IncomeBalanceRate = this.IncomeBalanceRate
		return conf.Update()
	}
}

type ConfigFunds struct{}

func (this ConfigFunds) Get() (*response.ConfigFunds, error) {
	funds := model.SetFunds{}
	if !funds.Get() {
		return nil, errors.New("参数错误")
	}
	return &response.ConfigFunds{
		Id:                  funds.Id,
		RechargeStartTime:   funds.RechargeStartTime,
		RechargeEndTime:     funds.RechargeEndTime,
		RechargeMinAmount:   funds.RechargeMinAmount,
		RechargeMaxAmount:   funds.RechargeMaxAmount,
		RechargeFee:         funds.RechargeFee,
		RechargeQuickAmount: funds.RechargeQuickAmount,
		WithdrawStartTime:   funds.WithdrawStartTime,
		WithdrawEndTime:     funds.WithdrawEndTime,
		MustPassword:        funds.MustPassword,
		PasswordFreeze:      funds.PasswordFreeze,
		WithdrawMinAmount:   funds.WithdrawMinAmount,
		WithdrawMaxAmount:   funds.WithdrawMaxAmount,
		WithdrawFee:         funds.WithdrawFee,
		WithdrawCount:       funds.WithdrawCount,
		ProductFee:          funds.ProductFee,
		ProductQuickAmount:  funds.ProductQuickAmount,
		DayTurnMoneyNum:     funds.DayTurnMoneyNum,
	}, nil
}

type ConfigFundsUpdate struct {
	request.ConfigFundsUpdate
}

func (this ConfigFundsUpdate) Update() error {
	conf := model.SetFunds{}
	//添加
	if !conf.Get() {
		c := model.SetFunds{
			RechargeStartTime:   conf.RechargeStartTime,
			RechargeEndTime:     conf.RechargeEndTime,
			RechargeMinAmount:   conf.RechargeMinAmount,
			RechargeMaxAmount:   conf.RechargeMaxAmount,
			RechargeFee:         conf.RechargeFee,
			RechargeQuickAmount: conf.RechargeQuickAmount,
			WithdrawStartTime:   conf.WithdrawStartTime,
			WithdrawEndTime:     conf.WithdrawEndTime,
			MustPassword:        conf.MustPassword,
			PasswordFreeze:      conf.PasswordFreeze,
			WithdrawMinAmount:   conf.WithdrawMinAmount,
			WithdrawMaxAmount:   conf.WithdrawMaxAmount,
			WithdrawFee:         conf.WithdrawFee,
			WithdrawCount:       conf.WithdrawCount,
			ProductFee:          conf.ProductFee,
			ProductQuickAmount:  conf.ProductQuickAmount,
			DayTurnMoneyNum:     conf.DayTurnMoneyNum,
		}
		return c.Insert()
	} else {
		//修改
		conf.Id = this.Id
		conf.RechargeStartTime = this.RechargeStartTime
		conf.RechargeEndTime = this.RechargeEndTime
		conf.RechargeMinAmount = this.RechargeMinAmount
		conf.RechargeMaxAmount = this.RechargeMaxAmount
		conf.RechargeFee = this.RechargeFee
		conf.RechargeQuickAmount = this.RechargeQuickAmount
		conf.WithdrawStartTime = this.WithdrawStartTime
		conf.WithdrawEndTime = this.WithdrawEndTime
		conf.MustPassword = this.MustPassword
		conf.PasswordFreeze = this.PasswordFreeze
		conf.WithdrawMinAmount = this.WithdrawMinAmount
		conf.WithdrawMaxAmount = this.WithdrawMaxAmount
		conf.WithdrawFee = this.WithdrawFee
		conf.WithdrawCount = this.WithdrawCount
		conf.ProductFee = this.ProductFee
		conf.ProductQuickAmount = this.ProductQuickAmount
		conf.DayTurnMoneyNum = this.DayTurnMoneyNum
		return conf.Update()
	}
}

// bank
type ConfigBankList struct {
}

func (this ConfigBankList) List() response.ConfigBankData {
	m := model.SetBank{}
	list := m.List(false)
	res := make([]response.ConfigBank, 0)
	for _, v := range list {
		i := response.ConfigBank{
			Id:         v.Id,
			BankName:   v.BankName,
			CardNumber: v.CardNumber,
			BranchBank: v.BranchBank,
			RealName:   v.RealName,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ConfigBankData{List: res}
}

type ConfigBankCreate struct {
	request.ConfigBankCreate
}

func (this ConfigBankCreate) Create() error {
	if this.BankName == "" {
		return errors.New("收款银行不能为空")
	}
	if this.RealName == "" {
		return errors.New("收款人姓名不能为空")
	}
	if this.CardNumber == "" {
		return errors.New("收款卡号不能为空")
	}
	if this.BranchBank == "" {
		return errors.New("支行不能为空")
	}
	m := model.SetBank{
		BankName:   this.BankName,
		CardNumber: this.CardNumber,
		BranchBank: this.BranchBank,
		RealName:   this.RealName,
		Lang:       "zh_cn",
		Status:     this.Status,
	}
	return m.Insert()
}

type ConfigBankUpdate struct {
	request.ConfigBankUpdate
}

func (this ConfigBankUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.BankName == "" {
		return errors.New("收款银行不能为空")
	}
	if this.RealName == "" {
		return errors.New("收款人姓名不能为空")
	}
	if this.CardNumber == "" {
		return errors.New("收款卡号不能为空")
	}
	if this.BranchBank == "" {
		return errors.New("支行不能为空")
	}
	m := model.SetBank{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款银行卡不存在")
	}
	m.BankName = this.BankName
	m.CardNumber = this.CardNumber
	m.RealName = this.RealName
	m.BranchBank = this.BranchBank
	m.Status = this.Status
	return m.Update("bank_name", "card_number", "real_name", "branch_bank", "status")
}

type ConfigBankUpdateStatus struct {
	request.ConfigBankUpdateStatus
}

func (this ConfigBankUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetBank{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款银行卡不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ConfigBankRemove struct {
	request.ConfigBankRemove
}

func (this ConfigBankRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetBank{
		Id: this.Id,
	}
	return m.Remove()
}

// alipay
type ConfigAlipayList struct {
}

func (this ConfigAlipayList) List() response.ConfigAlipayData {

	m := model.SetAlipay{}
	list := m.List(false)
	res := make([]response.ConfigAlipay, 0)
	for _, v := range list {
		i := response.ConfigAlipay{
			Id:         v.Id,
			Account:    v.Account,
			RealName:   v.RealName,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ConfigAlipayData{List: res}
}

type ConfigAlipayCreate struct {
	request.ConfigAlipayCreate
}

func (this ConfigAlipayCreate) Create() error {
	if this.Account == "" {
		return errors.New("收款账号不能为空")
	}
	if this.RealName == "" {
		return errors.New("收款人姓名不能为空")
	}
	m := model.SetAlipay{
		Account:  this.Account,
		RealName: this.RealName,
		Lang:     this.Lang,
		Status:   this.Status,
	}
	return m.Insert()
}

type ConfigAlipayUpdate struct {
	request.ConfigAlipayUpdate
}

func (this ConfigAlipayUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Account == "" {
		return errors.New("收款账号不能为空")
	}
	if this.RealName == "" {
		return errors.New("收款人姓名不能为空")
	}
	m := model.SetAlipay{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款支付宝不存在")
	}
	m.RealName = this.RealName
	m.Account = this.Account
	m.Status = this.Status
	m.Lang = this.Lang
	return m.Update("real_name", "account", "status", "lang")
}

type ConfigAlipayUpdateStatus struct {
	request.ConfigAlipayUpdateStatus
}

func (this ConfigAlipayUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetAlipay{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款支付宝不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ConfigAlipayRemove struct {
	request.ConfigAlipayRemove
}

func (this ConfigAlipayRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetAlipay{
		Id: this.Id,
	}
	return m.Remove()
}

// usdt
type ConfigUsdtList struct {
}

func (this ConfigUsdtList) List() response.ConfigUsdtData {
	m := model.SetUsdt{}
	list := m.List(false)
	res := make([]response.ConfigUsdt, 0)
	for _, v := range list {
		i := response.ConfigUsdt{
			Id:         v.Id,
			Address:    v.Address,
			Status:     v.Status,
			Proto:      v.Proto,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ConfigUsdtData{List: res}
}

type ConfigUsdtCreate struct {
	request.ConfigUsdtCreate
}

func (this ConfigUsdtCreate) Create() error {
	if this.Address == "" {
		return errors.New("usdt收款地址不能为空")
	}
	//if this.Proto != 1 && this.Proto != 2 {
	//	return errors.New("usdt收地址址协议错误")
	//}
	m := model.SetUsdt{
		Address: this.Address,
		Status:  this.Status,
		Proto:   this.Proto,
	}
	return m.Insert()
}

type ConfigUsdtUpdate struct {
	request.ConfigUsdtUpdate
}

func (this ConfigUsdtUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Address == "" {
		return errors.New("usdt收款地址不能为空")
	}
	//if this.Proto != 1 && this.Proto != 2 {
	//	return errors.New("usdt收地址址协议错误")
	//}
	m := model.SetUsdt{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款usdt不存在")
	}
	m.Address = this.Address
	m.Proto = this.Proto
	m.Status = this.Status
	return m.Update("address", "proto", "status")
}

type ConfigUsdtUpdateStatus struct {
	request.ConfigUsdtUpdateStatus
}

func (this ConfigUsdtUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetUsdt{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("收款usdt不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ConfigUsdtRemove struct {
	request.ConfigUsdtRemove
}

func (this ConfigUsdtRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetUsdt{
		Id: this.Id,
	}
	return m.Remove()
}

// kf
type ConfigKfList struct {
}

func (this ConfigKfList) List() response.ConfigKfData {
	m := model.SetKf{}
	list := m.List(false)
	res := make([]response.ConfigKf, 0)
	for _, v := range list {
		i := response.ConfigKf{
			Id:         v.Id,
			Name:       v.Name,
			StartTime:  v.StartTime,
			EndTime:    v.EndTime,
			Link:       v.Link,
			Key:        v.Key,
			Icon:       v.Icon,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ConfigKfData{List: res}
}

type ConfigKfUpdate struct {
	request.ConfigKfUpdate
}

func (this ConfigKfUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.StartTime == "" {
		return errors.New("上班时间不能为空")
	}
	if this.EndTime == "" {
		return errors.New("下班时间不能为空")
	}
	if this.Link == "" {
		return errors.New("跳转链接不能为空")
	}
	m := model.SetKf{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("客服不存在")
	}
	m.Name = this.Name
	m.StartTime = this.StartTime
	m.EndTime = this.EndTime
	m.Link = this.Link
	m.Key = this.Key
	m.Icon = this.Icon
	m.Status = this.Status
	return m.Update("name", "start_time", "end_time", "link", "key", "icon", "lang")
}

type ConfigKfUpdateStatus struct {
	request.ConfigKfUpdateStatus
}

func (this ConfigKfUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetKf{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("客服不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

// lang
type ConfigLangList struct {
	Status int `form:"status"`
}

func (this ConfigLangList) List() response.ConfigLangData {
	m := model.SetLang{}
	if this.Status != 0 {
		m.Status = this.Status
	}
	list := m.List(false)
	res := make([]response.ConfigLang, 0)
	for _, v := range list {
		i := response.ConfigLang{
			Id:         v.Id,
			Name:       v.Name,
			Code:       v.Code,
			Icon:       v.Icon,
			IsDefault:  v.IsDefault,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ConfigLangData{List: res}
}

type ConfigLangUpdateStatus struct {
	request.ConfigLangUpdateStatus
}

func (this ConfigLangUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.SetLang{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("语言不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

// recharge method
type ConfigRechargeMethodList struct {
}

func (this ConfigRechargeMethodList) List() response.ConfigRechargeMethodData {
	m := model.RechargeMethod{}
	list, _ := m.List()
	res := make([]response.ConfigRechargeMethod, 0)
	for _, v := range list {
		i := response.ConfigRechargeMethod{
			Id:     v.Id,
			Name:   v.Name,
			Code:   v.Code,
			Icon:   v.Icon,
			Lang:   v.Lang,
			Status: v.Status,
		}
		res = append(res, i)
	}
	return response.ConfigRechargeMethodData{List: res}
}

type ConfigRechargeMethodUpdate struct {
	request.ConfigRechargeMethodUpdate
}

func (this ConfigRechargeMethodUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.RechargeMethod{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("充值方式不存在")
	}
	m.Name = this.Name
	m.Icon = this.Icon
	return m.Update("name", "icon")
}

type ConfigRechargeMethodUpdateStatus struct {
	request.ConfigRechargeMethodUpdateStatus
}

func (this ConfigRechargeMethodUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.RechargeMethod{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("充值方式不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

// withdraw method
type ConfigWithdrawMethodList struct {
}

func (this ConfigWithdrawMethodList) List() response.ConfigWithdrawMethodData {
	m := model.WithdrawMethod{}
	list, _ := m.List()
	res := make([]response.ConfigWithdrawMethod, 0)
	for _, v := range list {
		i := response.ConfigWithdrawMethod{
			Id:     v.Id,
			Name:   v.Name,
			Code:   v.Code,
			Icon:   v.Icon,
			Status: v.Status,
			Fee:    v.Fee,
		}
		res = append(res, i)
	}
	return response.ConfigWithdrawMethodData{List: res}
}

type ConfigWithdrawMethodUpdate struct {
	request.ConfigWithdrawMethodUpdate
}

func (this ConfigWithdrawMethodUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.WithdrawMethod{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("提现方式不存在")
	}
	m.Name = this.Name
	m.Icon = this.Icon
	return m.Update("name", "icon")
}

type ConfigWithdrawMethodUpdateStatus struct {
	request.ConfigWithdrawMethodUpdateStatus
}

func (this ConfigWithdrawMethodUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.WithdrawMethod{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("提现方式不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}
