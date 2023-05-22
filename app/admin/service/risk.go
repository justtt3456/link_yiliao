package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/model"
	"errors"
)

type Risk struct {
}

func (this Risk) Get() (*response.RiskInfo, error) {
	risk := model.Risk{}
	if !risk.Get() {
		return nil, errors.New("参数错误")
	}
	return &response.RiskInfo{
		Id:        risk.Id,
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
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.WcLine == 0 {
		return errors.New("风控线不能为空")
	}
	if this.LoseModel == 0 {
		return errors.New("亏损模式不能为空")
	}
	m := model.Risk{
		Id: this.Id,
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
