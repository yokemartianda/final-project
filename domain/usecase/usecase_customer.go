package usecase

import (
	"context"
	"final-project/domain/entity"
)

type CustomerService interface {
	GetCustomerById(ctx context.Context, customer_id string) (*entity.Customer, error)
	GetListCustomer(ctx context.Context) ([]*entity.Customer, error)
}
