package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type MedicineOrderController struct {
	AuthController
}

// @Summary 药品订单列表
// @Tags 药品
// @Param token header string false "用户令牌"
// @Param object query request.MedicineOrderListRequest false "查询参数"
// @Success 200 {object} response.MedicineOrderPageListResponse
// @Router /medicine/order [get]
func (this MedicineOrderController) PageList(c *gin.Context) {
	s := service.MedicineOrderListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.PageList()
	this.Json(c, 0, "ok", list)
	return
}
