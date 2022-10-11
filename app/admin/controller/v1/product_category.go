package v1

import (
	"finance/app/admin/service"
	"github.com/gin-gonic/gin"
)

type ProductCategoryController struct {
	AuthController
}

// @Summary 产品分类列表
// @Tags 产品
// @Param object query request.ProductCategoryList false "查询参数"
// @Success 200 {object} response.ProductCategoryListResponse
// @Router /product_category/list [get]
func (this ProductCategoryController) List(c *gin.Context) {
	s := service.ProductCategoryList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 添加产品分类
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductCategoryCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /product_category/create [post]
func (this ProductCategoryController) Create(c *gin.Context) {
	s := service.ProductCategoryCreate{}
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

// @Summary 修改产品分类
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductCategoryUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /product_category/update [post]
func (this ProductCategoryController) Update(c *gin.Context) {
	s := service.ProductCategoryUpdate{}
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

// @Summary 修改产品分类状态
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductCategoryUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /product_category/update_status [post]
func (this ProductCategoryController) UpdateStatus(c *gin.Context) {
	s := service.ProductCategoryUpdateStatus{}
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

// @Summary 删除产品分类
// @Tags 产品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.ProductCategoryRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /product_category/remove [post]
func (this ProductCategoryController) Remove(c *gin.Context) {
	s := service.ProductCategoryRemove{}
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
