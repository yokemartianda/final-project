package customer

import "final-project/domain/repository"

type UsecaseCostumerInteractor struct {
	repoCustomer repository.InterfaceRepoCustomer
}

func NewCustomerUsecase(repoCostumer repository.InterfaceRepoCustomer) *UsecaseCostumerInteractor {
	return &UsecaseCostumerInteractor{
		repoCustomer: repoCostumer,
	}
}
