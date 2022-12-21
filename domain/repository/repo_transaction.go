package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoTransaction interface {
	InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (int64, error)
	GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error)
}
