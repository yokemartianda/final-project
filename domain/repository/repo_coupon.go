package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCoupon interface {
	InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) error
	GetCouponById(ctx context.Context, id_coupon string) (*entity.Coupon, error)
}
