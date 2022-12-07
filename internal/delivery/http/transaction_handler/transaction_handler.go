package transaction_handler

import (
	"context"
	"final-project/domain/repository"
)

type TransactionHandler struct {
	ctx             context.Context
	repoTransaction repository.InterfaceRepoTransaction
}

func NewTransactionHandler(ctx context.Context, repoTransaction repository.InterfaceRepoTransaction) *TransactionHandler {
	return &TransactionHandler{
		repoTransaction: repoTransaction,
		ctx:             ctx,
	}
}
