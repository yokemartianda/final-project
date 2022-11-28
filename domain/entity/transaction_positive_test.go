package entity_test

import (
	"final-project/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		TransactionID: 1,
		CustomerID:    3,
		ProductID:     5,
		Quantity:      1,
		Revenue:       5000000,
		CouponID:      "",
		PurchaseDate:  "2022-11-28",
	})

	assert.Equal(t, 1, transaction.GetTransactionID())
	assert.Equal(t, 3, transaction.GetCustomerID())
	assert.Equal(t, 5, transaction.GetProductID())
	assert.Equal(t, 1, transaction.GetQuantity())
	assert.Equal(t, 5000000, transaction.GetRevenue())
	assert.Equal(t, "2022-11-28", transaction.GetPurchaseDate())
	assert.Nil(t, err)
}
