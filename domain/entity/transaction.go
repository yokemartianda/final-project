package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	transactionID int
	customerID    int
	productID     int
	quantity      int
	revenue       int
	couponID      string
	purchaseDate  time.Time
}

type DTOTransaction struct {
	TransactionID int
	CustomerID    int
	ProductID     int
	Quantity      int
	Revenue       int
	CouponID      string
	PurchaseDate  string
}

func NewTransaction(dto DTOTransaction) (*Transaction, error) {
	if dto.CustomerID == 0 {
		return nil, errors.New("customer ID cannot be empty")
	}
	if dto.ProductID == 0 {
		return nil, errors.New("product ID cannot be empty")
	}
	if dto.Quantity == 0 {
		return nil, errors.New("quantity cannot be empty")
	}
	if dto.Revenue == 0 {
		return nil, errors.New("revenue cannot be empty")
	}
	if dto.PurchaseDate == "" {
		return nil, errors.New("purchase date cannot be empty")
	}

	convertPurchaseDate, _ := time.Parse("2006-01-02", dto.PurchaseDate)

	transaction := &Transaction{
		transactionID: dto.TransactionID,
		customerID:    dto.CustomerID,
		productID:     dto.ProductID,
		quantity:      dto.Quantity,
		revenue:       dto.Revenue,
		couponID:      dto.CouponID,
		purchaseDate:  convertPurchaseDate,
	}

	return transaction, nil
}

func (tr *Transaction) GetTransactionID() int {
	return tr.transactionID
}

func (tr *Transaction) GetCustomerID() int {
	return tr.customerID
}

func (tr *Transaction) GetProductID() int {
	return tr.productID
}

func (tr *Transaction) GetQuantity() int {
	return tr.quantity
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
