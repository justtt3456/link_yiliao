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
		//提现金额
		//amount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.Amount)/model.UNITY), 64)
		//freeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.Fee)/model.UNITY), 64)
		//totalAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.TotalAmount)/model.UNITY), 64)

		i := response.Withdraw{
			Id:         v.Id,
			OrderSn:    v.OrderSn,
			Type:       v.WithdrawType,
			TypeName:   v.WithdrawMethod.Name,
			BankName:   v.BankName,
			BranchBank: v.BranchBank,
			RealName:   v.RealName,
			CardNumber: v.CardNumber,
			BankPhone:  v.BankPhone,
			//Amount:      amount,
			//Fee:         freeAmount,
			//TotalAmount: totalAmount,
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
	//用户提现状态
	if member.FundsStatus != model.StatusOk {
		return errors.New(lang.Lang("The current account has been frozen!"))
	}
	//提金额
	//if this.TotalAmount <= 0 {
	//	return errors.New(lang.Lang("The withdrawal amount format is incorrect"))
	//}
	//提现方式
	if this.Method == 0 {
		return errors.New(lang.Lang("Wrong withdrawal method"))
	}
	//银行卡Id
	if this.Id == 0 {
		return errors.New(lang.Lang("Bank card does not exist"))
	}
	memberBank := model.MemberBank{Id: this.Id}
	if !memberBank.Get() {
		return errors.New(lang.Lang("Bank card does not exist"))
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

	//amount := int64(this.TotalAmount)
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
		//if amount < c.WithdrawMinAmount {
		//	return errors.New(fmt.Sprintf(lang.Lang("Minimum withdraw %.2f"), float64(c.WithdrawMinAmount)/model.UNITY))
		//}
		//if amount > c.WithdrawMaxAmount {
		//	return errors.New(fmt.Sprintf(lang.Lang("Maximum withdraw %.2f"), float64(c.WithdrawMaxAmount)/model.UNITY))
		//}
	}

	//添加Redis乐观锁
	//lockKey := fmt.Sprintf("redisLock:api:submitWithdrawOrder:memberId_%s_methodId_%d_amount_%d", member.Id, this.Method, amount)
	//redisLock := common.RedisLock{RedisClient: global.REDIS}
	//lockResult := redisLock.Lock(lockKey)
	//if !lockResult {
	//	return errors.New(lang.Lang("During data processing, Please try again later"))
	//}

	//获取用户当前余额
	memberModel := model.Member{Id: member.Id}
	if !memberModel.Get() {
		//解锁
		//redisLock.Unlock(lockKey)
		return errors.New(lang.Lang("Member info is not exits"))
	}
	//检查余额
	//if memberModel.WithdrawBalance < amount {
	//	//解锁
	//	redisLock.Unlock(lockKey)
	//	return errors.New(lang.Lang("Insufficient account balance"))
	//}

	//每日提现次数
	countModel := model.Withdraw{}
	countWhere := "uid = ? and create_time >= ? and `c_withdraw`.`status` != ?"
	countArgs := []interface{}{member.Id, common.GetTodayZero(), model.StatusRollback}
	count := countModel.Count(countWhere, countArgs)
	if count > 0 && count >= int64(c.WithdrawCount) {
		//解锁
		//redisLock.Unlock(lockKey)
		return errors.New(fmt.Sprintf(lang.Lang("You can only withdraw %d times per day"), c.WithdrawCount))
	}
	//当月未参与投资，不允许提现
	exist := model.OrderProduct{UId: member.Id}
	if exist.Get() {
		if exist.CreateTime+30*86400 < time.Now().Unix() {
			return errors.New("30天内未激活账户，不允许提现")
		}
	} else {
		return errors.New("30天内未激活账户，不允许提现")
	}
	//计算手续费
	//fee := int64(c.WithdrawFee) * amount / int64(model.UNITY)
	//生成提现记录
	order := model.Withdraw{
		UId:          member.Id,
		WithdrawType: this.Method,
		BankName:     memberBank.BankName,
		BranchBank:   memberBank.BranchBank,
		RealName:     memberBank.RealName,
		CardNumber:   memberBank.CardNumber,
		BankPhone:    memberBank.BankPhone,
		//Amount:       amount - fee,
		//Fee:          fee,
		//TotalAmount:  amount,
		Status:  model.StatusReview,
		OrderSn: common.OrderSn(),
	}
	if err := order.Insert(); err != nil {
		//解锁
		//redisLock.Unlock(lockKey)
		return err
	}
	//生成账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: 4,
		ItemId:    order.Id,
		//Amount:    amount,
		Before: memberModel.WithdrawBalance,
		//After:     memberModel.WithdrawBalance - amount,
		Desc: "提现申请",
	}
	if err := trade.Insert(); err != nil {
		//解锁
		//redisLock.Unlock(lockKey)
		return err
	}

	//更改用户余额
	//member.WithdrawBalance -= amount
	//member.TotalBalance -= amount
	err := member.Update("withdraw_balance")

	//解锁
	//redisLock.Unlock(lockKey)
	return err
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
