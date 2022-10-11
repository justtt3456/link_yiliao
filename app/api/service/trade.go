package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/common"
	"finance/lang"
	"finance/model"
	"github.com/sirupsen/logrus"
	"time"
)

type TradeService struct {
	request.Trade
}

func (this TradeService) PageList(member model.Member) response.TradeList {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Trade{}
	where, args, _ := this.getWhere(member.ID)

	list, page := m.PageList(where, args, this.Page, this.PageSize)
	return response.TradeList{List: this.formatList(list), Page: FormatPage(page)}
}
func (this TradeService) formatList(list []model.Trade) []response.Trade {
	res := make([]response.Trade, 0)
	for _, v := range list {
		item := response.Trade{
			ID:         v.ID,
			TradeType:  v.TradeType,
			Amount:     float64(v.Amount) / model.UNITY,
			Before:     float64(v.Before) / model.UNITY,
			After:      float64(v.After) / model.UNITY,
			CreateTime: v.CreateTime,
			Desc:       v.Desc,
		}
		res = append(res, item)
	}
	return res
}
func (this TradeService) getWhere(uid int) (string, []interface{}, error) {
	where := map[string]interface{}{
		"uid": uid,
	}
	if this.Type > 0 {
		where["trade_type"] = this.Type
	}

	var start int64
	var end int64
	now := time.Now().Unix()
	if this.StartTime != "" {
		start = common.DateToUnix(this.StartTime)
		if start > now {
			return "", nil, errors.New(lang.Lang("The start time cannot be greater than the current time"))
		}
		where["create_time >"] = start
	}
	if this.EndTime != "" {
		end = common.DateToUnix(this.EndTime) + 86400
		if end < start {
			return "", nil, errors.New(lang.Lang("The start time cannot be greater than the end time"))
		}
		where["create_time <"] = end
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals, nil
}

type Tradev2Service struct {
	request.Trade
}

func (this Tradev2Service) PageList(member model.Member) response.TradeList {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Trade{}

	list, page := m.PageList("uid = ? and trade_type in (?)", []interface{}{member.ID, []int{7, 8, 13, 16, 17, 18, 19, 20, 21}}, this.Page, this.PageSize)
	return response.TradeList{List: this.formatList(list), Page: FormatPage(page)}
}
func (this Tradev2Service) formatList(list []model.Trade) []response.Trade {
	res := make([]response.Trade, 0)
	for _, v := range list {
		item := response.Trade{
			ID:         v.ID,
			TradeType:  v.TradeType,
			Amount:     float64(v.Amount) / model.UNITY,
			Before:     float64(v.Before) / model.UNITY,
			After:      float64(v.After) / model.UNITY,
			CreateTime: v.CreateTime,
			Desc:       v.Desc,
		}
		res = append(res, item)
	}
	return res
}
