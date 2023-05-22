package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	controller.Controller
}

//	@Summary	配置项
//	@Tags		首页
//	@Param		object	query		request.Request	false	"查询参数"
//	@Success	200		{object}	response.ConfigResponse
//	@Router		/config [get]
func (this ConfigController) List(c *gin.Context) {
	s := service.Config{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "", s.Get())
	return
}
