package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoTransactionItem interface {
	InsertDataTransactionItems(ctx context.Context, dataTransactionItems *entity.TransactionItems, transactionID string) error
}
