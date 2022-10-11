package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/model"
)

type Risk struct {
}

func (this Risk) Get() (*response.RiskInfo, error) {
	risk := model.Risk{}
	if !risk.Get() {
		return nil, errors.New("参数错误")
	}
	return &response.RiskInfo{
		ID:        risk.ID,
		WinList:   risk.WinList,
		LoseList:  risk.LoseList,
		WcLine:    risk.WcLine,
		WcRatio:   risk.WcRatio,
		LoseModel: risk.LoseModel,
		WinTime:   risk.WinTime,
		LoseTime:  risk.LoseTime,
	}, nil
}

type RiskUpdate struct {
	request.RiskUpdate
}

func (this RiskUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.WcLine == 0 {
		return errors.New("风控线不能为空")
	}
	if this.LoseModel == 0 {
		return errors.New("亏损模式不能为空")
	}
	m := model.Risk{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	m.WinList = this.WinList
	m.LoseList = this.LoseList
	m.WcLine = this.WcLine
	m.WcRatio = this.WcRatio
	m.LoseModel = this.LoseModel
	m.LoseTime = this.LoseTime
	m.WinTime = this.WinTime
	return m.Update()
}
