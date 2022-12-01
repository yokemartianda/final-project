package customer_hendler

import (
	"context"
	"final-project/internal/repository/mysql"
)

type CustomerHandler struct {
	ctx          context.Context
	repoCustomer *mysql.CustomerMysqlInteractor
}

func NewCustomerHandler(ctx context.Context, repoCustomer *mysql.CustomerMysqlInteractor) *CustomerHandler {
	return &CustomerHandler{
		ctx:          ctx,
		repoCustomer: repoCustomer,
	}
}
