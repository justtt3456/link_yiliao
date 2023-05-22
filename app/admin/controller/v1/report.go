package v1

import (
	"china-russia/app/admin/service"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	AuthController
}

// Summary 用户报表
// Tags 报表
// Param token header string false "用户令牌"
// Param object query request.ReportSum false "查询参数"
// Success 200 {object} response.MemberReportResponse
// Router /report/member [get]
func (this ReportController) Member(c *gin.Context) {
	s := service.Report{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res := s.Member()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}

// Summary 代理报表
// Tags 报表
// Param token header string false "用户令牌"
// Param object query request.ReportSum false "查询参数"
// Success 200 {object} response.AgentReportResponse
// Router /report/agent [get]
func (this ReportController) Agent(c *gin.Context) {
	s := service.Report{}
	err := c.ShouldBindQuery(&s)
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	res := s.Agent()
	if err != nil {
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", res)
	return
}
