package entity

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	transactionID    string
	customerID       string
	customerName     string
	revenue          int
	couponID         string
	purchaseDate     time.Time
	transactionItems []*TransactionItems
}

type DTOTransaction struct {
	TransactionID    string
	CustomerID       string
	CustomerName     string
	Revenue          int
	CouponID         string
	PurchaseDate     string
	TransactionItems []*TransactionItems
}

func (tr *Transaction) AddDataIDCustomer(customer *Customer) *Transaction {
	tr.customerID = customer.customerID

	return tr
}

func NewTransaction(dto DTOTransaction) (*Transaction, error) {
	if dto.CustomerID == "" {
		return nil, errors.New("customer ID cannot be empty")
	}
	// if dto.Revenue == 0 {
	// 	return nil, errors.New("revenue cannot be empty")
	// }
	if dto.PurchaseDate == "" {
		return nil, errors.New("purchase date cannot be empty")
	}

	convertPurchaseDate, _ := time.Parse("2006-01-02", dto.PurchaseDate)

	transaction := &Transaction{
		transactionID:    dto.TransactionID,
		customerID:       dto.CustomerID,
		customerName:     dto.CustomerName,
		revenue:          dto.Revenue,
		couponID:         dto.CouponID,
		purchaseDate:     convertPurchaseDate,
		transactionItems: dto.TransactionItems,
	}

	return transaction, nil
}

func (tr *Transaction) SetUniqTransactionID() *Transaction {
	rand.Seed(time.Now().UnixNano())
	min := 10000000
	max := 99999999
	valueString := strconv.Itoa(rand.Intn(max-min+1) + min)
	tr.transactionID = "TRAX" + valueString

	return tr
}

func (tr *Transaction) GetTransactionID() string {
	return tr.transactionID
}

func (tr *Transaction) GetCustomerID() string {
	return tr.customerID
}

func (tr *Transaction) GetRevenue() int {
	return tr.revenue
}

func (tr *Transaction) GetCouponID() string {
	return tr.couponID
}

func (tr *Transaction) GetPurchaseDate() string {
	return tr.purchaseDate.Format("2006-01-02")
}

func (tr *Transaction) GetTransactionItems() []*TransactionItems {
	return tr.transactionItems
}

func (tr *Transaction) SetTransactionItems(transactionItems []*TransactionItems) *Transaction {
	tr.transactionItems = transactionItems
	return tr
}

func (tr *Transaction) GetCustomerName() string {
	return tr.customerName
}

func (tr *Transaction) SumTotalRevenue() int {
	var totalRevenue int
	for _, item := range tr.transactionItems {
		totalRevenue += item.GetRevenueItem()
	}
	tr.revenue = totalRevenue
	return totalRevenue
}
