package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCoupon interface {
	InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) error
	GetCouponByCustomerId(ctx context.Context, customer_id string) ([]*entity.Coupon, error)
}
