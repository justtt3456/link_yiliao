package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type UpgradeController struct {
	controller.Controller
}

// Summary 获取版本升级
// Tags 版本更新
// Param object query request.Upgrade false "查询参数"
// Success 200 {object} response.UpgradeResponse
// Router /upgrade [get]
func (this UpgradeController) Version(c *gin.Context) {
	//版本更新检查
	s := service.Upgrade{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	version, err := s.GetLastVersion()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "", version)
	return
}
