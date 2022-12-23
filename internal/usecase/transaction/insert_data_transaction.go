package transaction

import (
	"context"
	"errors"
	"final-project/domain/entity"
)

func (tr UsecaseTransactionInteractor) InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (string, error) {
	// coupon validation
	types, errCouponValidation := tr.repoCoupon.CouponValidation(ctx, dataTransaction)
	if errCouponValidation != nil {
		return "", errCouponValidation
	}
	// before insert, sum total revenue from items and update value
	totalRevenue := dataTransaction.SumTotalRevenue(types)
	if dataTransaction.GetCouponID() != "" && totalRevenue == 0 {
		return "", errors.New("list items not contain items with this coupon criteria")
	}
	traxID, err := tr.repoTransaction.InsertDataTransaction(ctx, dataTransaction)

	for _, item := range dataTransaction.GetTransactionItems() {
		errTransactionItem := tr.repoTransactionItems.InsertDataTransactionItems(ctx, item, dataTransaction.GetTransactionID())
		if errTransactionItem != nil {
			return "", errTransactionItem
		}
	}

	if err != nil {
		return "", err
	}

	if errCouponValidation == nil && totalRevenue != 0 {
		errUpdateStatus := tr.repoCoupon.UpdateCouponStatus(ctx, dataTransaction.GetCouponID(), 1)
		if errUpdateStatus != nil {
			errDeleteTrax := tr.repoTransaction.DeleteTransactionById(ctx, dataTransaction.GetTransactionID())
			if errDeleteTrax != nil {
				return "", errDeleteTrax
			}
			errDeleteItems := tr.repoTransactionItems.DeleteItemsByTransactionId(ctx, dataTransaction.GetTransactionID())
			if errDeleteItems != nil {
				return "", errDeleteItems
			}

			return "", errUpdateStatus
		}
	}

	return traxID, nil
}
