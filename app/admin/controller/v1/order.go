package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	AuthController
}

// @Summary 产品订单列表
// @Tags 订单
// @Param token header string false "用户令牌"
// @Param object query request.OrderListRequest false "查询参数"
// @Success 200 {object} response.BuyListResp
// @Router /order/product_list [get]
func (this OrderController) PageList(c *gin.Context) {
	s := service.OrderListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.PageList()
	this.Json(c, 0, "ok", list)
	return
}

// @Summary 股权订单列表
// @Tags 订单
// @Param token header string false "用户令牌"
// @Param object query request.OrderListRequest false "查询参数"
// @Success 200 {object} response.BuyGuquanResp
// @Router /order/guquan_list [get]
func (this OrderController) GuQuanPageList(c *gin.Context) {
	s := service.OrderListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.GuQuanPageList()
	this.Json(c, 0, "ok", list)
	return
}

// @Summary 修改中签率
// @Tags 订单
// @Param token header string false "用户令牌"
// @Param object body request.OrderUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /order/guquan_update [post]
func (this OrderController) Update(c *gin.Context) {
	s := service.OrderUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	err = s.Update()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
