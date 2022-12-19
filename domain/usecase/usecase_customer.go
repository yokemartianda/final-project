package usecase

import (
	"context"
	"final-project/domain/entity"
)

type CustomerService interface {
	GetListCustomerCoupon(ctx context.Context, include string) ([]*entity.Customer, error)
}
