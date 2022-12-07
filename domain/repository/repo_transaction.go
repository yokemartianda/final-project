package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoTransaction interface {
	InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) error
}
