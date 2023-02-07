package v1

import (
	"finance/app/api/controller"
	"finance/app/api/service"
	"github.com/gin-gonic/gin"
)

type NoticeController struct {
	controller.AuthController
}

// Summary 公告列表
// Tags 公告
// Param object query request.Pagination false "查询参数"
// Param token header string false "用户令牌"
// Success 200 {object} response.NoticeResponse
// Router /notice/page_list [get]
func (this NoticeController) PageList(c *gin.Context) {
	s := service.Notice{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

//获取滚动新闻列表
func (this NoticeController) NoticeList(c *gin.Context) {
	s := service.NoticeList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}
