package transaction

import (
	"context"
	"final-project/domain/entity"
)

func (tr UsecaseTransactionInteractor) InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) error {
	// coupon validation
	types, errCouponValidation := tr.repoCoupon.CouponValidation(ctx, dataTransaction)
	if errCouponValidation != nil {
		return errCouponValidation
	}
	// before insert, sum total revenue from items and update value
	dataTransaction.SumTotalRevenue(types)
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
