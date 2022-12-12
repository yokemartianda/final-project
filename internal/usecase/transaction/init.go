package transaction

import "final-project/domain/repository"

type UsecaseTransactionInteractor struct {
	repoTransaction      repository.InterfaceRepoTransaction
	repoTransactionItems repository.InterfaceRepoTransactionItem
}

func NewTransactionUsecase(repoTransaction repository.InterfaceRepoTransaction, repoTransactionItems repository.InterfaceRepoTransactionItem) *UsecaseTransactionInteractor {
	return &UsecaseTransactionInteractor{
		repoTransaction:      repoTransaction,
		repoTransactionItems: repoTransactionItems,
	}
}
