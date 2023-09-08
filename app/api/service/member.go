package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Member struct {
	request.MemberInfo
}

func (this Member) UpdateInfo(member *model.Member) error {
	//if this.Nickname != "" {
	//	member.Nickname = this.Nickname
	//}
	//if this.Email != "" {
	//	member.Email = this.Email
	//}
	//if this.Qq != "" {
	//	member.Qq = this.Qq
	//}
	//if this.Wechat != "" {
	//	member.Wechat = this.Wechat
	//}
	//if this.Avatar != "" {
	//member.Avatar = this.Avatar
	//}
	//return member.Update("nickname", "email", "qq", "wechat", "avatar")
	return nil
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
	if this.IdNumber == "" {
		return errors.New(lang.Lang("Id number cannot be empty"))
	}
	if this.Mobile == "" {
		return errors.New(lang.Lang("Phone number can not be blank"))
	}

	if !common.IsMobile(this.Mobile, global.Language) {
		return errors.New(lang.Lang("The phone format is incorrect"))
	}
	if !common.IsIdCard(this.IdNumber) {
		return errors.New(lang.Lang("The Id card format is incorrect"))
	}

	//会员实名认证后,无需再提交认证信息
	if member.IsReal == model.StatusAccept {
		return errors.New(lang.Lang("Real name authentication already exists"))
	}
	//分析会员是否提交实名认证信息
	info := model.MemberVerified{
		UId: member.Id,
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
		IdNumber: this.IdNumber,
	}
	//当身份证号已存在时
	if ex.Get() {
		if ex.Status == model.StatusAccept || ex.Status == model.StatusReview {
			return errors.New(lang.Lang("The Id number has been submitted for certification"))
		}
	}
	m := model.MemberVerified{
		UId:      member.Id,
		RealName: this.RealName,
		IdNumber: this.IdNumber,
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
	m := model.MemberParents{}
	where := "parent_id = ?"
	args := []interface{}{member.Id}
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	if this.Page == 0 {
		this.Page = 1
	}
	memberModel := model.Member{}
	//注册人数
	var total []int
	err := global.DB.Model(m).Select("uid").Where("parent_id = ?", member.Id).Find(&total).Error
	if err != nil {
		return nil, err
	}
	//激活人数
	var buyMember int64
	err = global.DB.Model(memberModel).Where("is_buy = ? and id in (?)", 1, total).Count(&buyMember).Error
	if err != nil {
		return nil, err
	}
	//充值金额
	var totalRecharge float64
	err = global.DB.Model(model.Recharge{}).Select("COALESCE(sum(amount),0)").Where("uid in (?) and status = 2", total).Scan(&totalRecharge).Error
	if err != nil {
		return nil, err
	}
	//总返佣
	//var totalRebate float64
	//err = global.DB.Model(memberModel).Select("COALESCE(sum(total_rebate),0)").Where("id in (?)", total).Scan(&totalRebate).Error
	//if err != nil {
	//	return nil, err
	//}
	config := model.SetBase{}
	config.Get()
	list, page := m.GetChildListByParentId(where, args, this.Page, this.PageSize)
	res.Page = FormatPage(page)
	res.List = make([]response.MyTeam, 0)
	for _, v := range list {
		ids := make([]int, 0)
		global.DB.Model(model.MemberParents{}).Select("uid").Where("parent_id = ?", v.Uid).Scan(&ids)

		var childRechargeMember int64
		//var childBuyMember int64
		var childBuyAmount float64
		if len(ids) > 0 {
			//下级用户总充值人数
			global.DB.Model(memberModel).Where("total_recharge > ? and id in (?)", 0, ids).Count(&childRechargeMember)
			//下级用户总激活人数
			//global.DB.Model(memberModel).Where("is_buy = ? and id in (?)", model.StatusOk, ids).Count(&childBuyMember)
		}
		//用户投资金额
		global.DB.Model(memberModel).Select("total_buy").Where("id = ?", v.Uid).Scan(&childBuyAmount)
		//用户返佣金额
		var childRebateAmount decimal.Decimal
		var payAmount float64
		global.DB.Model(model.OrderProduct{}).Select("pay_money").Where("uid = ?", v.Uid).Scan(&payAmount)
		switch v.Level {
		case 1:
			childRebateAmount = decimal.NewFromFloat(payAmount).Mul(config.OneSend).Div(decimal.NewFromInt(100)).Round(2)
		case 2:
			childRebateAmount = decimal.NewFromFloat(payAmount).Mul(config.TwoSend).Div(decimal.NewFromInt(100)).Round(2)
		case 3:
			childRebateAmount = decimal.NewFromFloat(payAmount).Mul(config.ThreeSend).Div(decimal.NewFromInt(100)).Round(2)
		}
		res.List = append(res.List, response.MyTeam{
			Id:             v.Member.Id,
			Username:       this.parseMobileNumber(v.Member.Username),
			RechargeMember: int(childRechargeMember),
			RealName:       v.Member.RealName,
			RegisterMember: len(ids),
			BuyAmount:      decimal.NewFromFloat(childBuyAmount),
			RebateAmount:   childRebateAmount,
			Level:          v.Level,
		})
	}
	res.RegisterMember = len(total)
	res.BuyMember = buyMember
	res.TotalRecharge = decimal.NewFromFloat(totalRecharge)
	res.TotalRebate = member.TotalRebate
	return &res, nil
}

func (this MemberTeam) parseMobileNumber(mobile string) string {
	numbers := []rune(mobile)
	length := len(numbers)
	return string(numbers[0:3]) + "****" + string(numbers[length-4:length])
}

type MemberTransfer struct {
	request.MemberTransfer
}

func (this *MemberTransfer) Transfer(member *model.Member) error {
	redis := model.Redis{}
	if err := redis.Lock("lock:" + member.Username); err != nil {
		return errors.New("操作频繁")
	}
	defer redis.Unlock("lock:" + member.Username)

	if this.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("金额必须大于0")
	}
	//if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
	//	return errors.New("交易密码错误")
	//}
	c := model.SetBase{}
	c.Get()
	switch this.Type {
	case 1: //可用U转R
		if member.UsdtBalance.LessThan(this.Amount) {
			return errors.New("usdt账户余额不足")
		}
		amount := this.Amount.Mul(c.UsdtSellRate)
		member.Balance = member.Balance.Add(amount)
		member.UsdtBalance = member.UsdtBalance.Sub(this.Amount)
		err := member.Update("balance", "usdt_balance")
		if err != nil {
			return err
		}
		//加入账单
		usdt := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     this.Amount,
			Before:     member.UsdtBalance.Add(this.Amount),
			After:      member.UsdtBalance,
			Desc:       "可用usdt转可用cny",
			IsFrontend: 1,
		}
		err = usdt.Insert()
		if err != nil {
			logrus.Errorf("可用usdt转可用cny 记录失败%v", err)
		}
		cny := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     amount,
			Before:     member.Balance.Sub(amount),
			After:      member.Balance,
			Desc:       "可用usdt转可用cny",
			IsFrontend: 1,
		}
		err = cny.Insert()
		if err != nil {
			logrus.Errorf("可用usdt转可用cny 记录失败%v", err)
		}

	case 2: //可用R转U
		amount := this.Amount.Div(c.UsdtBuyRate)
		if member.Balance.LessThan(this.Amount) {
			return errors.New("账户余额不足")
		}
		member.Balance = member.Balance.Sub(this.Amount)
		member.UsdtBalance = member.UsdtBalance.Add(amount)
		err := member.Update("balance", "usdt_balance")
		if err != nil {
			return err
		}
		//加入账单
		cny := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     this.Amount,
			Before:     member.Balance.Add(this.Amount),
			After:      member.Balance,
			Desc:       "可用cny转可用usdt",
			IsFrontend: 1,
		}
		err = cny.Insert()
		if err != nil {
			logrus.Errorf("可用cny转可用usdt 记录失败%v", err)
		}
		usdt := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     amount,
			Before:     member.UsdtBalance.Sub(amount),
			After:      member.UsdtBalance,
			Desc:       "可用cny转可用usdt",
			IsFrontend: 1,
		}
		err = usdt.Insert()
		if err != nil {
			logrus.Errorf("可用cny转可用usdt 记录失败%v", err)
		}

	case 3: //可提U转R
		if member.UsdtWithdrawBalance.LessThan(this.Amount) {
			return errors.New("usdt账户余额不足")
		}
		amount := this.Amount.Mul(c.UsdtSellRate)
		member.WithdrawBalance = member.WithdrawBalance.Add(amount)
		member.UsdtWithdrawBalance = member.UsdtWithdrawBalance.Sub(this.Amount)
		err := member.Update("withdraw_balance", "usdt_withdraw_balance")
		if err != nil {
			return err
		}
		//加入账单
		usdt := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     this.Amount,
			Before:     member.UsdtWithdrawBalance.Add(this.Amount),
			After:      member.UsdtWithdrawBalance,
			Desc:       "可提usdt转可提cny",
			IsFrontend: 1,
		}
		err = usdt.Insert()
		if err != nil {
			logrus.Errorf("可提usdt转可提cny 记录失败%v", err)
		}
		cny := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     amount,
			Before:     member.WithdrawBalance.Sub(amount),
			After:      member.WithdrawBalance,
			Desc:       "可提usdt转可提cny",
			IsFrontend: 1,
		}
		err = cny.Insert()
		if err != nil {
			logrus.Errorf("可提usdt转可提cny 记录失败%v", err)
		}

	case 4: //可提R转U
		amount := this.Amount.Div(c.UsdtBuyRate)
		if member.WithdrawBalance.LessThan(this.Amount) {
			return errors.New("账户余额不足")
		}
		member.WithdrawBalance = member.WithdrawBalance.Sub(this.Amount)
		member.UsdtWithdrawBalance = member.UsdtWithdrawBalance.Add(amount)
		err := member.Update("withdraw_balance", "usdt_withdraw_balance")
		if err != nil {
			return err
		}
		//加入账单
		cny := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     this.Amount,
			Before:     member.WithdrawBalance.Add(this.Amount),
			After:      member.WithdrawBalance,
			Desc:       "可提cny转可提usdt",
			IsFrontend: 1,
		}
		err = cny.Insert()
		if err != nil {
			logrus.Errorf("可提cny转可提usdt 记录失败%v", err)
		}
		usdt := model.Trade{
			UId:        member.Id,
			TradeType:  5,
			Amount:     amount,
			Before:     member.UsdtWithdrawBalance.Sub(amount),
			After:      member.UsdtWithdrawBalance,
			Desc:       "可提cny转可提usdt",
			IsFrontend: 1,
		}
		err = usdt.Insert()
		if err != nil {
			logrus.Errorf("可提cny转可提usdt 记录失败%v", err)
		}
	}
	return nil
}
