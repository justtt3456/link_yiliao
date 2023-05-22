package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"strings"
)

type Upgrade struct {
	request.Upgrade
}

func (this Upgrade) GetLastVersion() (*response.Upgrade, error) {
	if this.Platform == "" {
		return nil, errors.New(lang.Lang("Platform name cannot be empty"))
	}
	m := model.Upgrade{}
	platform := strings.ToLower(this.Platform)
	if platform != "ios" && platform != "android" {
		return nil, errors.New(lang.Lang("Platform name error"))
	}
	m.Platform = platform
	if m.GetLastVersion() {
		res := response.Upgrade{
			Platform:    m.Platform,
			Version:     m.Version,
			DownloadURL: m.DownloadURL,
			MustUpgrade: m.MustUpgrade,
			UpgradeDesc: m.UpgradeDesc,
		}
		return &res, nil
	}
	return nil, nil
}
