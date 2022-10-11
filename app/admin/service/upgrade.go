package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type UpgradeList struct {
	request.UpgradeList
}

func (this UpgradeList) PageList() response.UpgradeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Upgrade{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.UpgradeInfo, 0)
	for _, v := range list {
		i := response.UpgradeInfo{
			ID:          v.ID,
			Platform:    v.Platform,
			Version:     v.Version,
			DownloadURL: v.DownloadURL,
			MustUpgrade: v.MustUpgrade,
			UpgradeDesc: v.UpgradeDesc,
			Status:      v.Status,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.UpgradeData{List: res, Page: FormatPage(page)}
}
func (this UpgradeList) getWhere() (string, []interface{}) {
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

type UpgradeCreate struct {
	request.UpgradeCreate
}

func (this UpgradeCreate) Create() error {
	if this.Platform == "" {
		return errors.New("平台名称不能为空")
	}
	if this.Version == "" {
		return errors.New("版本号不能为空")
	}
	if this.DownloadURL == "" {
		return errors.New("下载链接不能为空")
	}
	m := model.Upgrade{
		Platform:    this.Platform,
		Version:     this.Version,
		DownloadURL: this.DownloadURL,
		MustUpgrade: this.MustUpgrade,
		UpgradeDesc: this.UpgradeDesc,
		Status:      this.Status,
	}
	return m.Insert()
}

type UpgradeUpdate struct {
	request.UpgradeUpdate
}

func (this UpgradeUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Platform == "" {
		return errors.New("平台名称不能为空")
	}
	if this.Version == "" {
		return errors.New("版本号不能为空")
	}
	if this.DownloadURL == "" {
		return errors.New("下载链接不能为空")
	}
	m := model.Upgrade{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("版本不存在")
	}
	m.Platform = this.Platform
	m.Version = this.Version
	m.DownloadURL = this.DownloadURL
	m.MustUpgrade = this.MustUpgrade
	m.UpgradeDesc = this.UpgradeDesc
	m.Status = this.Status
	return m.Update("platform", "version", "download_url", "must_upgrade", "upgrade_desc", "status")
}

type UpgradeUpdateStatus struct {
	request.UpgradeUpdateStatus
}

func (this UpgradeUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Upgrade{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("版本不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type UpgradeRemove struct {
	request.UpgradeRemove
}

func (this UpgradeRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Upgrade{
		ID: this.ID,
	}
	return m.Remove()
}
