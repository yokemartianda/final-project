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

	sqlQuery := "SELECT transaction_id, customer_id, revenue, coupon_id, purchase_date FROM transaction LIMIT ?"
	rows, errMysql := m.db.QueryContext(ctx, sqlQuery, limit)
	if errMysql != nil {
		return nil, errMysql
	}

	transactionCollection := make([]*entity.Transaction, 0)
	for rows.Next() {
		var (
			transaction_id string
			customer_id    string
			revenue        int
			coupon_id      string
			purchase_date  string
		)

		errTransaction := rows.Scan(&transaction_id, &customer_id, &revenue, &coupon_id, &purchase_date)

		if errTransaction != nil {
			return nil, errTransaction
		}
		transaction, errNewTransaction := entity.NewTransaction(entity.DTOTransaction{
			TransactionID: transaction_id,
			CustomerID:    customer_id,
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
