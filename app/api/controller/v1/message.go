package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/service"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	controller.AuthController
}

// @Summary	站内信列表
// @Tags		站内信
// @Param		object	query		request.Pagination	false	"查询参数"
// @Param		token	header		string				false	"用户令牌"
// @Success	200		{object}	response.MessageResponse
// @Router		/message/page_list [get]
func (this MessageController) PageList(c *gin.Context) {
	s := service.Message{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.PageList(*member))
	return
}

// @Summary	站内信已读
// @Tags		站内信
// @Param		token	header		string		false	"用户令牌"
// @Param		object	body		request.Msg	true	"查询参数"
// @Success	200		{object}	response.Response
// @Router		/message/read [post]
func (this MessageController) Read(c *gin.Context) {
	s := service.MessageRead{}
	//if err := c.ShouldBindJSON(&s); err != nil {
	//	this.Json(c, 10001, err.Error(), nil)
	//	return
	//}
	member := this.MemberInfo(c)
	this.Json(c, 0, "ok", s.Read(member))
	return
}
