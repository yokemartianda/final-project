package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCoupon interface {
	InsertDataCoupon(ctx context.Context, dataCustomer *entity.Coupon) error
	GetCouponByIdCustomer(ctx context.Context, idCustomer string) (*entity.Coupon, error)
}
