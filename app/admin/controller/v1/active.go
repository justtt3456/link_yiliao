package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type ActiveController struct {
	AuthController
}

// @Summary 优惠券列表
// @Tags 优惠券
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.CouponResp
// @Router /active/couponList [get]
func (this ActiveController) CouponList(c *gin.Context) {
	s := service.CouponList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.CouponList())
	return
}

// @Summary 添加优惠券
// @Tags 优惠券
// @Param token header string false "用户令牌"
// @Param object body request.AddCoupon false "查询参数"
// @Success 200 {object} response.Response
// @Router /active/addCoupon [post]
func (this ActiveController) AddCoupon(c *gin.Context) {
	s := service.CouponAdd{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.Add())
	return
}

// @Summary 优惠券活动列表
// @Tags 优惠券活动
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ActiveResp
// @Router /active/page_list [get]
func (this ActiveController) PageList(c *gin.Context) {
	s := service.ActiveList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加
// @Tags 优惠券活动
// @Param token header string false "用户令牌"
// @Param object body request.AddActive false "查询参数"
// @Success 200 {object} response.Response
// @Router /active/addActive [post]
func (this ActiveController) AddActive(c *gin.Context) {
	s := service.ActiveAdd{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err := s.Add(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除
// @Tags 优惠券活动
// @Param token header string false "用户令牌"
// @Param object body request.DelActive false "查询参数"
// @Success 200 {object} response.Response
// @Router /active/delActive [post]
func (this ActiveController) DelActive(c *gin.Context) {
	s := service.DelActive{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.Del())
	return
}
