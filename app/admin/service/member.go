package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	//"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type MemberList struct {
	request.MemberList
}

func (this MemberList) PageList() (response.MemberListData, error) {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Member{}
	where, args := this.getWhere()
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	res := make([]response.MemberInfo, 0)
	for _, v := range list {
		p := model.MemberParents{Uid: v.Id, Level: 1}
		p.Get2()
		invite := model.InviteCode{UId: v.Id}
		invite.Get()
		agent := model.Agent{Id: v.AgentId}
		agent.Get()
		i := response.MemberInfo{
			Id:                v.Id,
			Username:          v.Username,
			Balance:           v.Balance,
			WithdrawBalance:   v.WithdrawBalance,
			IsReal:            v.IsReal,
			RealName:          v.RealName,
			InvestFreeze:      v.InvestFreeze,
			InvestAmount:      v.InvestAmount,
			InvestIncome:      v.InvestIncome,
			Avatar:            v.Avatar,
			Status:            v.Status,
			FundsStatus:       v.FundsStatus,
			Level:             v.Level,
			Score:             v.Score,
			LastLoginTime:     v.LastLoginTime,
			LastLoginIP:       v.LastLoginIp,
			RegTime:           v.RegTime,
			RegisterIP:        v.RegisterIp,
			DisableLoginTime:  v.DisableLoginTime,
			DisableBetTime:    v.DisableBetTime,
			IsBuy:             v.IsBuy,
			Code:              invite.Code,
			TopId:             p.ParentId,
			TopName:           p.Parent.Username,
			AgentName:         agent.Account,
			WithdrawThreshold: v.WithdrawThreshold,
		}
		res = append(res, i)
	}
	return response.MemberListData{
		List: res,
		Page: FormatPage(page),
	}, nil

}
func (this MemberList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Id > 0 {
		where["id"] = this.Id
	}
	if this.Mobile != "" {
		where["mobile"] = this.Mobile
	}
	if this.Username != "" {
		where["username"] = this.Username
	}
	if this.RealName != "" {
		where["real_name"] = this.RealName
	}
	if this.StartTime != "" {
		where["reg_time >"] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where["reg_time <"] = common.DateToUnix(this.EndTime)
	}
	if this.AgentName != "" {
		agent := model.Agent{Account: this.AgentName}
		if agent.Get() {
			where["agent_id"] = agent.Id
		} else {
			where["agent_id"] = -1
		}
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type MemberCreate struct {
	request.MemberCreate
}

func (this MemberCreate) Create(c *gin.Context) error {
	if this.Username == "" {
		return errors.New("用户名不能为空")
	}
	if this.Password == "" {
		return errors.New("密码不能为空")
	}
	m := model.Member{
		Username: this.Username,
	}
	if m.Get() {
		return errors.New("用户名已存在")
	}
	m.Salt = common.RandStringRunes(6)
	m.Password = common.Md5String(this.Password + m.Salt)
	m.RegTime = time.Now().Unix()
	m.RegisterIp = c.ClientIP()
	//m.Balance = int64(this.Balance)
	return m.Insert()
}

type MemberUpdate struct {
	request.MemberUpdate
}

func (this MemberUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Member{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("用户不存在")
	}
	m.Description = this.Description
	return m.Update("agent_id", "score", "level", "description", "withdraw_amount")
}

type MemberUpdatePassword struct {
	request.MemberUpdatePassword
}

func (this MemberUpdatePassword) UpdatePassword() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Member{
		Id: this.Id,
	}
	if this.Password == "" && this.PayPassword == "" {
		return nil
	}
	if !m.Get() {
		return errors.New("用户不存在")
	}
	if this.Password != "" {
		m.Salt = common.RandStringRunes(6)
		m.Password = common.Md5String(this.Password + m.Salt)
	}
	if this.PayPassword != "" {
		m.WithdrawSalt = common.RandStringRunes(6)
		m.WithdrawPassword = common.Md5String(this.PayPassword + m.WithdrawSalt)
	}
	return m.Update("salt", "password", "withdraw_salt", "withdraw_password")
}

type MemberUpdateStatus struct {
	request.MemberUpdateStatus
}

func (this MemberUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Status != model.StatusOk && this.Status != model.StatusClose {
		return errors.New("状态错误")
	}
	member := model.Member{Id: this.Id}
	if !member.Get() {
		return errors.New("用户不存在")
	}
	switch this.Type {
	case "login":
		member.Status = this.Status
		if this.Status != model.StatusOk {
			member.Token = ""
		}
	case "funds":
		member.FundsStatus = this.Status
	default:
		return errors.New("类型错误")
	}
	return member.Update("status", "funds_status", "token")
}

type MemberRemove struct {
	request.MemberUpdateStatus
}

func (this MemberRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	member := model.Member{Id: this.Id}
	if !member.Get() {
		return errors.New("用户不存在")
	}
	return member.Remove()
}

type MemberVerifiedList struct {
	request.MemberVerifiedList
}

func (this MemberVerifiedList) PageList() response.MemberVerifiedData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.MemberVerified{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.MemberVerified, 0)
	for _, v := range list {
		i := response.MemberVerified{
			Id:         v.Id,
			UId:        v.UId,
			Username:   v.Member.Username,
			RealName:   v.RealName,
			IdNumber:   v.IdNumber,
			Mobile:     v.Mobile,
			Frontend:   v.Frontend,
			Backend:    v.Backend,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.MemberVerifiedData{
		List: res,
		Page: FormatPage(page),
	}
}
func (this MemberVerifiedList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.Status > 0 {
		where[model.MemberVerified{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type MemberVerifiedUpdate struct {
	request.MemberVerifiedUpdate
}

func (this MemberVerifiedUpdate) Update() error {
	if len(this.Ids) == 0 {
		return errors.New("参数错误")
	}
	for _, v := range this.Ids {
		m := model.MemberVerified{Id: v}
		if !m.Get() {
			return errors.New("记录不存在")
		}
		if m.Status == model.StatusAccept {
			return errors.New("已通过")
		}
		member := model.Member{Id: m.UId}
		if !member.Get() {
			return errors.New("用户不存在")
		}
		//获取基础配置表信息
		c := model.SetBase{}
		c.Get()
		switch this.Status {
		case model.StatusAccept:
			member.RealName = m.RealName
			if decimal.Zero.LessThan(c.VerifiedSend) {
				//加入账变记录
				trade := model.Trade{
					UId:        member.Id,
					TradeType:  8,
					Amount:     c.VerifiedSend,
					Before:     member.Balance,
					After:      member.Balance.Add(c.VerifiedSend),
					Desc:       "实名认证礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				trade.Insert()
				//第一次实名通过的时候送奖金
				member.Balance = member.Balance.Add(c.VerifiedSend)
			}
		case model.StatusRollback:
			member.RealName = ""
		default:
			return errors.New("状态错误")
		}
		m.Status = this.Status
		m.Update("status")
		member.IsReal = this.Status
		err := member.Update("is_real", "real_name", "balance")
		if err != nil {
			return err
		}
	}
	return nil
}

type MemberVerifiedInfoUpdate struct {
	request.MemberVerifiedInfoUpdate
}

func (this MemberVerifiedInfoUpdate) UpdateInfo() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.MemberVerified{Id: this.Id}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	member := model.Member{Id: m.UId}
	if !member.Get() {
		return errors.New("用户不存在")
	}
	m.RealName = this.RealName
	m.IdNumber = this.IdNumber
	if m.Status == model.StatusAccept {
		member.RealName = this.RealName
		member.Update("real_name")
	}
	return m.Update("real_name", "id_number")
}

type MemberVerifiedRemove struct {
	request.MemberVerifiedRemove
}

func (this MemberVerifiedRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.MemberVerified{Id: this.Id}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	member := model.Member{Id: m.UId}
	if !member.Get() {
		return errors.New("用户不存在")
	}
	member.IsReal = 0
	return member.Update("is_real")
}

type MemberTeam struct {
	request.MemberTeamReq
}

func (this *MemberTeam) GetTeam() response.MemberListData {
	var res response.MemberListData
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	if this.Page == 0 {
		this.Page = 1
	}
	m := model.MemberParents{}
	var where string
	var args []interface{}
	if this.Level > 0 {
		where = "parent_id = ? and c_member_parents.level = ?"
		args = []interface{}{this.UserId, this.Level}
	} else {
		where = "parent_id = ?"
		args = []interface{}{this.UserId}
	}

	list, page := m.GetByPuid(where, args, this.Page, this.PageSize)
	if len(list) == 0 {
		return res
	}
	//总投资，团队总可用 可提，团队总收益
	child := model.MemberParents{}
	where1 := "parent_id = ?"
	args2 := []interface{}{this.UserId}
	users, _ := child.GetByPuidAll(where1, args2)
	//总投资
	var childIds []int
	for i := range users {
		childIds = append(childIds, users[i].Member.Id)
	}
	where3 := "uid in (?) "
	args3 := []interface{}{childIds}
	product := model.OrderProduct{}
	totalSumProduct := product.Sum(where3, args3, "pay_money")
	m1 := model.Member{}
	where4 := "id in (?)"
	args4 := []interface{}{childIds}
	//总可用
	totalSumBalance := m1.Sum(where4, args4, "balance")
	//总可提
	totalSumUseBalance := m1.Sum(where4, args4, "withdraw_balance")
	//总总收益
	totalSumIncome := m1.Sum(where4, args4, "income")
	//总充值
	rechargeModel := model.Recharge{}
	where5 := "uid in (?) and status = 2"
	args5 := []interface{}{childIds}
	totalRechargeAmount := rechargeModel.Sum(where5, args5, "amount")
	//总提现
	withdrawModel := model.Withdraw{}
	where6 := "uid in (?) and status = 2"
	args6 := []interface{}{childIds}
	totalWithdrawAmount := withdrawModel.Sum(where6, args6, "total_amount")
	todayZeroTime := common.GetTodayZero()
	//今日总充值
	rechargeModel2 := model.Recharge{}
	where7 := "update_time >= ? and uid in (?) and status = 2"
	args7 := []interface{}{todayZeroTime, childIds}
	todayRechargeAmount := rechargeModel2.Sum(where7, args7, "amount")
	//今日总提现
	withdrawModel2 := model.Withdraw{}
	where8 := "update_time >= ? and uid in (?) and status = 2"
	args8 := []interface{}{todayZeroTime, childIds}
	todayWithdrawAmount := withdrawModel2.Sum(where8, args8, "total_amount")
	//充值总人数
	rechargeModel3 := model.Recharge{}
	where9 := "uid in (?) and status = 2"
	args9 := []interface{}{childIds}
	totalRechargeCount := rechargeModel3.GetMemberCount(where9, args9)

	//今日充值人数
	rechargeModel4 := model.Recharge{}
	where10 := "update_time >= ? and uid in (?) and status = 2"
	args10 := []interface{}{todayZeroTime, childIds}
	todayRechargeCount := rechargeModel4.GetMemberCount(where10, args10)
	//当月充值金额
	year, month, _ := time.Now().Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
	rechargeModel5 := model.Recharge{}
	where11 := "update_time >= ? and uid in (?) and status = 2"
	args11 := []interface{}{firstOfMonth.Unix(), childIds}
	monthRechargeAmount := rechargeModel5.Sum(where11, args11, "amount")

	res.TotalSumProduct = decimal.NewFromFloat(totalSumProduct)
	res.TotalSumBalance = decimal.NewFromFloat(totalSumBalance)
	res.TotalSumUseBalance = decimal.NewFromFloat(totalSumUseBalance)
	res.TotalSumIncome = decimal.NewFromFloat(totalSumIncome)
	res.TotalMemberCount = len(childIds)
	res.TotalRechargeAmount = decimal.NewFromFloat(totalRechargeAmount)
	res.TotalWithdrawAmount = decimal.NewFromFloat(totalWithdrawAmount)
	res.TodayRechargeAmount = decimal.NewFromFloat(todayRechargeAmount)
	res.TodayWithdrawAmount = decimal.NewFromFloat(todayWithdrawAmount)
	res.TotalRechargeCount = totalRechargeCount
	res.TodayRechargeCount = todayRechargeCount
	res.MonthRechargeAmount = decimal.NewFromFloat(monthRechargeAmount)
	res.Page = FormatPage(page)
	items := make([]response.MemberInfo, 0)
	for _, v := range list {
		p := model.MemberParents{Uid: v.Member.Id, Level: 1}
		p.Get()
		items = append(items, response.MemberInfo{
			Id:               v.Member.Id,
			Username:         v.Member.Username,
			Balance:          v.Member.Balance,
			WithdrawBalance:  v.Member.WithdrawBalance,
			IsReal:           v.Member.IsReal,
			RealName:         v.Member.RealName,
			InvestFreeze:     v.Member.InvestFreeze,
			InvestAmount:     v.Member.InvestAmount,
			InvestIncome:     v.Member.InvestIncome,
			Avatar:           v.Member.Avatar,
			Status:           v.Member.Status,
			FundsStatus:      v.Member.FundsStatus,
			Level:            v.Level,
			Score:            v.Member.Score,
			LastLoginTime:    v.Member.LastLoginTime,
			LastLoginIP:      v.Member.LastLoginIp,
			RegTime:          v.Member.RegTime,
			RegisterIP:       v.Member.RegisterIp,
			DisableLoginTime: v.Member.DisableLoginTime,
			DisableBetTime:   v.Member.DisableBetTime,
			//Code:             v.Member.Code,
			IsBuy:              v.Member.IsBuy,
			TopId:              p.Parent.Id,
			TopName:            p.Parent.Username,
			ProductOrderAmount: v.Member.TotalBuy,
			WithdrawThreshold:  v.Member.WithdrawThreshold,
		})
	}
	res.List = items
	return res
}

type SendCoupon struct {
	request.SendCouponReq
}

func (this SendCoupon) Send() error {
	s := strings.Split(this.Ids, ",")
	if len(s) == 0 {
		return errors.New("用户Id不能为空")
	}
	if this.CouponId == 0 {
		return errors.New("优惠券Id不能为空")
	}
	c := model.Coupon{Id: this.CouponId}
	if !c.Get() {
		return errors.New("优惠券Id不存在,请先创建")
	}

	for _, v := range s {
		id, _ := strconv.Atoi(v)
		memberCoupon := model.MemberCoupon{
			Uid:      id,
			CouponId: this.CouponId,
			IsUse:    1,
		}
		memberCoupon.Insert()
	}

	return nil
}

type GetCode struct {
	request.GetCodeReq
}

func (this *GetCode) GetCode() string {
	return global.REDIS.Get(fmt.Sprintf("reg_%v", this.Mobile)).Val()
}
