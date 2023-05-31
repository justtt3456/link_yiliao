package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

type CouponList struct {
	request.Request
}

func (this *CouponList) CouponList() response.CouponResp {
	res := make([]response.Coupon, 0)
	m := model.Coupon{}
	s := m.List()
	for _, v := range s {
		res = append(res, response.Coupon{
			Id:    v.Id,
			Price: v.Price,
		})
	}
	return response.CouponResp{List: res}
}

type CouponAdd struct {
	request.AddCoupon
}

func (this *CouponAdd) Add() error {
	m := model.Coupon{}
	if this.Price.LessThanOrEqual(decimal.Zero) {
		return errors.New("金额不能为0")
	}
	m.Price = this.Price
	return m.Insert()
}

type ActiveList struct {
	request.Request
}

func (this *ActiveList) PageList() response.ActiveResp {
	res := make([]response.Active, 0)
	m := model.CouponActivity{}
	s := m.List()
	for _, v := range s {
		res = append(res, response.Active{
			Id:       v.Id,
			Amount:   v.Amount,
			Price:    v.Coupon.Price,
			CouponId: v.CouponId,
		})
	}
	return response.ActiveResp{List: res}
}

type ActiveAdd struct {
	request.AddActive
}

func (this *ActiveAdd) Add() error {
	m := model.CouponActivity{}
	if this.CouponId == 0 {
		return errors.New("优惠券Id不能为空")
	}
	c := model.Coupon{Id: this.CouponId}
	if !c.Get() {
		return errors.New("优惠券不存在")
	}
	if this.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("满多少不能为0")
	}
	m.Amount = this.Amount
	m.CouponId = this.CouponId
	fmt.Println(m)
	return m.Insert()
}

type DelActive struct {
	request.DelActive
}

func (this *DelActive) Del() error {
	if this.Id == 0 {
		return errors.New("活动Id不能为空")
	}
	m := model.CouponActivity{Id: this.Id}
	if !m.Get() {
		return errors.New("活动不存在")
	}

	return m.Del()
}
