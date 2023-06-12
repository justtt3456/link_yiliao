package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
)

type GuquanList struct {
	request.Request
}

func (this *GuquanList) List() *response.GuquanResp {

	m := model.Equity{}
	if !m.Get(false) {
		return nil
	}

	return &response.GuquanResp{
		Id:              m.Id,
		TotalGuquan:     m.Total,
		OtherGuquan:     m.Current,
		Price:           m.Price,
		LimitBuy:        m.MinBuy,
		LuckyRate:       m.HitRate,
		ReturnRate:      m.MissRate,
		ReturnLuckyRate: m.SellRate,
		PreStartTime:    m.PreStartTime,
		PreEndTime:      m.PreEndTime,
		OpenTime:        m.OpenTime,
		ReturnTime:      m.RecoverTime,
		Status:          m.Status,
	}

}

type GuquanUpdate struct {
	request.GuquanReq
}

func (this *GuquanUpdate) Update() error {

	if this.Id == 0 {
		return errors.New("Id不能为空")
	}

	m := model.Equity{Id: this.Id}

	if !m.Get(false) {
		m.Total = this.TotalGuquan
		m.Current = this.OtherGuquan
		m.MinBuy = this.LimitBuy
		m.HitRate = this.LuckyRate
		m.MissRate = this.ReturnRate
		m.SellRate = this.ReturnLuckyRate
		m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
		m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
		m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
		m.RecoverTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
		m.Status = this.Status
		return m.Insert()
	}
	m.Total = this.TotalGuquan
	m.Current = this.OtherGuquan
	m.MinBuy = this.LimitBuy
	m.HitRate = this.LuckyRate
	m.MissRate = this.ReturnRate
	m.SellRate = this.ReturnLuckyRate
	m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
	m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
	m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
	m.RecoverTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
	m.Status = this.Status
	return m.Update("total", "current", "price", "min_buy", "hit_rate", "miss_rate", "sell_rate", "pre_start_time", "pre_end_time", "open_time", "recover_time", "status")
}
