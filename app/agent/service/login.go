package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	request.LoginRequest
}

func (this LoginService) Login(c *gin.Context) (*model.Agent, error) {
	if this.Username == "" {
		return nil, errors.New("账号不能为空")
	}
	if this.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	//if this.GoogleCode == "" {
	//	return nil, errors.New("验证码不能为空")
	//}
	//是否存在用户
	agent := model.Agent{
		Account: this.Username,
	}
	if !agent.Get() {
		return nil, errors.New("账号不存在")
	}
	//验证谷歌验证码
	//google := extends.NewGoogleAuth()
	//b, err := google.VerifyCode(admin.GoogleAuth, this.GoogleCode)
	//if err != nil {
	//	return nil, errors.New("验证码错误")
	//}
	//if !b {
	//	return nil, errors.New("验证码错误")
	//}
	//验证密码
	if agent.Password != common.Md5String(this.Password+agent.Salt) {
		return nil, errors.New("账号密码错误")
	}
	//重置token
	agent.Token = common.RandStringRunes(32)
	agent.Update("token")
	return &agent, nil
}
