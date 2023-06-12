package v1

import (
	"china-russia/common"
	"china-russia/lang"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

type UploadController struct {
	AuthController
}

// @Summary 图片上传
// @Tags 上传
// @Accept multipart/form-data
// @Produce multipart/form-data
// @Param token header string false "用户令牌"
// @Param file formData file true "文件"
// @Success 200 {object} response.UploadResponse
// @Router /upload/image [post]
func (this UploadController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Error(err)
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	//文件格式验证
	exts := []string{".jpg", ".jpeg", ".png", ".gif"}
	extension := common.FileExt(file.Filename)
	if !common.InArray(strings.ToLower(extension), exts) {
		this.Json(c, 10001, lang.Lang("Picture format error"), nil)
		return
	}
	//最大上传 20M
	var max int64 = 20
	if file.Size > max<<20 {
		this.Json(c, 10001, fmt.Sprintf(lang.Lang("Please upload a picture within %dM"), max), nil)
		return
	}
	//文件名
	sn := common.OrderSn()
	// 上传文件到指定的路径
	filename := sn + extension
	err = c.SaveUploadedFile(file, "upload/"+filename)
	if err != nil {
		logrus.Error(err)
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", map[string]string{
		"file_path": "/upload/" + filename,
	})
	//记录用户上传了文件
	admin := this.AgentInfo(c)
	log.Println(fmt.Sprintf("管理员%s上传文件%s,上传后文件名:%s", admin.Account, file.Filename, filename))
	return
}
