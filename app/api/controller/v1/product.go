package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	controller.AuthController
}

// @Summary 产品分类
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ProductCategoryResponse
// @Router /product/category [get]
func (this ProductController) Category(c *gin.Context) {
	s := service.ProductCategory{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.Category())
	return
}

// @Summary 产品列表
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.ProductList false "查询参数"
// @Success 200 {object} response.ProductListResponse
// @Router /product/page_list [get]
func (this ProductController) PageList(c *gin.Context) {
	s := service.ProductList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 推荐列表
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.ProductListResponse
// @Router /product/recommend [get]
func (this ProductController) Recommend(c *gin.Context) {
	s := service.RecommendList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 获取一个产品
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.GetProduct false "查询参数"
// @Success 200 {object} response.Product
// @Router /product/getproduct [get]
func (this ProductController) Getproduct(c *gin.Context) {
	s := service.GetProduct{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.GetOne())
	return
}

// @Summary 股权产品
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.Request false "查询参数"
// @Success 200 {object} response.GuquanListResp
// @Router /product/guquan [get]
func (this ProductController) Guquan(c *gin.Context) {
	s := service.GuQuanList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.List())
	return
}

// @Summary 购买产品
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object body request.BuyReq false "查询参数"
// @Success 200 {object} response.Response
// @Router /product/buy [post]
func (this ProductController) Buy(c *gin.Context) {
	s := service.ProductBuy{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	err = s.Buy(member)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}

// @Summary 投资记录
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.ProductBuyList false "查询参数"
// @Success 200 {object} response.BuyListResp
// @Router /product/buy_list [get]
func (this ProductController) BuyList(c *gin.Context) {
	s := service.BuyProducList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.List(member))
	return
}

// @Summary 用户购买的股权数据
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.BuyGuquanPageList false "查询参数"
// @Success 200 {object} response.BuyGuquanPageListResp
// @Router /product/buy_guquan_list [get]
func (this *ProductController) BuyGuquanList(c *gin.Context) {
	s := service.BuyGuquanPageList{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(member))
	return
}

// @Summary 获取用户股权证书内容
// @Tags 产品
// @Param token header string false "用户令牌"
// @Param object query request.StockCertificate false "查询参数"
// @Success 200 {object} response.StockCertificateResp
// @Router /product/stock_certificate [get]
func (this *ProductController) StockCertificate(c *gin.Context) {
	s := service.StockCertificate{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.GetInfo(member))
	return
}
