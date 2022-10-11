package v1

import (
	"finance/model"
	"github.com/gin-gonic/gin"
)

type IPController struct {
	AuthController
}

func (this IPController) Index(c *gin.Context) {
	ip := model.IP{}
	if !ip.Get() {
		this.Json(c, 10001, "白名单不存在", nil)
		return
	}
	this.Json(c, 0, "ok", ip.IP)
	return
}
func (this IPController) Update(c *gin.Context) {
	s := struct {
		IP string `json:"ip"`
	}{}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	if s.IP == "" {
		this.Json(c, 10001, "ip不能为空", nil)
		return
	}
	ip := model.IP{}
	if !ip.Get() {
		this.Json(c, 10001, "ip不能为空", nil)
		return
	}
	ip.IP = s.IP
	if err = ip.Update("ip"); err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", nil)
	return
}
