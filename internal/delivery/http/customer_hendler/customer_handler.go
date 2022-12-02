package customer_hendler

import (
	"context"
	"final-project/domain/repository"
)

type CustomerHandler struct {
	ctx          context.Context
	repoCustomer repository.InterfaceRepoCustomer
}

func NewCustomerHandler(ctx context.Context, repoCustomer repository.InterfaceRepoCustomer) *CustomerHandler {
	return &CustomerHandler{
		ctx:          ctx,
		repoCustomer: repoCustomer,
	}
}
