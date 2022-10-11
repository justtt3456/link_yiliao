package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type BannerList struct {
	request.BannerList
}

func (this BannerList) PageList() response.BannerData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Banner{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.BannerInfo, 0)
	for _, v := range list {
		i := response.BannerInfo{
			Id:         v.ID,
			Lang:       v.Lang,
			Image:      v.Image,
			Link:       v.Link,
			Sort:       v.Sort,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Type:       v.Type,
		}
		res = append(res, i)
	}
	return response.BannerData{List: res, Page: FormatPage(page)}
}
func (this BannerList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	if this.Lang != "" {
		where["lang"] = this.Lang
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type BannerCreate struct {
	request.BannerCreate
}

func (this BannerCreate) Create() error {
	if this.Image == "" {
		return errors.New("banner图片不能为空")
	}
	m := model.Banner{
		Image: this.Image,
		Sort:  this.Sort,
		Link:  this.Link,
		Lang:  this.Lang,
		Type:  this.Type,
	}
	return m.Insert()
}

type BannerUpdate struct {
	request.BannerUpdate
}

func (this BannerUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Image == "" {
		return errors.New("banner图片不能为空")
	}
	m := model.Banner{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("banner不存在")
	}
	m.Image = this.Image
	m.Link = this.Link
	m.Sort = this.Sort
	m.Lang = this.Lang
	m.Type = this.Type
	return m.Update("image", "link", "sort", "lang", "type")
}

type BannerUpdateStatus struct {
	request.BannerUpdateStatus
}

func (this BannerUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Banner{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("banner不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type BannerRemove struct {
	request.BannerRemove
}

func (this BannerRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Banner{
		ID: this.ID,
	}
	return m.Remove()
}
