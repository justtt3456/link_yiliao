package response

type UpgradeResponse struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Upgrade `json:"data"`
}
type Upgrade struct {
	Platform    string `json:"platform"`     //平台
	Version     string `json:"version"`      //版本
	DownloadURL string `json:"download_url"` //下载地址
	MustUpgrade int    `json:"must_upgrade"` //强制更新 1是
	UpgradeDesc string `json:"upgrade_desc"` //更新说明
}
