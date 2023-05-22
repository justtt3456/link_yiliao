package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type MemberLevelController struct {
	AuthController
}

// Summary 用户等级列表
// Tags 用户
// Param token header string false "用户令牌"
// Param object query request.Request false "查询参数"
// Success 200 {object} response.MemberLevelListResponse
// Router /member_level/list [get]
func (this MemberLevelController) List(c *gin.Context) {
	s := service.MemberLevel{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	list := s.List()
	this.Json(c, 0, "ok", list)
	return
}

// Summary 修改用户等级
// Tags 用户
// Accept application/json
// Produce application/json
// Param token header string false "用户令牌"
// Param object body request.MemberLevelUpdate false "查询参数"
// Success 200 {object} response.Response
// Router /member_level/update [post]
func (this MemberLevelController) Update(c *gin.Context) {
	s := service.MemberLevelUpdate{}
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
