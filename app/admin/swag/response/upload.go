package response

type UploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FilePath string `json:"file_path"`
	} `json:"data"`
}
