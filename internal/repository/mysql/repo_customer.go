package mysql

import (
	"context"
	"database/sql"
	"final-project/domain/entity"
	"final-project/internal/repository/mysql/mapper"
	"fmt"
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

	insertQuery := "INSERT INTO customer (customer_id, name, alamat, phone_number, created_time)" +
		"VALUES( ?, ?, ?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCustomer.GetCustomerID(), dataCustomer.GetName(), dataCustomer.GetAlamat(), dataCustomer.GetPhoneNumber(), dataCustomer.GetCreatedTime())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

func (c *CustomerMysqlInteractor) GetListCustomer(ctx context.Context) ([]*entity.Customer, error) {
	var (
		errMysql error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT customer_id, name, alamat, phone_number, created_time FROM customer"
	rows, errMysql := c.db.QueryContext(ctx, sqlQuery)

	if errMysql != nil {
		return nil, errMysql
	}

	dataCustomerCollection := make([]*entity.Customer, 0)
	for rows.Next() {
		var (
			customerID  string
			name        string
			alamat      string
			phoneNumber string
			createdTime string
		)

		err := rows.Scan(&customerID, &name, &alamat, &phoneNumber, &createdTime)
		if err != nil {
			return nil, err
		}
		dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", createdTime)
		if errParse != nil {
			return nil, errParse
		}
		createdTime = dateParse.Format("2006-01-02")

		dataCustomer, errCustomer := mapper.DataCustomerDbToEntity(entity.DTOCustomer{
			CustomerID:  customerID,
			Name:        name,
			Alamat:      alamat,
			PhoneNumber: phoneNumber,
			CreatedTime: createdTime,
		})

		if errCustomer != nil {
			return nil, errCustomer
		}

		dataCustomerCollection = append(dataCustomerCollection, dataCustomer)

	}
	defer rows.Close()

	return dataCustomerCollection, nil
}

func (c *CustomerMysqlInteractor) GetCustomerById(ctx context.Context, customer_id string) (*entity.Customer, error) {
	var (
		errMysql    error
		customerid  string
		name        string
		alamat      string
		phoneNumber string
		createdTime string
	)
	_, cancel := context.WithTimeout(ctx, 60*time.Second)

	defer cancel()
	fmt.Println(customer_id)
	sqlQuery := "SELECT customer_id, name, alamat, phone_number, created_time FROM customer WHERE customer_id = ?"
	errMysql = c.db.QueryRowContext(ctx, sqlQuery, customer_id).Scan(&customerid, &name, &alamat, &phoneNumber, &createdTime)
	if errMysql != nil {
		return nil, errMysql
	}
	dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", createdTime)
	if errParse != nil {
		return nil, errParse
	}
	createdTime = dateParse.Format("2006-01-02")

	dataCustomer, errCustomer := entity.NewCustomer(entity.DTOCustomer{
		CustomerID:  customer_id,
		Name:        name,
		Alamat:      alamat,
		PhoneNumber: phoneNumber,
		CreatedTime: createdTime,
	})
	if errCustomer != nil {
		return nil, errCustomer
	}

	return dataCustomer, nil
}
