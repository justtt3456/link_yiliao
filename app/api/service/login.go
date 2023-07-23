package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
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
	member.LastLoginIp = ip
	member.LastLoginTime = time.Now().Unix()
	member.Token = common.RandStringRunes(32)
	return member.Info(), member.Update("token", "last_login_ip", "last_login_time")
}

func (this RegisterService) Insert(c *gin.Context) (*response.Member, error) {
	//参数分析
	if this.Username == "" {
		return nil, errors.New(lang.Lang("Username cannot be empty"))
	}
	//添加Redis乐观锁
	lockKey := fmt.Sprintf("member_register:%s", this.Username)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	if !redisLock.Lock(lockKey) {
		return nil, errors.New(lang.Lang("During data processing, Please try again later"))
	}
	defer redisLock.Unlock(lockKey)
	//同一ip注册数量限制
	ex := model.Member{}
	log.Println("当前ip:", c.ClientIP())
	if ex.Count("register_ip = ?", []interface{}{c.ClientIP()}) >= 3 {
		return nil, errors.New("同一ip注册数量不能超过3个")
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
	//if this.WithdrawPassword == "" {
	//	return nil, errors.New(lang.Lang("Withdraw password cannot be empty"))
	//}
	if this.InviteCode == "" {
		return nil, errors.New(lang.Lang("Invitation code cannot be empty"))
	}
	//验证码
	//if !common.CaptchaVerify(c, this.Code) {
	//	return nil, errors.New("验证码错误")
	//}
	//检查邀请码
	invite := model.InviteCode{
		Code: this.InviteCode,
	}
	if !invite.Get() {
		return nil, errors.New(lang.Lang("Wrong invitation code"))
	}
	//是否存在用户
	memberModel := model.Member{
		Username: this.Username,
	}
	if memberModel.Get() {
		return nil, errors.New(lang.Lang("Username already exists"))
	}
	salt := common.RandStringRunes(6)
	password := common.Md5String(this.Password + salt)
	//入库
	member := model.Member{
		Username:         this.Username,
		Salt:             salt,
		WithdrawSalt:     salt,
		WithdrawPassword: password,
		Password:         password,
		Token:            common.RandStringRunes(32),
		RegisterIp:       c.ClientIP(),
		LastLoginIp:      c.ClientIP(),
		LastLoginTime:    time.Now().Unix(),
		Status:           model.StatusOk,
		FundsStatus:      model.StatusOk,
		IsBuy:            2,
		AgentId:          invite.AgentId,
	}
	err := member.Insert()
	if err != nil {
		return nil, err
	}
	inviteCode := model.InviteCode{}
	inviteCode.Code = inviteCode.InviteCode()
	inviteCode.UId = member.Id
	inviteCode.Username = member.Username
	inviteCode.AgentId = invite.AgentId
	inviteCode.AgentName = invite.AgentName
	inviteCode.Insert()
	//三级代理
	if invite.UId > 0 {
		current := model.MemberParents{
			Uid:      member.Id,
			ParentId: invite.UId,
			Level:    1,
		}
		current.Insert()
	}
	parent := model.MemberParents{Uid: invite.UId}
	pres, err := parent.GetByUid()
	if err != nil {
		logrus.Errorf("绑定关系查询失败%v", err)
	}
	parents := make([]model.MemberParents, 0)
	for _, v := range pres {
		//只做三级
		//if v.Level >= 3 {
		//	break
		//}
		if v.ParentId <= 0 {
			break
		}
		parents = append(parents, model.MemberParents{
			Uid:      member.Id,
			ParentId: v.ParentId,
			Level:    v.Level + 1,
		})
	}
	err = parent.InsertAll(parents)
	if err != nil {
		logrus.Errorf("绑定关系插入失败%v", err)
	}

	go memberLoginLog(member, c.ClientIP())
	return member.Info(), nil
}

func memberLoginLog(member model.Member, ip string) {
	m := model.MemberLoginLog{
		UId:       member.Id,
		Username:  member.Username,
		LoginIP:   ip,
		LoginTime: time.Now().Unix(),
	}
	m.Insert()
}
