package mysql

import (
	"context"
	"database/sql"
	"final-project/domain/entity"
	"time"
)

type TransactionItemsMysqlInteractor struct {
	db *sql.DB
}

func NewTransactionItemsMysql(db *sql.DB) *TransactionItemsMysqlInteractor {
	return &TransactionItemsMysqlInteractor{
		db: db,
	}
}

func (m *TransactionItemsMysqlInteractor) InsertDataTransactionItems(ctx context.Context, dataTransactionItems *entity.TransactionItems, transactionID string) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO transaction_items (transaction_id, criteria_id, revenue_item, date_created)" +
		"VALUES(?, ?, ?, ?)"

	_, errMysql = m.db.Exec(insertQuery, transactionID, dataTransactionItems.GetCriteriaID(),
		dataTransactionItems.GetRevenueItem(), dataTransactionItems.GetDateCreated())

	if errMysql != nil {
		return errMysql
	}

	return nil
}
