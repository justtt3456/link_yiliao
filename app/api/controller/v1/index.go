package v1

import (
	"finance/app/api/controller"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/model"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	controller.AuthController
}

// @Summary 首页banner和公告
// @Tags 首页
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.IndexResponse
// @Router /index [get]
func (this IndexController) Index(c *gin.Context) {
	res := response.Index{
		Banner: this.banner(),
		Notice: this.notice(),
	}
	this.Json(c, 0, "ok", res)
	return
}
func (this IndexController) notice() response.IndexNotice {
	//公告
	notice := model.Notice{}
	return response.IndexNotice{
		Pop:  notice.Pop(),
		Roll: notice.Roll(),
	}
}
func (this IndexController) banner() []response.Banner {
	//banner
	banner := model.Banner{
		Lang:   global.Language,
		Status: model.StatusOk,
	}
	bs := banner.List()
	res := make([]response.Banner, 0)
	for _, v := range bs {
		i := response.Banner{
			Image: v.Image,
			Link:  v.Link,
			Type:  v.Type,
		}
		res = append(res, i)
	}
	return res
}
