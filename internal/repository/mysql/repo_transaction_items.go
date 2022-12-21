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

func (m *TransactionItemsMysqlInteractor) GetItemsByTransactionID(ctx context.Context, transaction_id string) ([]*entity.TransactionItems, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT item_id, transaction_id, criteria_id, revenue_item, date_created FROM transaction_items WHERE transaction_id = ?"
	rows, errMysql := m.db.QueryContext(ctx, sqlQuery, transaction_id)
	if errMysql != nil {
		return nil, errMysql
	}

	itemsCollection := make([]*entity.TransactionItems, 0)
	for rows.Next() {
		var (
			item_id       int
			transactionID string
			criteria_id   int
			revenue_item  int
			date_created  string
		)

		errTransaction := rows.Scan(&item_id, &transactionID, &criteria_id, &revenue_item, &date_created)

		if errTransaction != nil {
			return nil, errTransaction
		}
		item, errNewItem := entity.NewTransactionItems(entity.DTOTransactionItems{
			ItemID:        item_id,
			TransactionID: transactionID,
			CriteriaID:    criteria_id,
			RevenueItem:   revenue_item,
			DateCreated:   date_created,
		})

		if errNewItem != nil {
			return nil, errNewItem
		}
		itemsCollection = append(itemsCollection, item)
	}
	defer rows.Close()

	return itemsCollection, nil
}
