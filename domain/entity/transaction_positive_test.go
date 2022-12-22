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
		CouponID:     "",
		PurchaseDate: "2022-11-28",
	})
	transaction.SetUniqTransactionID()

	transactionItems := make([]*entity.TransactionItems, 0)
	items1, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    2,
		RevenueItem:   1000000,
	})
	items2, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    5,
		RevenueItem:   500000,
	})
	transactionItems = append(transactionItems, items1, items2)
	transaction.SetTransactionItems(transactionItems)
	transaction.SumTotalRevenue()

	fmt.Println(transaction.GetRevenue())
	assert.Equal(t, "CUST-001", transaction.GetCustomerID())
	assert.Equal(t, 1500000, transaction.GetRevenue())
	assert.Equal(t, "2022-11-28", transaction.GetPurchaseDate())
	assert.Nil(t, err)
}
