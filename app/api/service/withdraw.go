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
	"github.com/shopspring/decimal"

	//"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type WithdrawList struct {
	request.WithdrawList
}

func (this WithdrawList) PageList(member model.Member) response.WithdrawData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Withdraw{}
	where, args, _ := this.getWhere(member.Id)
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	res := make([]response.Withdraw, 0)
	for _, v := range list {
		if v.WithdrawMethod.Code == "usdt" {
			v.Amount = v.UsdtAmount
		}
		i := response.Withdraw{
			Id:          v.Id,
			OrderSn:     v.OrderSn,
			Type:        v.WithdrawType,
			TypeName:    v.WithdrawMethod.Name,
			BankName:    v.BankName,
			BranchBank:  v.BranchBank,
			RealName:    v.RealName,
			CardNumber:  v.CardNumber,
			BankPhone:   v.BankPhone,
			Amount:      v.Amount,
			Fee:         v.Fee,
			TotalAmount: v.TotalAmount,
			Description: v.Description,
			Status:      v.Status,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}

		res = append(res, i)
	}
	return response.WithdrawData{List: res, Page: FormatPage(page)}
}

func (this WithdrawList) getWhere(uid int) (string, []interface{}, error) {
	where := map[string]interface{}{
		"uid": uid,
	}
	if this.Status > 0 {
		where[model.Withdraw{}.TableName()+".status"] = this.Status
		//where["o.draw_time >"] = time.Now().Unix()
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals, nil
}

type WithdrawCreate struct {
	request.WithdrawCreate
}

func (this WithdrawCreate) Create(member model.Member) error {
	//添加Redis乐观锁
	lockKey := fmt.Sprintf("withdraw:%d", member.Id)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	if !redisLock.Lock(lockKey) {
		return errors.New(lang.Lang("During data processing, Please try again later"))
	}
	defer redisLock.Unlock(lockKey)
	var err error
	tx := global.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	//用户提现状态
	if member.FundsStatus != model.StatusOk {
		return errors.New(lang.Lang("The current account has been frozen!"))
	}
	//提现金额
	if this.TotalAmount.LessThanOrEqual(decimal.Zero) {
		return errors.New(lang.Lang("The withdrawal amount format is incorrect"))
	}
	//提现方式
	if this.Method == 0 {
		return errors.New(lang.Lang("Wrong withdrawal method"))
	}
	//银行卡Id
	if this.Id == 0 {
		return errors.New("提款方式错误")
	}
	//验证交易密码
	if common.Md5String(this.WithdrawPassword+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New(lang.Lang("Incorrect withdraw password"))
	}
	//提现方式分析
	method := model.WithdrawMethod{Id: this.Method}
	if !method.Get() {
		return errors.New(lang.Lang("Wrong withdrawal method"))
	}
	config := model.SetBase{}
	config.Get()
	//提现时间 金额验证
	c := model.SetFunds{}
	if c.Get() {
		now := time.Now().Unix()
		if now < common.TimeToUnix(c.WithdrawStartTime) {
			return errors.New(fmt.Sprintf(lang.Lang("Please withdraw after %s"), c.WithdrawStartTime))
		}
		if now > common.TimeToUnix(c.WithdrawEndTime) {
			return errors.New(fmt.Sprintf(lang.Lang("Please withdraw before %s"), c.WithdrawEndTime))
		}
		if method.Code == "bank" {
			if this.TotalAmount.LessThan(c.WithdrawMinAmount) {
				return errors.New(fmt.Sprintf(lang.Lang("Minimum withdraw %v"), c.WithdrawMinAmount))
			}
			if c.WithdrawMaxAmount.LessThan(this.TotalAmount) {
				return errors.New(fmt.Sprintf(lang.Lang("Maximum withdraw %v"), c.WithdrawMaxAmount))
			}
		} else if method.Code == "usdt" {
			if this.TotalAmount.LessThan(c.WithdrawMinAmount.Div(config.UsdtBuyRate)) {
				return errors.New(fmt.Sprintf(lang.Lang("Minimum withdraw %v"), c.WithdrawMinAmount.Div(config.UsdtBuyRate).Round(2)))
			}
			if c.WithdrawMaxAmount.Div(config.UsdtBuyRate).LessThan(this.TotalAmount) {
				return errors.New(fmt.Sprintf(lang.Lang("Maximum withdraw %v"), c.WithdrawMaxAmount.Div(config.UsdtBuyRate).Round(2)))
			}
		}
	}

	//股权分开启 验证额度
	if time.Now().Unix() >= config.EquityStartDate {
		equityScore := model.EquityScoreOrder{}
		score := equityScore.SumScore("uid = ? and status = ? and create_time < ?", []interface{}{member.Id, model.StatusOk, common.GetTodayZero()}, "pay_money")
		threshold := config.EquityRate.Mul(score).Div(decimal.NewFromInt(100)).Round(2)
		if method.Code == "bank" {
			if threshold.LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		} else if method.Code == "usdt" {
			if threshold.Div(config.UsdtBuyRate).LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		}
		//当日已使用额度
		sumModel := model.Withdraw{}
		sumWhere := "uid = ? and create_time >= ? and status in (?)"
		sumArgs := []interface{}{member.Id, common.GetTodayZero(), []int{model.StatusReview, model.StatusAccept}}
		sum := sumModel.Sum(sumWhere, sumArgs, "total_amount")

		if method.Code == "bank" {
			if threshold.Sub(decimal.NewFromFloat(sum)).LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		} else if method.Code == "usdt" {
			if threshold.Sub(decimal.NewFromFloat(sum)).Div(config.UsdtBuyRate).LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		}
	} else {
		if method.Code == "bank" {
			if member.WithdrawThreshold.LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		} else if method.Code == "usdt" {
			if member.WithdrawThreshold.Div(config.UsdtBuyRate).LessThan(this.TotalAmount) {
				return errors.New("提现额度不足")
			}
		}
	}
	//每日提现次数
	countModel := model.Withdraw{}
	countWhere := "uid = ? and create_time >= ? and `c_withdraw`.`status` != ?"
	countArgs := []interface{}{member.Id, common.GetTodayZero(), model.StatusRollback}
	count := countModel.Count(countWhere, countArgs)
	if count > 0 && count >= int64(c.WithdrawCount) {
		return errors.New(fmt.Sprintf(lang.Lang("You can only withdraw %d times per day"), c.WithdrawCount))
	}
	//当月未参与投资，不允许提现
	//exist := model.OrderProduct{UId: member.Id}
	//if exist.Get() {
	//	if exist.CreateTime+30*86400 < time.Now().Unix() {
	//		return errors.New("30天内未激活账户，不允许提现")
	//	}
	//} else {
	//	return errors.New("30天内未激活账户，不允许提现")
	//}
	//计算手续费
	fee := c.WithdrawFee.Mul(this.TotalAmount).Div(decimal.NewFromInt(100)).Round(2)
	var bankName string
	var branchBank string
	var realName string
	var cardNumber string
	var bankPhone string
	var usdtAmount decimal.Decimal
	var realAmount decimal.Decimal
	var totalAmount decimal.Decimal
	switch method.Code {
	case "bank":
		//检查余额
		if member.WithdrawBalance.LessThan(this.TotalAmount) {
			return errors.New(lang.Lang("Insufficient account balance"))
		}
		bank := model.MemberBank{Id: this.Id}
		if !bank.Get() {
			return errors.New("银行卡不存在")
		}
		bankName = bank.BankName
		branchBank = bank.BranchBank
		realName = bank.RealName
		cardNumber = bank.CardNumber
		bankPhone = bank.BankPhone
		realAmount = this.TotalAmount.Sub(fee)
		totalAmount = this.TotalAmount
	case "usdt":
		//检查余额
		if member.UsdtWithdrawBalance.LessThan(this.TotalAmount) {
			return errors.New(lang.Lang("Insufficient account balance"))
		}
		usdt := model.MemberUsdt{Id: this.Id}
		if !usdt.Get() {
			return errors.New("usdt地址不存在")
		}
		bankName = usdt.Protocol
		cardNumber = usdt.Address
		totalAmount = this.TotalAmount.Mul(config.UsdtBuyRate)
		realAmount = this.TotalAmount.Sub(fee)
		usdtAmount = this.TotalAmount
		//usdtAmount = this.TotalAmount.Sub(fee).Div(config.UsdtRate).Round(2)
		//realAmount = this.TotalAmount.Div(config.UsdtRate).Round(2)
	}
	//生成提现记录
	order := model.Withdraw{
		UId:          member.Id,
		WithdrawType: this.Method,
		BankName:     bankName,
		BranchBank:   branchBank,
		RealName:     realName,
		CardNumber:   cardNumber,
		BankPhone:    bankPhone,
		Amount:       realAmount,
		Fee:          fee,
		TotalAmount:  totalAmount,
		Status:       model.StatusReview,
		OrderSn:      common.OrderSn(),
		UsdtAmount:   usdtAmount,
	}
	if err := tx.Create(&order).Error; err != nil {
		//解锁
		//redisLock.Unlock(lockKey)
		return err
	}
	//生成账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: 4,
		ItemId:    order.Id,
		Amount:    this.TotalAmount,
		Before:    member.WithdrawBalance,
		After:     member.WithdrawBalance.Sub(this.TotalAmount),
		Desc:      "提现申请",
	}
	if method.Code == "usdt" {
		trade.Before = member.UsdtWithdrawBalance
		trade.After = member.UsdtWithdrawBalance.Sub(this.TotalAmount)
	}
	if err := tx.Create(&trade).Error; err != nil {
		//解锁
		//redisLock.Unlock(lockKey)
		return err
	}
	//更改用户余额
	if method.Code == "bank" {
		member.WithdrawBalance = member.WithdrawBalance.Sub(this.TotalAmount)
		member.WithdrawThreshold = member.WithdrawThreshold.Sub(this.TotalAmount)
	} else if method.Code == "usdt" {
		member.UsdtWithdrawBalance = member.UsdtWithdrawBalance.Sub(this.TotalAmount)
		member.WithdrawThreshold = member.WithdrawThreshold.Sub(this.TotalAmount.Mul(config.UsdtBuyRate))
	}
	return tx.Select("withdraw_balance", "usdt_withdraw_balance", "withdraw_threshold").Updates(member).Error
}

type WithdrawMethod struct {
}

func (this WithdrawMethod) List() *response.WithdrawMethodData {
	m := model.WithdrawMethod{
		Lang:   global.Language,
		Status: model.StatusOk,
	}
	methods, err := m.List()
	if err != nil {
		return nil
	}
	return &response.WithdrawMethodData{
		List: this.formatList(methods),
	}
}
func (this WithdrawMethod) formatList(lists []model.WithdrawMethod) []response.WithdrawMethod {
	res := make([]response.WithdrawMethod, 0)
	for _, v := range lists {
		i := response.WithdrawMethod{
			Id:   v.Id,
			Name: v.Name,
			Code: v.Code,
			Icon: v.Icon,
			//Fee:  float64(v.Fee) ,
		}
		res = append(res, i)
	}
	return res
}
