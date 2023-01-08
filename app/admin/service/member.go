package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/global"
	"finance/model"
	"fmt"
	"github.com/gin-gonic/gin"
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
		p := model.MemberRelation{UID: v.ID, Level: 1}
		p.Get2()
		i := response.MemberInfo{
			ID:               v.ID,
			Username:         v.Username,
			Balance:          float64(v.Balance) / model.UNITY,
			UseBalance:       float64(v.UseBalance) / model.UNITY,
			TotalBalance:     float64(v.TotalBalance) / model.UNITY,
			IsReal:           v.IsReal,
			RealName:         v.RealName,
			InvestFreeze:     float64(v.InvestFreeze) / model.UNITY,
			InvestAmount:     float64(v.InvestAmount) / model.UNITY,
			InvestIncome:     float64(v.InvestIncome) / model.UNITY,
			Avatar:           v.Avatar,
			Status:           v.Status,
			FundsStatus:      v.FundsStatus,
			Level:            v.Level,
			Score:            v.Score,
			LastLoginTime:    v.LastLoginTime,
			LastLoginIP:      v.LastLoginIP,
			RegTime:          v.RegTime,
			RegisterIP:       v.RegisterIP,
			Nickname:         v.Nickname,
			Mobile:           v.Mobile,
			Email:            v.Email,
			Qq:               v.Qq,
			Wechat:           v.Wechat,
			DisableLoginTime: v.DisableLoginTime,
			DisableBetTime:   v.DisableBetTime,
			IsBuy:            v.IsBuy,
			Code:             v.Code,
			TopId:            p.Puid,
			TopName:          p.Member2.Username,
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
	if this.ID > 0 {
		where["id"] = this.ID
	}
	if this.Mobile != "" {
		where["mobile"] = this.Mobile
	}
	if this.Username != "" {
		where["username"] = this.Username
	}
	if this.StartTime != "" {
		where["reg_time >"] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where["reg_time <"] = common.DateToUnix(this.EndTime)
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
	m.RegisterIP = c.ClientIP()
	m.Balance = int64(this.Balance * model.UNITY)
	return m.Insert()
}

type MemberUpdate struct {
	request.MemberUpdate
}

func (this MemberUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Member{
		ID: this.ID,
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
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Member{
		ID: this.ID,
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
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Status != model.StatusOk && this.Status != model.StatusClose {
		return errors.New("状态错误")
	}
	member := model.Member{ID: this.ID}
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
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	member := model.Member{ID: this.ID}
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
			ID:         v.ID,
			UID:        v.UID,
			Username:   v.Member.Username,
			RealName:   v.RealName,
			IDNumber:   v.IDNumber,
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
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.MemberVerified{ID: this.ID}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	if this.Status != model.StatusAccept && this.Status != model.StatusRollback {
		return errors.New("状态错误")
	}
	m.Status = this.Status
	m.Update("status")
	member := model.Member{ID: m.UID}
	if !member.Get() {
		return errors.New("用户不存在")
	}
	c := model.SetBase{}
	c.Get()
	if member.IsOneShiming == 1 && this.Status == 2 {
		//加入账变记录
		trade := model.Trade{
			UID:        member.ID,
			TradeType:  8,
			Amount:     int64(c.VerifiedSend),
			Before:     member.UseBalance,
			After:      member.UseBalance + int64(c.VerifiedSend),
			Desc:       "实名认证礼金",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		trade.Insert()

		//第一次实名通过的时候送奖金
		member.IsOneShiming = 2
		member.UseBalance += int64(c.VerifiedSend)
		member.TotalBalance += int64(c.VerifiedSend)
		member.Income += int64(c.VerifiedSend)
	}

	member.IsReal = this.Status
	member.RealName = m.RealName
	member.Mobile = m.Mobile
	return member.Update("is_real", "real_name", "mobile", "income", "use_balance", "is_one_shiming", "total_balance")
}

type MemberVerifiedRemove struct {
	request.MemberVerifiedRemove
}

func (this MemberVerifiedRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.MemberVerified{ID: this.ID}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	member := model.Member{ID: m.UID}
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
	m := model.MemberRelation{}
	var where string
	var args []interface{}
	if this.Level != nil {
		where = "puid = ? and level = ?"
		args = []interface{}{this.UserId, this.Level}
	} else {
		where = "puid = ?"
		args = []interface{}{this.UserId}
	}

	list, page := m.GetByPuid(where, args, this.Page, this.PageSize)
	if len(list) == 0 {
		return res
	}
	//总投资，团队总可用 可提，团队总收益
	child := model.MemberRelation{}
	where1 := "puid = ?"
	args2 := []interface{}{this.UserId}
	users, _ := child.GetByPuidAll(where1, args2)
	//总投资
	var childIds []int
	for i := range users {
		childIds = append(childIds, users[i].Member.ID)
	}

	where3 := "uid in (?) "
	args3 := []interface{}{childIds}
	product := model.OrderProduct{}
	totalSumProduct := product.Sum(where3, args3, "pay_money")

	m1 := model.Member{}
	where4 := "id in (?) "
	args4 := []interface{}{childIds}
	//总可用
	totalSumBalance := m1.Sum(where4, args4, "balance")
	//总可提
	totalSumUseBalance := m1.Sum(where4, args4, "use_balance")
	//总总收益
	totalSumIncome := m1.Sum(where4, args4, "income")
	res.TotalSumProduct = float64(totalSumProduct) / model.UNITY
	res.TotalSumBalance = float64(totalSumBalance) / model.UNITY
	res.TotalSumUseBalance = float64(totalSumUseBalance) / model.UNITY
	res.TotalSumIncome = float64(totalSumIncome) / model.UNITY
	res.TotalMemberCount = len(childIds)
	res.Page = FormatPage(page)
	items := make([]response.MemberInfo, 0)
	for i := range list {
		p := model.MemberRelation{UID: list[i].Member.ID, Level: 1}
		p.Get2()
		//获取用户投资金额
		orderModel := model.OrderProduct{}
		payMondyAmount := orderModel.Sum("id = ?", []interface{}{list[i].Member.ID}, "pay_money")

		items = append(items, response.MemberInfo{
			ID:                 list[i].Member.ID,
			Username:           list[i].Member.Username,
			TotalBalance:       float64(list[i].Member.TotalBalance) / model.UNITY,
			Balance:            float64(list[i].Member.Balance) / model.UNITY,
			UseBalance:         float64(list[i].Member.UseBalance) / model.UNITY,
			IsReal:             list[i].Member.IsReal,
			RealName:           list[i].Member.RealName,
			InvestFreeze:       float64(list[i].Member.InvestFreeze) / model.UNITY,
			InvestAmount:       float64(list[i].Member.InvestAmount) / model.UNITY,
			InvestIncome:       float64(list[i].Member.InvestIncome) / model.UNITY,
			Avatar:             list[i].Member.Avatar,
			Status:             list[i].Member.Status,
			FundsStatus:        list[i].Member.FundsStatus,
			Level:              int(list[i].Level),
			Score:              list[i].Member.Score,
			LastLoginTime:      list[i].Member.LastLoginTime,
			LastLoginIP:        list[i].Member.LastLoginIP,
			RegTime:            list[i].Member.RegTime,
			RegisterIP:         list[i].Member.RegisterIP,
			Nickname:           list[i].Member.Nickname,
			Mobile:             list[i].Member.Mobile,
			Email:              list[i].Member.Email,
			Qq:                 list[i].Member.Qq,
			Wechat:             list[i].Member.Wechat,
			DisableLoginTime:   list[i].Member.DisableLoginTime,
			DisableBetTime:     list[i].Member.DisableBetTime,
			Code:               list[i].Member.Code,
			IsBuy:              list[i].Member.IsBuy,
			TopId:              p.Puid,
			TopName:            p.Member2.Username,
			ProductOrderAmount: float64(payMondyAmount) / model.UNITY,
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
		return errors.New("用户ID不能为空")
	}
	if this.CouponId == 0 {
		return errors.New("优惠券ID不能为空")
	}
	c := model.Coupon{ID: this.CouponId}
	if !c.Get() {
		return errors.New("优惠券ID不存在,请先创建")
	}

	for _, v := range s {
		id, _ := strconv.Atoi(v)
		memberCoupon := model.MemberCoupon{
			Uid:      int64(id),
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
