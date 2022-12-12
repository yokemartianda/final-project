package transaction_handler

import (
	"context"
	"final-project/domain/repository"
	"final-project/domain/usecase"
	"final-project/internal/usecase/transaction"
)

type TransactionHandler struct {
	ctx                context.Context
	repoTransaction    repository.InterfaceRepoTransaction
	transactionUsecase usecase.TransactionService
}

func NewTransactionHandler(ctx context.Context, repoTransaction repository.InterfaceRepoTransaction, repoTransactionItem repository.InterfaceRepoTransactionItem) *TransactionHandler {
	transactionUsecas := transaction.NewTransactionUsecase(repoTransaction, repoTransactionItem)
	return &TransactionHandler{
		repoTransaction:    repoTransaction,
		transactionUsecase: transactionUsecas,
		ctx:                ctx,
	}
}
