package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
)

type GuquanList struct {
	request.Request
}

func (this *GuquanList) List() *response.GuquanResp {

	m := model.Guquan{}
	if !m.Get(false) {
		return nil
	}

	return &response.GuquanResp{
		ID:              m.ID,
		TotalGuquan:     m.TotalGuquan,
		OtherGuquan:     m.OtherGuquan,
		ReleaseRate:     float64(m.ReleaseRate) / model.UNITY,
		Price:           float64(m.Price) / model.UNITY,
		LimitBuy:        m.LimitBuy,
		LuckyRate:       float64(m.LuckyRate) / model.UNITY,
		ReturnRate:      float64(m.ReturnRate) / model.UNITY,
		ReturnLuckyRate: float64(m.ReturnLuckyRate) / model.UNITY,
		PreStartTime:    m.PreStartTime,
		PreEndTime:      m.PreEndTime,
		OpenTime:        m.OpenTime,
		ReturnTime:      m.ReturnTime,
		Status:          m.Status,
	}

}

type GuquanUpdate struct {
	request.GuquanReq
}

func (this *GuquanUpdate) Update() error {

	if this.ID == 0 {
		return errors.New("Id不能为空")
	}

	m := model.Guquan{ID: this.ID}

	if !m.Get(false) {
		m.TotalGuquan = this.TotalGuquan
		m.OtherGuquan = this.OtherGuquan
		m.ReleaseRate = int(this.ReleaseRate * model.UNITY)
		m.LimitBuy = this.LimitBuy
		m.LuckyRate = int(this.LuckyRate * model.UNITY)
		m.ReturnRate = int(this.ReturnRate * model.UNITY)
		m.ReturnLuckyRate = int(this.ReturnLuckyRate * model.UNITY)
		m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
		m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
		m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
		m.ReturnTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
		m.Status = this.Status
		return m.Insert()
	}
	m.TotalGuquan = this.TotalGuquan
	m.OtherGuquan = this.OtherGuquan
	m.ReleaseRate = int(this.ReleaseRate * model.UNITY)
	m.LimitBuy = this.LimitBuy
	m.LuckyRate = int(this.LuckyRate * model.UNITY)
	m.ReturnRate = int(this.ReturnRate * model.UNITY)
	m.ReturnLuckyRate = int(this.ReturnLuckyRate * model.UNITY)
	m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
	m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
	m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
	m.ReturnTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
	m.Status = this.Status
	return m.Update("total_guquan", "other_guquan", "release_rate", "price", "limit_buy", "lucky_rate", "return_rate", "return_lucky_rate", "pre_start_time", "pre_end_time", "open_time", "return_time", "status")
}
