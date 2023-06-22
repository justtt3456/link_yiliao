package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type ProductCategoryList struct {
	request.ProductCategoryList
}

func (this ProductCategoryList) List() response.ProductCategoryData {
	m := model.ProductCategory{}
	where, args := this.getWhere()
	list := m.List(where, args)
	res := make([]response.ProductCategory, 0)
	for _, v := range list {
		i := response.ProductCategory{
			Id:         v.Id,
			Name:       v.Name,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.ProductCategoryData{List: res}
}
func (this ProductCategoryList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type ProductCategoryCreate struct {
	request.ProductCategoryCreate
}

func (this ProductCategoryCreate) Create() error {
	if this.Name == "" {
		return errors.New("分类名称不能为空")
	}
	//if this.Lang == "" {
	//	return errors.New("分类语言名称不能为空")
	//}

	m := model.ProductCategory{
		Name:   this.Name,
		Status: this.Status,
	}
	return m.Insert()
}

type ProductCategoryUpdate struct {
	request.ProductCategoryUpdate
}

func (this ProductCategoryUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("分类名称不能为空")
	}
	//if this.Lang == "" {
	//	return errors.New("分类语言不能为空")
	//}

	m := model.ProductCategory{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("分类不存在")
	}
	m.Name = this.Name
	m.Status = this.Status
	return m.Update("name", "status")
}

type ProductCategoryUpdateStatus struct {
	request.ProductCategoryUpdateStatus
}

func (this ProductCategoryUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.ProductCategory{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("分类不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ProductCategoryRemove struct {
	request.ProductCategoryRemove
}

func (this ProductCategoryRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.ProductCategory{
		Id: this.Id,
	}
	if err := m.Remove(); err != nil {
		return err
	}
	//删除分类下产品
	p := model.Product{
		Category: this.Id,
	}
	return p.Remove()
}
