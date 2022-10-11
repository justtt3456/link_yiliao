package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type ProductList struct {
	request.ProductList
}

func (this ProductList) PageList() response.ProductData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Product{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Product, 0)
	for _, v := range list {
		i := response.Product{
			ID:           v.ID,
			Name:         v.Name,
			Category:     v.Category,
			CategoryName: v.ProductCategory.Name,
			CreateTime:   v.CreateTime,
			Status:       v.Status,
			Tag:          v.Tag,
			TimeLimit:    v.TimeLimit,
			IsRecommend:  v.IsRecommend,
			Dayincome:    float64(v.Dayincome) / model.UNITY,
			Price:        float64(v.Price) / model.UNITY,
			TotalPrice:   float64(v.TotalPrice) / model.UNITY,
			OtherPrice:   float64(v.OtherPrice) / model.UNITY,
			MoreBuy:      v.MoreBuy,
			Desc:         v.Desc,
			IsFinish:     v.IsFinish,
			IsManjian:    v.IsManjian,
			BuyTimeLimit: v.BuyTimeLimit,
			Progress:     float64(v.Progress) / model.UNITY,
		}
		res = append(res, i)
	}
	return response.ProductData{List: res, Page: FormatPage(page)}
}
func (this ProductList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where[model.Product{}.TableName()+".name"] = this.Name
	}
	if this.Category > 0 {
		where[model.Product{}.TableName()+".category"] = this.Category
	}
	if this.Status > 0 {
		where[model.Product{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type ProductCreate struct {
	request.ProductCreate
}

func (this ProductCreate) Create() error {
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Category == 0 {
		return errors.New("产品类别不能为空")
	}

	if this.TimeLimit == 0 {
		return errors.New("产品投资期限不能为空")
	}
	if this.Dayincome == 0 {
		return errors.New("每日收益不能为空")
	}
	if this.Price == 0 {
		return errors.New("最低买入不能为空")
	}
	if this.TotalPrice == 0 {
		return errors.New("项目规模不能为空")
	}
	if this.OtherPrice == 0 {
		return errors.New("可投余额不能为空")
	}

	if this.MoreBuy == 0 {
		this.MoreBuy = 99999
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.IsRecommend == 0 {
		return errors.New("是否推荐到首页不能为空")
	}

	if this.IsFinish == 0 {
		return errors.New("是否投资完毕不能为空")
	}

	if this.IsManjian == 0 {
		return errors.New("是否有满减活动不能为空")
	}
	if this.BuyTimeLimit == 0 {
		return errors.New("产品限时多少天不能为空")
	}

	ex := model.Product{
		Name: this.Name,
	}
	if ex.Get() {
		return errors.New("产品已存在")
	}
	m := model.Product{
		Name:         this.Name,
		Category:     this.Category,
		CreateTime:   this.CreateTime,
		Status:       this.Status,
		Tag:          this.Tag,
		TimeLimit:    this.TimeLimit,
		IsRecommend:  this.IsRecommend,
		Dayincome:    int(this.Dayincome * model.UNITY),
		Price:        int64(this.Price * model.UNITY),
		TotalPrice:   int64(this.TotalPrice * model.UNITY),
		OtherPrice:   int64(this.OtherPrice * model.UNITY),
		MoreBuy:      this.MoreBuy,
		Desc:         this.Desc,
		IsFinish:     this.IsFinish,
		IsManjian:    this.IsManjian,
		BuyTimeLimit: this.BuyTimeLimit,
		Progress:     int(this.Progress * model.UNITY),
	}
	return m.Insert()
}

type ProductUpdate struct {
	request.ProductUpdate
}

func (this ProductUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Category == 0 {
		return errors.New("产品类别不能为空")
	}

	if this.TimeLimit == 0 {
		return errors.New("产品投资期限不能为空")
	}
	if this.Dayincome == 0 {
		return errors.New("每日收益不能为空")
	}
	if this.Price == 0 {
		return errors.New("最低买入不能为空")
	}
	if this.TotalPrice == 0 {
		return errors.New("项目规模不能为空")
	}
	if this.OtherPrice == 0 {
		return errors.New("可投余额不能为空")
	}
	if this.MoreBuy == 0 {
		this.MoreBuy = 99999
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.IsRecommend == 0 {
		return errors.New("是否推荐到首页不能为空")
	}

	if this.IsFinish == 0 {
		return errors.New("是否投资完毕不能为空")
	}

	if this.IsManjian == 0 {
		return errors.New("是否有满减活动不能为空")
	}
	if this.BuyTimeLimit == 0 {
		return errors.New("产品限时多少天不能为空")
	}
	m := model.Product{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}

	m.Name = this.Name
	m.Category = this.Category
	m.Status = this.Status
	m.TimeLimit = this.TimeLimit
	m.Dayincome = int(this.Dayincome * model.UNITY)
	m.Price = int64(this.Price * model.UNITY)
	m.TotalPrice = int64(this.TotalPrice * model.UNITY)
	m.OtherPrice = int64(this.OtherPrice * model.UNITY)
	m.MoreBuy = this.MoreBuy
	m.Desc = this.Desc
	m.Status = this.Status
	m.IsRecommend = this.IsRecommend
	m.IsFinish = this.IsFinish
	m.IsManjian = this.IsManjian
	m.BuyTimeLimit = this.BuyTimeLimit
	m.Tag = this.Tag
	m.Progress = int(this.Progress * model.UNITY)

	return m.Update("name", "progress", "buy_time_limit", "category", "create_time", "status", "tag", "time_limit", "is_recommend", "day_income", "price", "total_price", "other_price", "more_buy", "desc", "is_finish", "is_manjian")
}

type ProductUpdateStatus struct {
	request.ProductUpdateStatus
}

func (this ProductUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ProductRemove struct {
	request.ProductRemove
}

func (this ProductRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		ID: this.ID,
	}
	return m.Remove()
}
