package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCustomer interface {
	InsertDataCustomer(ctx context.Context, dataCustomer *entity.Customer) error
	GetListCustomerCoupon(ctx context.Context) ([]*entity.Customer, error)
	GetCustomerById(ctx context.Context, customerid string) (*entity.Customer, error)
}
