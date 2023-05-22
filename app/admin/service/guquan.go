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

	m := model.Guquan{}
	if !m.Get(false) {
		return nil
	}

	return &response.GuquanResp{
		Id: m.Id,
		//TotalGuquan:     m.TotalGuquan,
		//OtherGuquan:     m.OtherGuquan,
		//ReleaseRate:     float64(m.ReleaseRate) ,
		//Price:           float64(m.Price) ,
		//LimitBuy:        m.LimitBuy,
		//LuckyRate:       float64(m.LuckyRate) ,
		//ReturnRate:      float64(m.ReturnRate) ,
		//ReturnLuckyRate: float64(m.ReturnLuckyRate) ,
		PreStartTime: m.PreStartTime,
		PreEndTime:   m.PreEndTime,
		OpenTime:     m.OpenTime,
		ReturnTime:   m.ReturnTime,
		Status:       m.Status,
	}

}

type GuquanUpdate struct {
	request.GuquanReq
}

func (this *GuquanUpdate) Update() error {

	if this.Id == 0 {
		return errors.New("Id不能为空")
	}

	m := model.Guquan{Id: this.Id}

	if !m.Get(false) {
		//m.TotalGuquan = this.TotalGuquan
		//m.OtherGuquan = this.OtherGuquan
		//m.ReleaseRate = int(this.ReleaseRate)
		//m.LimitBuy = this.LimitBuy
		//m.LuckyRate = int(this.LuckyRate)
		//m.ReturnRate = int(this.ReturnRate)
		//m.ReturnLuckyRate = int(this.ReturnLuckyRate)
		m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
		m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
		m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
		m.ReturnTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
		m.Status = this.Status
		return m.Insert()
	}
	//m.TotalGuquan = this.TotalGuquan
	//m.OtherGuquan = this.OtherGuquan
	//m.ReleaseRate = int(this.ReleaseRate)
	//m.LimitBuy = this.LimitBuy
	//m.LuckyRate = int(this.LuckyRate)
	//m.ReturnRate = int(this.ReturnRate)
	//m.ReturnLuckyRate = int(this.ReturnLuckyRate)
	m.PreStartTime = common.DateTimeToNewYorkUnix(this.PreStartTime)
	m.PreEndTime = common.DateTimeToNewYorkUnix(this.PreEndTime)
	m.OpenTime = common.DateTimeToNewYorkUnix(this.OpenTime)
	m.ReturnTime = common.DateTimeToNewYorkUnix(this.ReturnTime)
	m.Status = this.Status
	return m.Update("total_guquan", "other_guquan", "release_rate", "price", "limit_buy", "lucky_rate", "return_rate", "return_lucky_rate", "pre_start_time", "pre_end_time", "open_time", "return_time", "status")
}
