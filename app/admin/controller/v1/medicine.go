package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type MedicineController struct {
	AuthController
}

// @Summary 药品列表
// @Tags 药品
// @Param object query request.MedicineList false "查询参数"
// @Success 200 {object} response.MedicineListResponse
// @Router /medicine/page_list [get]
func (this MedicineController) PageList(c *gin.Context) {
	s := service.MedicineList{}
	if err := c.ShouldBindQuery(&s); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s.PageList())
	return
}

// @Summary 添加药品
// @Tags 药品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MedicineCreate false "查询参数"
// @Success 200 {object} response.Response
// @Router /medicine/create [post]
func (this MedicineController) Create(c *gin.Context) {
	s := service.MedicineCreate{}
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

// @Summary 修改药品
// @Tags 药品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MedicineUpdate false "查询参数"
// @Success 200 {object} response.Response
// @Router /medicine/update [post]
func (this MedicineController) Update(c *gin.Context) {
	s := service.MedicineUpdate{}
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

// @Summary 修改药品状态
// @Tags 药品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MedicineUpdateStatus false "查询参数"
// @Success 200 {object} response.Response
// @Router /medicine/update_status [post]
func (this MedicineController) UpdateStatus(c *gin.Context) {
	s := service.MedicineUpdateStatus{}
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

// @Summary 删除药品
// @Tags 药品
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param object body request.MedicineRemove false "查询参数"
// @Success 200 {object} response.Response
// @Router /medicine/remove [post]
func (this MedicineController) Remove(c *gin.Context) {
	s := service.MedicineRemove{}
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
