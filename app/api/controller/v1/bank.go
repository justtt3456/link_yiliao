package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/swag/response"
	"china-russia/global"
	"china-russia/model"
	"github.com/gin-gonic/gin"
)

type BankController struct {
	controller.AuthController
}

// Summary 银行列表
// Tags 银行
// Param object query request.Request false "查询参数" @
// Success 200 {object} response.BankResponse
// Router /bank/list [get]
func (this BankController) List(c *gin.Context) {
	bank := model.Bank{
		Status: model.StatusOk,
		Lang:   global.Language,
	}
	bankList := bank.List()
	res := make([]response.Bank, 0)
	for _, v := range bankList {
		i := response.Bank{
			Id:       v.Id,
			BankName: v.BankName,
		}
		res = append(res, i)
	}
	this.Json(c, 0, "", map[string]interface{}{
		"list": res,
	})
	return
}
