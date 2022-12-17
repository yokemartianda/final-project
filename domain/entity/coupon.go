package entity

import (
	"errors"
	"time"
)

type Coupon struct {
	couponID    string
	discount    int
	expiredDate time.Time
}

type DTOCoupon struct {
	CouponID    string
	Discount    int
	ExpiredDate string
}

func NewCoupon(dto DTOCoupon) (*Coupon, error) {
	if dto.Discount == 0 {
		return nil, errors.New("discount cannot be empty")
	}
	if dto.ExpiredDate == "" {
		return nil, errors.New("expired date cannot be empty")
	}

	convertExpiredDate, _ := time.Parse("2006-01-02", dto.ExpiredDate)

	coupon := &Coupon{
		couponID:    dto.CouponID,
		discount:    dto.Discount,
		expiredDate: convertExpiredDate,
	}

	return coupon, nil
}

func (c *Coupon) GetCouponID() string {
	return c.couponID
}

func (c *Coupon) GetDiscount() int {
	return c.discount
}

func (c *Coupon) GetExpiredDate() string {
	return c.expiredDate.Format("2006-01-02")
}

//
