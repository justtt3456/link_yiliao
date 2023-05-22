package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	AuthController
}

// @Summary 产品列表
// @Tags 产品
// @Param object query request.ProductList false "查询参数"
// @Success 200 {object} response.ProductListResponse
// @Router /product/page_list [get]
func (this ProductController) PageList(c *gin.Context) {
	s := service.ProductList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加产品
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /product/create [post]
func (this ProductController) Create(c *gin.Context) {
	s := service.ProductCreate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Create(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改产品
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /product/update [post]
func (this ProductController) Update(c *gin.Context) {
	s := service.ProductUpdate{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Update(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 修改产品状态
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /product/update_status [post]
func (this ProductController) UpdateStatus(c *gin.Context) {
	s := service.ProductUpdateStatus{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.UpdateStatus(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 删除产品
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /product/remove [post]
func (this ProductController) Remove(c *gin.Context) {
	s := service.ProductRemove{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.Remove(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 赠品产品列表(下拉菜单内容)
// @Tags 产品
// @Param object query request.ProductList false "查询参数"
// @Success 200 {object} response.ProductListResponse
// @Router /product/gift_options [get]
func (this ProductController) GiftOptions(c *gin.Context) {
	s := service.GiftProductOptions{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.GiftList())
	return
}
