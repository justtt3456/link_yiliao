package model

import (
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/extends"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type Member struct {
	Id                int             `gorm:"column:id"`
	Username          string          `gorm:"column:username"`          //手机号
	Password          string          `gorm:"column:password"`          //密码，sha1加密
	Salt              string          `gorm:"column:salt"`              //盐
	WithdrawPassword  string          `gorm:"column:withdraw_password"` //提现密码
	WithdrawSalt      string          `gorm:"column:withdraw_salt"`
	AgentId           int             `gorm:"column:agent_id"`           //代理id
	Balance           decimal.Decimal `gorm:"column:balance"`            //可用余额
	WithdrawBalance   decimal.Decimal `gorm:"column:withdraw_balance"`   //可提现余额
	IsReal            int             `gorm:"column:is_real"`            //是否实名 1审核中 2通过 3驳回
	RealName          string          `gorm:"column:real_name"`          //真实姓名
	InvestFreeze      decimal.Decimal `gorm:"column:invest_freeze"`      //余额宝冻结金额
	InvestAmount      decimal.Decimal `gorm:"column:invest_amount"`      //余额宝有效金额
	InvestIncome      decimal.Decimal `gorm:"column:invest_income"`      //余额宝总收益
	Avatar            string          `gorm:"column:avatar"`             //头像
	Status            int             `gorm:"column:status"`             //帐号启用状态，1启用2禁用
	FundsStatus       int             `gorm:"column:funds_status"`       //资金冻结状态
	Level             int             `gorm:"column:level"`              //等级
	Score             int             `gorm:"column:score"`              //信誉分
	LastLoginTime     int64           `gorm:"column:last_login_time"`    //最后登录时间
	LastLoginIp       string          `gorm:"column:last_login_ip"`      //最后登录ip
	RegTime           int64           `gorm:"column:reg_time"`           //注册时间
	RegisterIp        string          `gorm:"column:register_ip"`        //注册ip
	Token             string          `gorm:"column:token"`              //token盐
	DisableLoginTime  int64           `gorm:"column:disable_login_time"` //禁止登录时间
	DisableBetTime    int64           `gorm:"column:disable_bet_time"`   //禁止投注时间
	WithdrawAmount    int64           `gorm:"column:withdraw_amount"`    //提现流水
	Description       string          `gorm:"column:description"`        //用户备注
	IsBuy             int             `gorm:"column:is_buy"`             //1=有效 2=无效
	Equity            int             `gorm:"column:equity"`             //股权
	EquityScore       int             `gorm:"column:equity_score"`       //股权分
	PreIncome         decimal.Decimal `gorm:"column:pre_income"`         //待收益
	PreCapital        decimal.Decimal `gorm:"column:pre_capital"`        //待收本金
	TotalRebate       decimal.Decimal `gorm:"column:total_rebate"`       //总返佣
	TotalRecharge     decimal.Decimal `gorm:"column:total_recharge"`     //总充值
	TotalIncome       decimal.Decimal `gorm:"column:total_income"`       //总收益
	WithdrawThreshold decimal.Decimal `gorm:"column:withdraw_threshold"` //提现额度
	TotalBuy          decimal.Decimal `gorm:"column:total_buy"`
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
	token := jwtService.NewToken(this.Id, this.Token)
	var hasWithdrawPassword int
	if this.WithdrawPassword != "" {
		hasWithdrawPassword = 1
	}
	var username string
	if len(this.Username) >= 8 {
		username = this.Username[:3] + "****" + this.Username[len(this.Username)-4:]
	}
	invite := InviteCode{
		UId: this.Id,
	}
	invite.Get()
	coupon := MemberCoupon{Uid: this.Id, IsUse: 1}
	list := coupon.List()
	var coupons []response.Coupon
	if len(list) > 0 {
		for _, v := range list {
			coupons = append(coupons, response.Coupon{
				UseId: v.Id,
				Id:    v.Coupon.Id,
				Price: v.Coupon.Price,
			})
		}
	}
	//实名信息
	var mobile string
	mv := MemberVerified{UId: this.Id}
	mv.Get()
	if len(mv.Mobile) >= 8 {
		mobile = mv.Mobile[:3] + "****" + mv.Mobile[len(mv.Mobile)-4:]
	}
	where := "uid = ? and is_read = ?"
	args := []interface{}{this.Id, StatusClose}
	msg := MemberMessage{}
	//股权分开启提现额度修改
	config := SetBase{}
	config.Get()
	if time.Now().Unix() >= config.EquityStartDate {
		equityScore := EquityScoreOrder{}
		score := equityScore.SumScore("uid = ? and status = ? and create_time < ?", []interface{}{this.Id, StatusOk, common.GetTodayZero()}, "pay_money")
		this.WithdrawThreshold = config.EquityRate.Mul(score).Div(decimal.NewFromInt(100)).Round(2)
	}
	return &response.Member{
		Id:                  this.Id,
		Username:            username,
		Balance:             this.Balance,
		WithdrawBalance:     this.WithdrawBalance,
		IsReal:              this.IsReal,
		RealName:            mv.RealName,
		IdNumber:            mv.IdNumber,
		Mobile:              mobile,
		InvestFreeze:        this.InvestFreeze,
		InvestAmount:        this.InvestAmount,
		InvestIncome:        this.InvestIncome,
		Avatar:              this.Avatar,
		Status:              this.Status,
		FundsStatus:         this.FundsStatus,
		Level:               this.Level,
		Score:               this.Score,
		LastLoginTime:       this.LastLoginTime,
		LastLoginIP:         this.LastLoginIp,
		RegTime:             this.RegTime,
		RegisterIP:          this.RegisterIp,
		Token:               token,
		HasWithdrawPassword: hasWithdrawPassword,
		InviteCode:          invite.Code,
		Coupon:              coupons,
		//Income: this.TotalIncome,
		//Guquan:  this.Guquan,
		Message:           msg.Count(where, args),
		PreIncome:         this.PreIncome,
		PreCapital:        this.PreCapital,
		WithdrawThreshold: this.WithdrawThreshold,
		EquityScore:       this.EquityScore,
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

func (this *Member) Sum(where string, args []interface{}, field string) float64 {
	var total float64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
