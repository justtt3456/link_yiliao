package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"github.com/sirupsen/logrus"
)

type EquityScoreService struct {
	request.EquityScorePageList
}

func (this EquityScoreService) PageList() *response.EquityScorePageListResponse {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.EquityScoreOrder{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.EquityScoreOrder, 0)

	for _, v := range list {
		items = append(items, response.EquityScoreOrder{
			Id:         v.Id,
			UId:        v.UId,
			Username:   v.Member.Username,
			PayMoney:   v.PayMoney,
			Rate:       v.Rate,
			Interval:   v.Interval,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			EndTime:    v.EndTime,
		})

	}
	return &response.EquityScorePageListResponse{List: items, Page: FormatPage(page)}
}
func (this EquityScoreService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.Uid != 0 {
		where["Member.id"] = this.Uid
	}
	//if this.StartTime != "" {
	//	where["create_time >="] = common.DateToUnix(this.StartTime)
	//}
	//if this.EndTime != "" {
	//	where["create_time <"] = common.DateToUnix(this.EndTime)
	//}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
