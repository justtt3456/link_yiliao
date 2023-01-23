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
	"time"
)

type Member struct {
	request.MemberInfo
}

func (this Member) UpdateInfo(member *model.Member) error {
	if this.Nickname != "" {
		member.Nickname = this.Nickname
	}
	if this.Email != "" {
		member.Email = this.Email
	}
	if this.Qq != "" {
		member.Qq = this.Qq
	}
	if this.Wechat != "" {
		member.Wechat = this.Wechat
	}
	if this.Avatar != "" {
		member.Avatar = this.Avatar
	}
	return member.Update("nickname", "email", "qq", "wechat", "avatar")
}
func (this Member) Logout(member *model.Member) error {
	member.Token = ""
	return member.Update("token")
}

type MemberPassword struct {
	request.MemberPassword
}

func (this MemberPassword) UpdatePassword(member *model.Member) error {
	if this.Password == "" {
		return errors.New(lang.Lang("The original password cannot be empty"))
	}
	if this.NewPassword == "" {
		return errors.New(lang.Lang("Password cannot be empty"))
	}
	if this.PasswordConfirm != this.NewPassword {
		return errors.New(lang.Lang("The two passwords are inconsistent"))
	}
	if common.Md5String(this.Password+member.Salt) != member.Password {
		return errors.New(lang.Lang("The original password is wrong"))
	}
	member.Salt = common.RandStringRunes(6)
	member.Password = common.Md5String(this.NewPassword + member.Salt)
	return member.Update("password", "salt")
}

func (this MemberPassword) UpdatePayPassword(member *model.Member) error {

	if this.NewPassword == "" {
		return errors.New(lang.Lang("Password cannot be empty"))
	}
	if this.PasswordConfirm != this.NewPassword {
		return errors.New(lang.Lang("The two passwords are inconsistent"))
	}
	if member.WithdrawPassword != "" {
		if this.Password == "" {
			return errors.New(lang.Lang("The original password cannot be empty"))
		}
		if common.Md5String(this.Password+member.WithdrawSalt) != member.WithdrawPassword {
			return errors.New(lang.Lang("The original password is wrong"))
		}
	}
	member.WithdrawSalt = common.RandStringRunes(6)
	member.WithdrawPassword = common.Md5String(this.NewPassword + member.WithdrawSalt)
	return member.Update("withdraw_password", "withdraw_salt")
}

type MemberVerified struct {
	request.MemberVerified
}

func (this MemberVerified) Verified(member model.Member) error {
	if this.RealName == "" {
		return errors.New(lang.Lang("Real name cannot be empty"))
	}
	if this.IDNumber == "" {
		return errors.New(lang.Lang("ID number cannot be empty"))
	}
	if this.Mobile == "" {
		return errors.New(lang.Lang("Phone number can not be blank"))
	}

	if !common.IsMobile(this.Mobile, global.Language) {
		return errors.New(lang.Lang("The phone format is incorrect"))
	}
	if !common.IsIdCard(this.IDNumber) {
		return errors.New(lang.Lang("The ID card format is incorrect"))
	}

	//会员实名认证后,无需再提交认证信息
	if member.IsReal == model.StatusAccept {
		return errors.New(lang.Lang("Real name authentication already exists"))
	}
	//分析会员是否提交实名认证信息
	info := model.MemberVerified{
		UID: member.ID,
	}
	//当会员已提交认证信息时
	if info.Get() {
		//当实名状态为待审核或已审核时,直接返回信息提示
		if info.Status == model.StatusAccept || info.Status == model.StatusReview {
			return errors.New(lang.Lang("You have submitted real name authentication"))
		}
		//当实名状态为已驳回时,将旧数据删除
		if info.Status == model.StatusRollback {
			//删除驳回认证
			info.Remove()
		}
	}

	//同一个身份证号只可以认证一次
	ex := model.MemberVerified{
		IDNumber: this.IDNumber,
	}
	//当身份证号已存在时
	if ex.Get() {
		return errors.New(lang.Lang("The ID number has been submitted for certification"))
	}

	m := model.MemberVerified{
		UID:      member.ID,
		RealName: this.RealName,
		IDNumber: this.IDNumber,
		Mobile:   this.Mobile,
		Frontend: this.Frontend,
		Backend:  this.Backend,
		Status:   model.StatusReview,
	}
	err := m.Insert()
	if err != nil {
		return err
	}
	//用户实名审核中
	member.IsReal = model.StatusReview
	return member.Update("is_real")
}

type MemberTeam struct {
	request.Pagination
}

func (this MemberTeam) GetTeam(member model.Member) (*response.MyTeamList, error) {
	var res response.MyTeamList
	m := model.MemberRelation{}
	where := "puid = ?"
	args := []interface{}{member.ID}
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	if this.Page == 0 {
		this.Page = 1
	}
	list, page := m.GetChildListByParentId(where, args, this.Page, this.PageSize)
	if len(list) == 0 {
		return nil, errors.New("无数据")
	}
	where2 := "trade_type in (?) and uid = ?"
	args2 := []interface{}{[]int{18, 19, 20}, member.ID}
	trade := model.Trade{}
	income2 := trade.Sum(where2, args2, "amount")
	res.TotalIncome = float64(income2) / model.UNITY
	res.Page = FormatPage(page)
	for i := range list {
		where1 := "trade_type in (?) and uid = ?"
		args1 := []interface{}{[]int{18, 19, 20}, member.ID}
		if list[i].Member.ID != member.ID {
			where1 += " and item_id = ?"
			args1 = append(args1, list[i].Member.ID)
		}
		trade := model.Trade{}
		income := trade.Sum(where1, args1, "amount")

		res.List = append(res.List, response.MyTeam{
			ID:       list[i].Member.ID,
			Username: list[i].Member.Username,
			Level:    int(list[i].Level),
			RegTime:  list[i].Member.RegTime,
			Income:   float64(income) / model.UNITY,
			RealName: list[i].MemberVerified.RealName,
		})
	}
	return &res, nil
}

type MemberTransfer struct {
	request.MemberTransfer
}

func (this *MemberTransfer) Transfer(member *model.Member) error {
	//1=可用转可提  2=可提转可用
	c := model.SetFunds{}
	c.Get()
	if this.Amount <= 0 {
		return errors.New("金额必须大于0")
	}
	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
	}
	amount := int64(this.Amount * model.UNITY)
	switch this.Type {
	//case 1:
	//	trade := model.Trade{UID: member.ID, TradeType: 5}
	//	count, err := trade.CountByToday()
	//	if err != nil {
	//		return err
	//	}
	//	if count >= c.DayTurnMoneyNum {
	//		return errors.New(fmt.Sprintf("每日只能转%v次", c.DayTurnMoneyNum))
	//	}
	//	member.Balance -= amount
	//	member.UseBalance += amount
	//	err = member.Update("balance", "use_balance")
	//	if err != nil {
	//		return err
	//	}
	//	//加入账单
	//	inc := model.Trade{
	//		UID:        member.ID,
	//		TradeType:  5,
	//		Amount:     amount,
	//		Before:     member.UseBalance - amount,
	//		After:      member.UseBalance,
	//		Desc:       "可用转可提",
	//		CreateTime: time.Now().Unix(),
	//		UpdateTime: time.Now().Unix(),
	//		IsFrontend: 1,
	//	}
	//	err = inc.Insert()
	//	if err != nil {
	//		logrus.Errorf("可用转可提 记录失败%v", err)
	//	}
	case 2:
		//当前余额分析
		if amount > member.UseBalance {
			return errors.New("可用提现金额不足!")
		}
		trade := model.Trade{UID: member.ID, TradeType: 6}
		count, err := trade.CountByToday()
		if err != nil {
			return err
		}
		if count >= c.DayTurnMoneyNum {
			return errors.New(fmt.Sprintf("每日只能转账%v次", c.DayTurnMoneyNum))
		}

		member.UseBalance -= amount
		member.Balance += amount
		err = member.Update("balance", "use_balance")
		if err != nil {
			return err
		}
		//加入账单
		inc := model.Trade{
			UID:        member.ID,
			TradeType:  6,
			Amount:     amount,
			Before:     member.Balance - amount,
			After:      member.Balance,
			Desc:       "可提转可用",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err = inc.Insert()
		if err != nil {
			logrus.Errorf("可提转可用 记录失败%v", err)
		}
	}
	return nil
}
