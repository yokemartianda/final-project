package entity

import (
	"errors"
	"time"
)

type Coupon struct {
	couponID    string
	customerID  string
	discount    int
	expiredDate time.Time
	customer    *Customer
}

type DTOCoupon struct {
	CouponID    string
	CustomerID  string
	Discount    int
	ExpiredDate string
	Customer    *Customer
}

func NewCoupon(dto DTOCoupon) (*Coupon, error) {
	if dto.CustomerID == "" {
		return nil, errors.New("customer ID cannot be empty")
	}
	if dto.Discount == 0 {
		return nil, errors.New("discount cannot be empty")
	}
	if dto.ExpiredDate == "" {
		return nil, errors.New("expired date cannot be empty")
	}

	convertExpiredDate, _ := time.Parse("2006-01-02", dto.ExpiredDate)

	coupon := &Coupon{
		couponID:    dto.CouponID,
		customerID:  dto.CustomerID,
		discount:    dto.Discount,
		expiredDate: convertExpiredDate,
	}

	return coupon, nil
}

func (c *Coupon) AddDataCustomer(cust *Customer) *Coupon {
	c.customer = cust

	return c
}

func (c *Coupon) GetCouponID() string {
	return c.couponID
}

func (c *Coupon) GetCustomerID() string {
	return c.customerID
}

func (c *Coupon) GetDiscount() int {
	return c.discount
}

func (c *Coupon) GetExpiredDate() string {
	return c.expiredDate.Format("2006-01-02")
}

func (c *Coupon) GetDataCustomer() *Customer {
	return c.customer
}

//
