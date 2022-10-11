package v1

import (
	"finance/app/api/controller"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/model"
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
			ID:       v.ID,
			BankName: v.BankName,
		}
		res = append(res, i)
	}
	this.Json(c, 0, "", map[string]interface{}{
		"list": res,
	})
	return
}
