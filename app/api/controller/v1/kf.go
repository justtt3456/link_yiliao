package v1

import (
	"finance/app/api/controller"
	"finance/lang"
	"finance/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type KfController struct {
	controller.Controller
}

func (this KfController) Redirect(c *gin.Context) {
	var param struct {
		ID  int `form:"id"`
		UID int `form:"uid"`
	}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	if param.ID == 0 {
		param.ID = 2
		//c.String(http.StatusOK, lang.Lang("Parameter error"))
	}
	//if param.UID == 0 {
	//	c.String(http.StatusOK, lang.Lang("Configuration error"))
	//	return
	//}
	config := model.SetKf{ID: param.ID}
	if !config.Get() {
		c.String(http.StatusOK, lang.Lang("Configuration error"))
		return
	}

	member := model.Member{ID: param.UID}
	if param.UID > 0 {
		if !member.Get() {
			//c.String(http.StatusOK, lang.Lang("User is not logged in"))
			//return
		}
	}

	var agentName string

	//检查用户认证 传参到客服
	c.HTML(http.StatusOK, "kf_"+strconv.Itoa(param.ID)+".html", gin.H{
		"config":    config,
		"member":    member,
		"agentName": agentName,
	})
	return
}
