package entity_test

import (
	"gamestore/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCoupon(t *testing.T) {
	coupon, err := entity.NewCoupon(entity.DTOCoupon{
		CouponID:    "ULTI-RND7821387123456",
		CustomerID:  2,
		Discount:    30,
		ExpiredDate: "2023-11-28",
	})

	assert.Equal(t, "ULTI-RND7821387123456", coupon.GetCouponID())
	assert.Equal(t, 2, coupon.GetCustomerID())
	assert.Equal(t, 30, coupon.GetDiscount())
	assert.Equal(t, "2023-11-28", coupon.GetExpiredDate())
	assert.Nil(t, err)
}