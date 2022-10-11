package response

type UpgradeListResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data BannerData `json:"data"`
}
type UpgradeData struct {
	List []UpgradeInfo `json:"list"`
	Page Page          `json:"page"`
}
type UpgradeInfo struct {
	ID          int    `json:"id"`           //
	Platform    string `json:"platform"`     //平台
	Version     string `json:"version"`      //版本
	DownloadURL string `json:"download_url"` //下载地址
	MustUpgrade int    `json:"must_upgrade"` //强制更新 1是
	UpgradeDesc string `json:"upgrade_desc"` //更新说明
	Status      int    `json:"status"`       //状态 1启用
	CreateTime  int64  `json:"create_time"`  //
	UpdateTime  int64  `json:"update_time"`  //
}
