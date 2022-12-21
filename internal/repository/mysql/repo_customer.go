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

// GetCustomerById implements repository.InterfaceRepoCustomer
func (*CustomerMysqlInteractor) GetCustomerById(ctx context.Context, customerid string) (*entity.Customer, error) {
	panic("unimplemented")
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

	insertQuery := "INSERT INTO customer (coupon_id, customer_id, name, alamat, phone_number, created_time)" +
		"VALUES(?, ?, ?, ?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCustomer.GetCouponID(), dataCustomer.GetCustomerID(), dataCustomer.GetName(), dataCustomer.GetAlamat(), dataCustomer.GetPhoneNumber(), dataCustomer.GetCreatedTime())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

func (c *CustomerMysqlInteractor) GetListCustomerCoupon(ctx context.Context) ([]*entity.Customer, error) {
	var (
		errMysql error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT a.customer_id,a.name,a.alamat,a.phone_number,a.created_date,b.*" +
		"FROM customer a JOIN coupon b ON a.customer_id = b.id"
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
			couponID    string
			types       string
			expiredDate string
		)

		err := rows.Scan(&customerID, &name, &alamat, &phoneNumber, &createdTime, &couponID, &types, &expiredDate)
		if err != nil {
			return nil, err
		}
		coupon, errCoupon := mapper.DataCouponDbToEntity(entity.DTOCoupon{
			CouponID:    couponID,
			Types:       types,
			ExpiredDate: expiredDate,
		})
		if errCoupon != nil {
			return nil, errCoupon
		}

		dataCustomer, errCustomer := mapper.DataCustomerDbToEntity(entity.DTOCustomer{
			CustomerID:  customerID,
			Name:        name,
			Alamat:      alamat,
			PhoneNumber: phoneNumber,
			CreatedTime: createdTime,
			Coupon:      coupon,
		})

		if errCustomer != nil {
			return nil, errCustomer
		}

		dataCustomerCollection = append(dataCustomerCollection, dataCustomer)

	}
	defer rows.Close()

	return dataCustomerCollection, nil
}

func (c *CouponMysqlInteractor) GetCustomerById(ctx context.Context, kodecustomerId string) (*entity.Customer, error) {
	var (
		dataCustomer *entity.Customer
		errMysql     error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := fmt.Sprintf("SELECT a.customer_id,a.name,a.alamat,a.phone_number,a.created_date,b.*"+
		"FROM customer a JOIN coupon b ON a.customer_id = b.id WHERE customer_id ='%s'", kodecustomerId)
	rows, errMysql := c.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}
	for rows.Next() {
		var (
			customerID  string
			name        string
			alamat      string
			phoneNumber string
			createdTime string
			couponID    string
			types       string
			expiredDate string
		)

		err := rows.Scan(&customerID, &name, &alamat, &phoneNumber, &createdTime, &couponID, &types, &expiredDate)
		if err != nil {
			return nil, err
		}
		coupon, errCoupon := mapper.DataCouponDbToEntity(entity.DTOCoupon{
			CouponID:    couponID,
			Types:       types,
			ExpiredDate: expiredDate,
		})
		if errCoupon != nil {
			return nil, errCoupon
		}

		dataCustomer, _ = mapper.DataCustomerDbToEntity(entity.DTOCustomer{
			CustomerID:  customerID,
			Name:        name,
			Alamat:      alamat,
			PhoneNumber: phoneNumber,
			CreatedTime: createdTime,
			Coupon:      coupon,
		})
	}
	defer rows.Close()

	return dataCustomer, nil
}
