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
