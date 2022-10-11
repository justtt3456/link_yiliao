package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	controller.AuthController
}

// Summary 资讯列表
// Tags 资讯
// Param object query request.Pagination false "查询参数"
// Success 200 {object} response.NewsResponse
// Router /news/page_list [get]
func (this NewsController) PageList(c *gin.Context) {
	s := service.News{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}
