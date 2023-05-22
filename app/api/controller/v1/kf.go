package v1

import (
	"china-russia/app/api/controller"
	"china-russia/lang"
	"china-russia/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type KfController struct {
	controller.Controller
}

func (this KfController) Redirect(c *gin.Context) {
	var param struct {
		Id  int `form:"id"`
		UId int `form:"uid"`
	}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	if param.Id == 0 {
		param.Id = 2
		//c.String(http.StatusOK, lang.Lang("Parameter error"))
	}
	//if param.UId == 0 {
	//	c.String(http.StatusOK, lang.Lang("Configuration error"))
	//	return
	//}
	config := model.SetKf{Id: param.Id}
	if !config.Get() {
		c.String(http.StatusOK, lang.Lang("Configuration error"))
		return
	}

	member := model.Member{Id: param.UId}
	if param.UId > 0 {
		if !member.Get() {
			//c.String(http.StatusOK, lang.Lang("User is not logged in"))
			//return
		}
	}

	var agentName string

	//检查用户认证 传参到客服
	c.HTML(http.StatusOK, "kf_"+strconv.Itoa(param.Id)+".html", gin.H{
		"config":    config,
		"member":    member,
		"agentName": agentName,
	})
	return
}
