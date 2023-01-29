package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/common"
	"finance/global"
	"finance/lang"
	"finance/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type LoginService struct {
	request.Login
}
type RegisterService struct {
	request.Register
}

func (this LoginService) DoLogin(c *gin.Context) (*response.Member, error) {

	if this.Username == "" {
		return nil, errors.New(lang.Lang("Username cannot be empty"))
	}
	if !common.IsMobile(this.Username, global.Language) {
		return nil, errors.New("手机格式不正确")
	}
	if this.Password == "" {
		return nil, errors.New(lang.Lang("Password cannot be empty"))
	}
	//if this.Code == "" {
	//	return nil, errors.New("验证码不能为空")
	//}
	//if !common.CaptchaVerify(c, this.Code) {
	//	return nil, errors.New("验证码错误")
	//}

	//是否存在用户
	member := model.Member{
		Username: this.Username,
	}
	if !member.Get() {
		return nil, errors.New(lang.Lang("Username does not exist"))
	}

	if member.Password != common.Md5String(this.Password+member.Salt) {
		return nil, errors.New(lang.Lang("Incorrect username and password"))
	}
	if member.Status != model.StatusOk {
		return nil, errors.New(lang.Lang("User is forbidden to log in"))
	}

	ip := c.ClientIP()
	//用户登录日志
	go memberLoginLog(member, ip)
	//重置token
	member.LastLoginIP = ip
	member.LastLoginTime = time.Now().Unix()
	member.Token = common.RandStringRunes(32)
	return member.Info(), member.Update("token", "last_login_ip", "last_login_time")
}

func (this RegisterService) Insert(c *gin.Context) (*response.Member, error) {

	key := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(this.Username)
	value := c.ClientIP() + c.Request.URL.Path + fmt.Sprint(this.Username)
	if global.REDIS.Get(key).Val() == value {
		return nil, errors.New("请勿重复提交表单")
	}
	global.REDIS.Set(key, value, time.Second*3)

	if this.Username == "" {
		return nil, errors.New(lang.Lang("Username cannot be empty"))
	}
	if !common.IsMobile(this.Username, global.Language) {
		return nil, errors.New("手机格式不正确")
	}
	if this.Password == "" {
		return nil, errors.New(lang.Lang("Password cannot be empty"))
	}
	if this.RePassword != this.Password {
		return nil, errors.New(lang.Lang("The two passwords are inconsistent"))
	}
	if this.WithdrawPassword == "" {
		return nil, errors.New(lang.Lang("Withdraw password cannot be empty"))
	}
	if this.InviteCode == "" {
		return nil, errors.New(lang.Lang("Invitation code cannot be empty"))
	}
	//验证码
	key = fmt.Sprintf("reg_%v", this.Username)
	if this.Code != global.REDIS.Get(key).Val() {
		return nil, errors.New("验证码错误")
	}
	//if this.Code == "" {
	//	return nil, errors.New("验证码不能为空")
	//}
	//if !common.CaptchaVerify(c, this.Code) {
	//	return nil, errors.New("验证码错误")
	//}
	//是否存在用户
	m := model.Member{
		Username: this.Username,
	}
	if m.Get() {
		return nil, errors.New(lang.Lang("Username already exists"))
	}
	//检查邀请码
	puser := model.Member{
		Code: this.InviteCode,
	}
	if !puser.Get() {
		return nil, errors.New(lang.Lang("Wrong invitation code"))
	}
	code := common.RandIntRunes(6)
	salt := common.RandStringRunes(6)
	withdrawSalt := common.RandStringRunes(6)
	//入库
	member := model.Member{
		Username:         this.Username,
		Salt:             salt,
		Password:         common.Md5String(this.Password + salt),
		WithdrawSalt:     withdrawSalt,
		WithdrawPassword: common.Md5String(this.WithdrawPassword + withdrawSalt),
		Token:            common.RandStringRunes(32),
		RegisterIP:       c.ClientIP(),
		LastLoginIP:      c.ClientIP(),
		LastLoginTime:    time.Now().Unix(),
		Code:             code,
		IsBuy:            2,
		IsOneShiming:     1,
	}
	err := member.Insert()
	if err != nil {
		return nil, err
	}
	//绑定父级关系
	relation := make([]model.MemberRelation, 0)
	relation = append(relation, model.MemberRelation{
		UID:   member.ID,
		Puid:  member.ID,
		Level: 0,
	})
	//查询祖先
	prelation := model.MemberRelation{UID: puser.ID}
	pres, err := prelation.GetByUid()
	if err != nil {
		logrus.Errorf("绑定关系查询失败%v", err)
	}
	if len(pres) > 0 {
		for i := range pres {
			relation = append(relation, model.MemberRelation{
				UID:   member.ID,
				Puid:  pres[i].Puid,
				Level: pres[i].Level + 1,
			})
		}
	}
	err = prelation.InsertAll(relation)
	if err != nil {
		logrus.Errorf("绑定关系插入失败%v", err)
	}
	go memberLoginLog(member, c.ClientIP())
	return member.Info(), nil
}

func memberLoginLog(member model.Member, ip string) {
	m := model.MemberLoginLog{
		UID:       member.ID,
		Username:  member.Username,
		LoginIP:   ip,
		LoginTime: time.Now().Unix(),
	}
	m.Insert()
}
