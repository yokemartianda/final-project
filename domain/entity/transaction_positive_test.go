package entity_test

import (
	"final-project/domain/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerID:   "CUST-001",
		Revenue:      5000000,
		CouponID:     "",
		PurchaseDate: "2022-11-28",
	})
	transaction.SetUniqTransactionID()

	fmt.Println(transaction.GetTransactionID())
	assert.Equal(t, "CUST-001", transaction.GetCustomerID())
	assert.Equal(t, 5000000, transaction.GetRevenue())
	assert.Equal(t, "2022-11-28", transaction.GetPurchaseDate())
	assert.Nil(t, err)
}
