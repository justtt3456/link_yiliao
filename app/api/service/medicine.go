package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/logic"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MedicineList struct {
	request.MedicineList
}

func (this MedicineList) PageList() response.MedicineListData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Medicine{}
	where, args, _ := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Medicine, 0)
	for _, v := range list {
		i := response.Medicine{
			Id:                v.Id,
			Name:              v.Name,
			Price:             v.Price,
			Img:               v.Img,
			Desc:              v.Desc,
			WithdrawThreshold: v.WithdrawThreshold,
			Interval:          v.Interval,
		}
		res = append(res, i)
	}
	return response.MedicineListData{List: res, Page: FormatPage(page)}
}

func (this MedicineList) getWhere() (string, []interface{}, error) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where[model.Medicine{}.TableName()+".name"] = this.Name
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals, nil
}

type MedicineBuy struct {
	request.MedicineBuyReq
}

func (this *MedicineBuy) Buy(member *model.Member) error {
	//实名认证
	if member.IsReal != 2 {
		return errors.New("请实名认证！")
	}
	//验证收货地址
	addressId := this.AddressId
	if addressId <= 0 {
		return errors.New("请选择收货地址")
	}
	address := model.MemberAddress{Id: addressId, UId: member.Id}
	if !address.Get() {
		return errors.New("收货地址不存在")
	}
	//产品Id
	if this.Id <= 0 {
		return errors.New("药品Id格式不正确！")
	}
	//添加Redis乐观锁
	lockKey := fmt.Sprintf("medicine_buy:%v:%v", member.Id, this.Id)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	if !redisLock.Lock(lockKey) {
		return errors.New(lang.Lang("During data processing, Please try again later"))
	}
	defer redisLock.Unlock(lockKey)
	if this.Quantity <= 0 {
		return errors.New("数量错误！")
	}
	product := model.Medicine{Id: this.Id}
	if !product.Get() {
		return errors.New("药品不存在！")
	}
	if product.Status != model.StatusOk {
		return errors.New("项目已下架！")
	}
	amount := product.Price.Mul(decimal.NewFromInt(int64(this.Quantity)))
	//余额检查
	if member.Balance.LessThan(amount) {
		return errors.New("余额不足,请先充值！")
	}
	//交易密码验证
	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
	}
	buyLogic := logic.NewMedicineBuyLogic()
	err := buyLogic.MedicineBuy(member, product, amount, addressId, this.Quantity)
	if err != nil {
		return err
	}
	return nil
}

type BuyMedicineList struct {
	request.MedicineOrder
}

func (this BuyMedicineList) List(member *model.Member) *response.MedicineBuyListResp {
	if this.Page == 0 {
		this.Page = 1
	}
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	m := model.MedicineOrder{}
	res := response.MedicineBuyListResp{}
	list, page := m.PageList(m.TableName()+".uid = ?", []interface{}{member.Id}, this.Page, this.PageSize)
	if len(list) == 0 {
		return nil
	}
	items := make([]response.MedicineBuyList, 0)
	for _, v := range list {

		product := model.Medicine{Id: m.Pid}
		if !product.Get() {
			continue
		}
		items = append(items, response.MedicineBuyList{
			Name:       v.Medicine.Name,
			CreateTime: v.CreateTime,
			Amount:     v.PayMoney,
			Address:    v.Address.Name + v.Address.Phone + v.Address.Address + v.Address.Other,
		})
	}
	res.List = items
	res.Page = FormatPage(page)
	return &res
}
