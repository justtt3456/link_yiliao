package v1

import (
	"finance/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (this *Controller) Json(c *gin.Context, code int, msg string, data interface{}) {
	type res struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	if data != nil {
		c.JSON(http.StatusOK, res{
			Code: code,
			Msg:  msg,
			Data: data,
		})
	} else {
		c.JSON(http.StatusOK, res{
			Code: code,
			Msg:  msg,
			Data: nil,
		})
	}
	return
}

type AuthController struct {
	Controller
}

func (AuthController) AdminInfo(c *gin.Context) *model.Admin {
	res, b := c.Get("admin")
	if !b {
		return nil
	}
	if claims, ok := res.(model.Admin); ok {
		return &claims
	}
	return nil
}
