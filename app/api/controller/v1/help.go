package v1

import (
	"finance/app/api/controller"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/lang"
	"finance/model"
	"github.com/gin-gonic/gin"
)

type HelpController struct {
	controller.Controller
}

//	@Summary	公司简介和推荐奖励
//	@Tags		公司简介和推荐奖励
//	@Param		object	query		request.Help	false	"查询参数"
//	@Success	200		{object}	response.HelpListResponse
//	@Router		/help/list [get]
func (this HelpController) List(c *gin.Context) {
	var param request.Help
	err := c.ShouldBindQuery(&param)
	if err != nil {
		this.Json(c, 10001, lang.Lang("Parameter error"), nil)
		return
	}
	m := model.Help{
		Status:   model.StatusOk,
		Lang:     global.Language,
		Category: param.Category,
	}
	list := m.List()
	res := make([]response.Help, 0)
	for _, v := range list {
		i := response.Help{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, i)
	}
	this.Json(c, 0, "ok", response.HelpData{List: res})
	return
}
