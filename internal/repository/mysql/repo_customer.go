package mysql

import (
	"context"
	"database/sql"
	"final-project/domain/entity"
	"time"
)

type CustomerMysqlInteractor struct {
	db *sql.DB
}

func NewCustomerMysql(db *sql.DB) *CustomerMysqlInteractor {
	return &CustomerMysqlInteractor{
		db: db,
	}
}

func (c *CustomerMysqlInteractor) InsertDataCustomer(ctx context.Context, dataCustomer *entity.Customer) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO customer (user_id, name, alamat, phone_number, created_time)" +
		"VALUES(?, ?, ?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCustomer.GetCustomerID(), dataCustomer.GetName(), dataCustomer.GetAlamat(), dataCustomer.GetPhoneNumber(), dataCustomer.GetCreatedTime())

	if errMysql != nil {
		return errMysql
	}

	return nil
}
