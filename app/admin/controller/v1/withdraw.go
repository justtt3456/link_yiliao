package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type WithdrawController struct {
	AuthController
}

// @Summary 提现列表
// @Tags 提现
// @Param token header string false "用户令牌"
// @Param object query request.WithdrawListRequest false "查询参数"
// @Success 200 {object} response.WithdrawResponse
// @Router /withdraw/page_list [get]
func (this WithdrawController) PageList(c *gin.Context) {
	s := service.WithdrawListService{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	err, data := s.PageList()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", data)
	return
}

// @Summary 审核提现
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.WithdrawUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /withdraw/update [post]
func (this WithdrawController) Update(c *gin.Context) {
	s := service.WithdrawUpdateService{}
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

// @Summary 修改提现
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.WithdrawUpdateRequest false "查询参数"
// @Success 200 {object} response.Response
// @Router /withdraw/update_info [post]
func (this WithdrawController) UpdateInfo(c *gin.Context) {
	s := service.WithdrawUpdateInfoService{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if err = s.UpdateInfo(); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
