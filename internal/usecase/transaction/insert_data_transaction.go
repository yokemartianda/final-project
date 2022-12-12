package transaction

import (
	"context"
	"final-project/domain/entity"
)

func (tr UsecaseTransactionInteractor) InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) error {
	_, err := tr.repoTransaction.InsertDataTransaction(ctx, dataTransaction)

	for _, item := range dataTransaction.GetTransactionItems() {
		errTransactionItem := tr.repoTransactionItems.InsertDataTransactionItems(ctx, item, dataTransaction.GetTransactionID())
		if errTransactionItem != nil {
			return errTransactionItem
		}
	}

	if err != nil {
		return err
	}

	return nil
}
