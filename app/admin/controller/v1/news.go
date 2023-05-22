package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	AuthController
}

// Summary 资讯列表
// Tags 资讯
// Param object query request.Pagination false "查询参数"
// Success 200 {object} response.NewsResponse
// Router /news/page_list [get]
func (this NewsController) PageList(c *gin.Context) {
	s := service.NewsList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// Summary 添加资讯
// Tags 资讯
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.NewsCreate false "查询参数"
// Success 200 {object} response.Response
// Router /news/create [post]
func (this NewsController) Create(c *gin.Context) {
	s := service.NewsCreate{}
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

// Summary 修改资讯
// Tags 资讯
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.NewsUpdate false "查询参数"
// Success 200 {object} response.Response
// Router /news/update [post]
func (this NewsController) Update(c *gin.Context) {
	s := service.NewsUpdate{}
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

// Summary 修改资讯状态
// Tags 资讯
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.NewsUpdateStatus false "查询参数"
// Success 200 {object} response.Response
// Router /news/update_status [post]
func (this NewsController) UpdateStatus(c *gin.Context) {
	s := service.NewsUpdateStatus{}
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

// Summary 删除资讯
// Tags 资讯
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.NewsRemove false "查询参数"
// Success 200 {object} response.Response
// Router /news/remove [post]
func (this NewsController) Remove(c *gin.Context) {
	s := service.NewsRemove{}
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
