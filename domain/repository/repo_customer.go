package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCustomer interface {
	InsertDataCustomer(ctx context.Context, dataCustomer *entity.Customer) error
}
