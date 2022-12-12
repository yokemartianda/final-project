package usecase

import (
	"context"
	"final-project/domain/entity"
)

type TransactionService interface {
	InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) error
}
