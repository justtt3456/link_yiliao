package v1

import (
	"finance/app/api/controller"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/model"
	"github.com/gin-gonic/gin"
)

type CountryController struct {
	controller.Controller
}

// Summary 国家列表
// Tags 国家
// Param object query request.Request false "查询参数"
// Success 200 {object} response.CountryListResponse
// Router /country/list [get]
func (this CountryController) List(c *gin.Context) {
	m := model.Country{
		IsReg: model.StatusOk,
	}
	list := m.List()
	res := make([]response.Country, 0)
	for _, v := range list {
		i := response.Country{
			Code: v.Code,
		}
		if global.Language == "zh_cn" {
			i.Name = v.ZhName
		} else {
			i.Name = v.EnName
		}
		res = append(res, i)
	}
	this.Json(c, 0, "ok", response.CountryData{List: res})
	return
}
