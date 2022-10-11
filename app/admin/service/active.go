package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/model"
)

type CouponList struct {
	request.Request
}

func (this *CouponList) CouponList() response.CouponResp {
	res := make([]response.Coupon, 0)
	m := model.Coupon{}
	s := m.List()
	for i := range s {
		res = append(res, response.Coupon{
			Id:    s[i].ID,
			Price: float64(s[i].Price) / model.UNITY,
		})
	}
	return response.CouponResp{List: res}
}

type CouponAdd struct {
	request.AddCoupon
}

func (this *CouponAdd) Add() error {
	m := model.Coupon{}
	if this.Price == 0 {
		return errors.New("金额不能为0")
	}
	m.Price = int64(this.Price * model.UNITY)
	return m.Insert()
}

type ActiveList struct {
	request.Request
}

func (this *ActiveList) PageList() response.ActiveResp {
	res := make([]response.Active, 0)
	m := model.FullDelivery{}
	s := m.List()
	for i := range s {
		res = append(res, response.Active{
			Id:       s[i].ID,
			Amout:    float64(s[i].Amout) / model.UNITY,
			Price:    float64(s[i].Coupon.Price) / model.UNITY,
			CouponId: s[i].CouponId,
		})
	}
	return response.ActiveResp{List: res}
}

type ActiveAdd struct {
	request.AddActive
}

func (this *ActiveAdd) Add() error {
	m := model.FullDelivery{}
	if this.CouponId == 0 {
		return errors.New("优惠券ID不能为空")
	}
	c := model.Coupon{ID: this.CouponId}
	if !c.Get() {
		return errors.New("优惠券不存在")
	}
	if this.Amout == 0 {
		return errors.New("满多少不能为0")
	}
	m.Amout = int64(this.Amout * model.UNITY)
	m.CouponId = this.CouponId

	return m.Insert()
}

type DelActive struct {
	request.DelActive
}

func (this *DelActive) Del() error {
	if this.Id == 0 {
		return errors.New("活动ID不能为空")
	}
	m := model.FullDelivery{ID: this.Id}
	if !m.Get() {
		return errors.New("活动不存在")
	}

	return m.Del()
}
