package service

import (
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type Report struct {
	request.ReportSum
}

func (this Report) Member() response.MemberReportData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.MemberReport{}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.MemberReport, 0)
	for _, v := range list {
		item := response.MemberReport{
			ID:             v.ID,
			UID:            v.UID,
			Username:       v.Username,
			RechargeCount:  v.RechargeCount,
			RechargeAmount: float64(v.RechargeAmount) / 100,
			WithdrawCount:  v.WithdrawCount,
			WithdrawAmount: float64(v.WithdrawAmount) / 100,
			BetCount:       v.BetCount,
			BetAmount:      float64(v.BetAmount) / 100,
			BetResult:      float64(v.BetResult) / 100,
			SysUp:          float64(v.SysUp) / 100,
			SysDown:        float64(v.SysDown) / 100,
			Freeze:         float64(v.Freeze) / 100,
			Unfreeze:       float64(v.Unfreeze) / 100,
			CreateTime:     v.CreateTime,
			UpdateTime:     v.UpdateTime,
		}
		res = append(res, item)
	}
	return response.MemberReportData{
		List: res,
		Page: FormatPage(page),
	}
}

func (this Report) Agent() response.AgentReportData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.AgentReport{}
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	res := make([]response.AgentReport, 0)
	for _, v := range list {
		item := response.AgentReport{
			ID:             v.ID,
			Aid:            v.Aid,
			Username:       v.Username,
			RechargeCount:  v.RechargeCount,
			RechargeAmount: float64(v.RechargeAmount) / 100,
			WithdrawCount:  v.WithdrawCount,
			WithdrawAmount: float64(v.WithdrawAmount) / 100,
			BetCount:       v.BetCount,
			BetAmount:      float64(v.BetAmount) / 100,
			BetResult:      float64(v.BetResult) / 100,
			SysUp:          float64(v.SysUp) / 100,
			SysDown:        float64(v.SysDown) / 100,
			Freeze:         float64(v.Freeze) / 100,
			Unfreeze:       float64(v.Unfreeze) / 100,
			RegisterCount:  v.RegisterCount,
			CreateTime:     v.CreateTime,
			UpdateTime:     v.UpdateTime,
		}
		res = append(res, item)
	}
	return response.AgentReportData{
		List: res,
		Page: FormatPage(page),
	}
}
func (this Report) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["username"] = this.Username
	}
	if this.StartTime != "" {
		where["create_time >="] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where["create_time <="] = common.DateToUnix(this.EndTime)
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
