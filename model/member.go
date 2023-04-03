package model

import (
	"finance/app/api/swag/response"
	"finance/common"
	"finance/extends"
	"finance/global"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Member struct {
	ID               int    `gorm:"column:id;primary_key"`          //
	Username         string `gorm:"column:username"`                //手机号
	Password         string `gorm:"column:password"`                //密码，sha1加密
	Salt             string `gorm:"column:salt"`                    //盐
	WithdrawPassword string `gorm:"column:withdraw_password"`       //提现密码
	WithdrawSalt     string `gorm:"column:withdraw_salt"`           //盐
	TotalBalance     int64  `gorm:"column:total_balance"`           //所有余额   可用余额+可提现余额+当前所投资项目的总额
	Balance          int64  `gorm:"column:balance"`                 //可用余额
	UseBalance       int64  `gorm:"column:use_balance"`             //可提现余额
	IsReal           int    `gorm:"column:is_real"`                 //是否实名 1审核中 2通过 3驳回
	RealName         string `gorm:"column:real_name"`               //真实姓名
	InvestFreeze     int64  `gorm:"column:invest_freeze"`           //余额宝冻结金额
	InvestAmount     int64  `gorm:"column:invest_amount"`           //余额宝有效金额
	InvestIncome     int64  `gorm:"column:invest_income"`           //余额宝总收益
	Avatar           string `gorm:"column:avatar"`                  //头像
	Status           int    `gorm:"column:status;default:1"`        //帐号启用状态，1启用2禁用
	FundsStatus      int    `gorm:"column:funds_status;default:1"`  //资金状态，1启用2禁用
	Level            int    `gorm:"column:level"`                   //等级
	Score            int    `gorm:"column:score;default:100"`       //信用分
	LastLoginTime    int64  `gorm:"column:last_login_time"`         //最后登录时间
	LastLoginIP      string `gorm:"column:last_login_ip"`           //最后登录ip
	RegTime          int64  `gorm:"column:reg_time;autoCreateTime"` //注册时间
	RegisterIP       string `gorm:"column:register_ip"`             //注册ip
	Token            string `gorm:"column:token"`                   //token盐
	Nickname         string `gorm:"column:nickname"`                //昵称
	Mobile           string `gorm:"column:mobile"`                  //手机号
	Email            string `gorm:"column:email"`                   //邮箱
	Qq               string `gorm:"column:qq"`                      //qq
	Wechat           string `gorm:"column:wechat"`                  //微信
	DisableLoginTime int64  `gorm:"column:disable_login_time"`      //禁止登录时间
	DisableBetTime   int64  `gorm:"column:disable_bet_time"`        //禁止投注时间
	Description      string `gorm:"column:description"`             //备注
	WithdrawAmount   int64  `gorm:"column:withdraw_amount"`         //提现流水
	Code             string `gorm:"column:code"`                    //邀请码
	IsBuy            int    `gorm:"column:is_buy"`                  //1=有效 2=无效
	IsOneShiming     int    `gorm:"column:is_one_shiming"`          //1=是 2=不是
	Income           int64  `gorm:"column:income"`                  //总收益
	PIncome          int64  `gorm:"column:p_income"`                //产品收益
	WillIncome       int64  `gorm:"column:wll_income"`              //待收益
	Guquan           int64  `gorm:"column:guquan"`                  //股权
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "c_member"
}
func (m *Member) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *Member) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Member) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Member) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Member) Info() *response.Member {
	//加密token
	jwtService := extends.JwtUtils{}
	token := jwtService.NewToken(this.ID, this.Token)
	var hasWithdrawPassword int
	if this.WithdrawPassword != "" {
		hasWithdrawPassword = 1
	}
	var mobile string
	if len(this.Mobile) >= 8 {
		mobile = this.Mobile[:4] + "****" + this.Mobile[len(this.Mobile)-4:]
	}
	coupon := MemberCoupon{Uid: int64(this.ID), IsUse: 1}
	list := coupon.List()
	var coupons []response.Coupon
	if len(list) > 0 {
		for i := range list {
			coupons = append(coupons, response.Coupon{
				UseId: list[i].ID,
				Id:    list[i].Coupon.ID,
				Price: float64(list[i].Coupon.Price) / UNITY,
			})
		}
	}
	where := "uid = ? or uid = ? and status = ? and is_read = ?"
	args := []interface{}{this.ID, -1, StatusOk, 1}
	msg := Message{}

	//金额分析(精确小数点后两位小数)
	incomeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.Income)/UNITY), 64)
	willIncomeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.WillIncome)/UNITY), 64)

	balanceAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.Balance)/UNITY), 64)
	useBalanceAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Floor(float64(this.UseBalance)/UNITY)), 64)
	totalBalanceAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.TotalBalance)/UNITY), 64)

	investFreezeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.InvestFreeze)/UNITY), 64)
	investAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.InvestAmount)/UNITY), 64)
	investIncomeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.InvestIncome)/UNITY), 64)

	return &response.Member{
		ID:                  this.ID,
		Username:            this.Username,
		Balance:             balanceAmount,
		UseBalance:          useBalanceAmount,
		IsReal:              this.IsReal,
		RealName:            this.RealName,
		InvestFreeze:        investFreezeAmount,
		InvestAmount:        investAmount,
		InvestIncome:        investIncomeAmount,
		Avatar:              this.Avatar,
		Status:              this.Status,
		FundsStatus:         this.FundsStatus,
		Level:               this.Level,
		Score:               this.Score,
		LastLoginTime:       this.LastLoginTime,
		LastLoginIP:         this.LastLoginIP,
		RegTime:             this.RegTime,
		RegisterIP:          this.RegisterIP,
		Token:               token,
		HasWithdrawPassword: hasWithdrawPassword,
		Nickname:            this.Nickname,
		Mobile:              mobile,
		Email:               this.Email,
		Qq:                  this.Qq,
		Wechat:              this.Wechat,
		InviteCode:          this.Code,
		TotalBalance:        totalBalanceAmount,
		Coupon:              coupons,
		Income:              incomeAmount,
		Guquan:              this.Guquan,
		Message:             msg.Count(where, args),
		WillIncome:          willIncomeAmount,
	}
}

// Get the member list from agent
func (m *Member) GetPageList(where string, args []interface{}, page, pageSize int) ([]Member, common.Page) {
	res := make([]Member, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(m).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Member) List(where string, args []interface{}) []Member {
	res := make([]Member, 0)
	tx := global.DB.Model(this).Where(where, args...).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *Member) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (this *Member) Count(where string, args []interface{}) int64 {
	var t int64
	res := global.DB.Model(this).Where(where, args...).Count(&t)
	if res.Error != nil {
		return 0
	}
	return t
}

func (this *Member) Sum(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
