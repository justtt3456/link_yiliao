package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"github.com/sirupsen/logrus"
)

type TradeService struct {
	request.TradeRequest
}

func (this TradeService) PageList() response.TradeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.Trade{}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	sli := make([]response.TradeInfo, 0)
	for _, v := range list {
		item := response.TradeInfo{
			Tid:       v.Id,
			Username:  v.Member.Username,
			TradeType: v.TradeType,
			//Amount:       float64(v.Amount),
			//BeforeAmount: float64(v.Before),
			//AfterAmount:  float64(v.After),
			Desc:       v.Desc,
			CreateTime: v.CreateTime,
		}
		sli = append(sli, item)
	}
	return response.TradeData{
		List: sli,
		Page: FormatPage(page),
	}
}

func (this TradeService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.StartTime != "" {
		where[model.Trade{}.TableName()+".create_time >="] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where[model.Trade{}.TableName()+".create_time <"] = common.DateToUnix(this.EndTime)
	}
	if this.TradeType != 0 {
		where[model.Trade{}.TableName()+".trade_type"] = this.TradeType
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
