package mysql

import (
	"context"
	"database/sql"
	"final-project/domain/entity"
	"time"
)

type TransactionMysqlInteractor struct {
	db *sql.DB
}

func NewTransactionMysql(db *sql.DB) *TransactionMysqlInteractor {
	return &TransactionMysqlInteractor{
		db: db,
	}
}

func (m *TransactionMysqlInteractor) InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (int64, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO transaction (transaction_id, customer_id, revenue, coupon_id, purchase_date)" +
		"VALUES(?, ?, ?, ?, ?)"

	res, errMysql := m.db.Exec(insertQuery, dataTransaction.GetTransactionID(), dataTransaction.GetCustomerID(),
		dataTransaction.GetRevenue(), dataTransaction.GetCouponID(), dataTransaction.GetPurchaseDate())

	lid, _ := res.LastInsertId()

	if errMysql != nil {
		return 0, errMysql
	}

	return lid, nil
}

func (m *TransactionMysqlInteractor) GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT transaction_id, transaction.customer_id, customer.name, revenue, coupon_id, purchase_date FROM transaction " +
		"LEFT JOIN customer ON transaction.customer_id = customer.customer_id ORDER BY purchase_date LIMIT ?"
	rows, errMysql := m.db.QueryContext(ctx, sqlQuery, limit)
	if errMysql != nil {
		return nil, errMysql
	}

	transactionCollection := make([]*entity.Transaction, 0)
	for rows.Next() {
		var (
			transaction_id string
			customer_id    string
			customer_name  string
			revenue        int
			coupon_id      string
			purchase_date  string
		)

		errTransaction := rows.Scan(&transaction_id, &customer_id, &customer_name, &revenue, &coupon_id, &purchase_date)

		if errTransaction != nil {
			return nil, errTransaction
		}

		dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", purchase_date)
		if errParse != nil {
			return nil, errParse
		}
		purchase_date = dateParse.Format("2006-01-02")

		transaction, errNewTransaction := entity.NewTransaction(entity.DTOTransaction{
			TransactionID: transaction_id,
			CustomerID:    customer_id,
			CustomerName:  customer_name,
			Revenue:       revenue,
			CouponID:      coupon_id,
			PurchaseDate:  purchase_date,
		})

		if errNewTransaction != nil {
			return nil, errNewTransaction
		}
		transactionCollection = append(transactionCollection, transaction)
	}
	defer rows.Close()

	return transactionCollection, nil
}
