package entity

import (
	"errors"
	"time"
)

type Coupon struct {
	couponID    string
	types       string
	expiredDate time.Time
	customerID  string
}

type DTOCoupon struct {
	CouponID    string
	Types       string
	ExpiredDate string
	CustomerID  string
}

func NewCoupon(dto DTOCoupon) (*Coupon, error) {
	if dto.Types == "" {
		return nil, errors.New("Type cannot be empty")
	}
	if dto.ExpiredDate == "" {
		return nil, errors.New("expired date cannot be empty")
	}

	convertExpiredDate, _ := time.Parse("2006-01-02", dto.ExpiredDate)

	coupon := &Coupon{
		couponID:    dto.CouponID,
		types:       dto.Types,
		expiredDate: convertExpiredDate,
		customerID:  dto.CustomerID,
	}

	return coupon, nil
}

func (c *Coupon) GetCouponID() string {
	return c.couponID
}

func (c *Coupon) GetTypes() string {
	return c.types
}
func (c *Coupon) GetCustomerID() string {
	return c.customerID
}

func (c *Coupon) GetExpiredDate() string {
	return c.expiredDate.Format("2006-01-02")
}

//
