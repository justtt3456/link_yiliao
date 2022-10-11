package request

type UpgradeList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
}
type UpgradeCreate struct {
	Platform    string `json:"platform"`     //平台
	Version     string `json:"version"`      //版本
	DownloadURL string `json:"download_url"` //下载地址
	MustUpgrade int    `json:"must_upgrade"` //强制更新 1是
	UpgradeDesc string `json:"upgrade_desc"` //更新说明
	Status      int    `json:"status"`       //状态 1启用
}
type UpgradeUpdate struct {
	ID          int    `json:"id"`
	Platform    string `json:"platform"`     //平台
	Version     string `json:"version"`      //版本
	DownloadURL string `json:"download_url"` //下载地址
	MustUpgrade int    `json:"must_upgrade"` //强制更新 1是
	UpgradeDesc string `json:"upgrade_desc"` //更新说明
	Status      int    `json:"status"`       //状态 1启用
}
type UpgradeUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"` //状态
}
type UpgradeRemove struct {
	ID int `json:"id"`
}
