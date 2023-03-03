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
	"github.com/sirupsen/logrus"
	"strconv"
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
	where, args, _ := this.getWhere(member.ID)
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	res := make([]response.Withdraw, 0)
	for _, v := range list {
		//提现金额
		amount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.Amount)/model.UNITY), 64)
		freeAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.Fee)/model.UNITY), 64)
		totalAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(v.TotalAmount)/model.UNITY), 64)

		i := response.Withdraw{
			ID:          v.ID,
			OrderSn:     v.OrderSn,
			Type:        v.WithdrawType,
			TypeName:    v.WithdrawMethod.Name,
			BankName:    v.BankName,
			BranchBank:  v.BranchBank,
			RealName:    v.RealName,
			CardNumber:  v.CardNumber,
			BankPhone:   v.BankPhone,
			Amount:      amount,
			Fee:         freeAmount,
			TotalAmount: totalAmount,
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
	if this.TotalAmount <= 0 {
		return errors.New(lang.Lang("The withdrawal amount format is incorrect"))
	}
	//提现方式
	if this.Method == 0 {
		return errors.New(lang.Lang("Wrong withdrawal method"))
	}
	//银行卡ID
	if this.ID == 0 {
		return errors.New(lang.Lang("Bank card does not exist"))
	}
	memberBank := model.MemberBank{ID: this.ID}
	if !memberBank.Get() {
		return errors.New(lang.Lang("Bank card does not exist"))
	}
	amount := int64(this.TotalAmount * model.UNITY)
	//检查余额
	if member.UseBalance < amount {
		return errors.New(lang.Lang("Insufficient account balance"))
	}
	//验证交易密码
	if common.Md5String(this.WithdrawPassword+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New(lang.Lang("Incorrect withdraw password"))
	}
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
		if amount < c.WithdrawMinAmount {
			return errors.New(fmt.Sprintf(lang.Lang("Minimum withdraw %.2f"), float64(c.WithdrawMinAmount)/model.UNITY))
		}
		if amount > c.WithdrawMaxAmount {
			return errors.New(fmt.Sprintf(lang.Lang("Maximum withdraw %.2f"), float64(c.WithdrawMaxAmount)/model.UNITY))
		}
	}
	method := model.WithdrawMethod{ID: this.Method}
	if !method.Get() {
		return errors.New(lang.Lang("Wrong withdrawal method"))
	}

	//每日提现次数
	countModel := model.Withdraw{}
	countWhere := "uid = ? and create_time >= ? and `c_withdraw`.`status` != ?"
	countArgs := []interface{}{member.ID, common.GetTodayZero(), model.StatusRollback}
	count := countModel.Count(countWhere, countArgs)
	if count > 0 && count >= int64(c.WithdrawCount) {
		return errors.New(fmt.Sprintf(lang.Lang("You can only withdraw %d times per day"), c.WithdrawCount))
	}
	//计算手续费
	fee := int64(c.WithdrawFee) * amount / int64(model.UNITY)
	m := model.Withdraw{
		UID:          member.ID,
		WithdrawType: this.Method,
		BankName:     memberBank.BankName,
		BranchBank:   memberBank.BranchBank,
		RealName:     memberBank.RealName,
		CardNumber:   memberBank.CardNumber,
		BankPhone:    memberBank.BankPhone,
		Amount:       amount - fee,
		Fee:          fee,
		TotalAmount:  amount,
		Status:       model.StatusReview,
		OrderSn:      common.OrderSn(),
	}
	//生成账单
	trade := model.Trade{
		UID:       m.UID,
		TradeType: 4,
		ItemID:    m.ID,
		Amount:    amount,
		Before:    member.UseBalance,
		After:     member.UseBalance - amount,
		Desc:      "提现申请",
	}
	if err := trade.Insert(); err != nil {
		return err
	}

	member.UseBalance -= amount
	member.TotalBalance -= amount
	member.Update("use_balance", "total_balance")
	return m.Insert()
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
			ID:   v.ID,
			Name: v.Name,
			Code: v.Code,
			Icon: v.Icon,
			Fee:  float64(v.Fee) / model.UNITY,
		}
		res = append(res, i)
	}
	return res
}
