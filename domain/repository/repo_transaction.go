package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoTransaction interface {
	InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (string, error)
	GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error)
	DeleteTransactionById(ctx context.Context, transaction_id string) error
}
