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
	key := fmt.Sprintf("reg_%v", this.Username)
	if this.Code != global.REDIS.Get(key).Val() {
		return nil, errors.New(lang.Lang("Verification code error"))
	}
	//检查邀请码
	puser := model.Member{
		Code: this.InviteCode,
	}
	if !puser.Get() {
		return nil, errors.New(lang.Lang("Wrong invitation code"))
	}

	//添加Redis乐观锁
	lockKey := fmt.Sprintf("redisLock:api:memberRegister:memberNamee_%s", this.Username)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	lockResult := redisLock.Lock(lockKey)
	if !lockResult {
		return nil, errors.New(lang.Lang("During data processing, Please try again later"))
	}

	//是否存在用户
	memberModel := model.Member{
		Username: this.Username,
	}
	if memberModel.Get() {
		//解锁
		redisLock.Unlock(lockKey)
		return nil, errors.New(lang.Lang("Username already exists"))
	}

	code := InviteCode()
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
		RegisterIp:       c.ClientIP(),
		LastLoginIp:      c.ClientIP(),
		LastLoginTime:    time.Now().Unix(),
		Code:             code,
		IsBuy:            2,
		IsOneShiming:     1,
	}
	err := member.Insert()
	if err != nil {
		//解锁
		redisLock.Unlock(lockKey)
		return nil, err
	}

	//绑定父级关系
	relation := make([]model.MemberRelation, 0)
	relation = append(relation, model.MemberRelation{
		UId:   member.Id,
		Puid:  member.Id,
		Level: 0,
	})
	//查询祖先
	prelation := model.MemberRelation{UId: puser.Id}
	pres, err := prelation.GetByUid()
	if err != nil {
		logrus.Errorf("绑定关系查询失败%v", err)
	}
	if len(pres) > 0 {
		for i := range pres {
			relation = append(relation, model.MemberRelation{
				UId:   member.Id,
				Puid:  pres[i].Puid,
				Level: pres[i].Level + 1,
			})
		}
	}
	err = prelation.InsertAll(relation)
	if err != nil {
		logrus.Errorf("绑定关系插入失败%v", err)
	}

	//解锁
	redisLock.Unlock(lockKey)

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

// 获取邀请码
func InviteCode() string {
	randCode := common.RandIntRunes(6)
	memberModel := model.Member{Code: randCode}
	//当邀请码重复时
	if !memberModel.Get() {
		return randCode
	}
	return InviteCode()
}
