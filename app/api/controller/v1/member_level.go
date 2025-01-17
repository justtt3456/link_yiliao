package v1

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/swag/response"
	"china-russia/model"
	"github.com/gin-gonic/gin"
)

type MemberLevelController struct {
	controller.AuthController
}

// Summary 会员等级列表
// Tags 会员等级
// Param object query request.Request false "查询参数"
// Success 200 {object} response.MemberLevelResponse
// Router /member_level/list [get]
func (this MemberLevelController) List(c *gin.Context) {
	level := model.MemberLevel{}
	memberLevel := level.List()
	res := make([]response.MemberLevel, 0)
	for _, v := range memberLevel {
		i := response.MemberLevel{
			Id:   v.Id,
			Name: v.Name,
			Img:  v.Img,
		}
		res = append(res, i)
	}
	this.Json(c, 0, "", map[string]interface{}{
		"list": res,
	})
	return
}
