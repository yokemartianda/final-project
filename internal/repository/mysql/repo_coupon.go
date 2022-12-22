package mysql

import (
	"context"
	"database/sql"
	"errors"
	"final-project/domain/entity"
	"time"
)

type CouponMysqlInteractor struct {
	db *sql.DB
}

func NewCouponMysql(db *sql.DB) *CouponMysqlInteractor {
	return &CouponMysqlInteractor{
		db: db,
	}
}

func (c *CouponMysqlInteractor) GetCouponById(ctx context.Context, id_coupon string) (*entity.Coupon, error) {
	var (
		errMysql     error
		coupon_id    string
		types        string
		expired_date string
		customer_id  string
		status       int
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT coupon_id, types, expired_date, customer_id, status FROM coupon WHERE coupon_id = ?"
	errMysql = c.db.QueryRowContext(ctx, sqlQuery, id_coupon).Scan(&coupon_id, &types, &expired_date, &customer_id, &status)

	if errMysql != nil {
		return nil, errMysql
	}

	coupon, errCoupon := entity.NewCoupon(entity.DTOCoupon{
		CouponID:    coupon_id,
		Types:       types,
		CustomerID:  customer_id,
		ExpiredDate: expired_date,
		Status:      status,
	})

	if errCoupon != nil {
		return nil, errCoupon
	}

	return coupon, nil
}

func (c *CouponMysqlInteractor) InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO coupon (coupon_id, types, expired_date, customer_id, status)" +
		"VALUES(?, ?, ?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCoupon.GetCouponID(), dataCoupon.GetTypes(), dataCoupon.GetExpiredDate(), dataCoupon.GetCustomerID(), dataCoupon.GetStatus())

	if errMysql != nil {
		return errMysql
	}
	return nil
}

func (c *CouponMysqlInteractor) CouponValidation(ctx context.Context, dataTransaction *entity.Transaction) (string, error) {
	coupon, errCoupon := c.GetCouponById(ctx, dataTransaction.GetCouponID())
	if errCoupon != nil {
		if errCoupon == sql.ErrNoRows {
			return "", errors.New("coupon not found")
		}
		return "", errCoupon
	}

	if coupon.GetCustomerID() != dataTransaction.GetCustomerID() {
		return "", errors.New("this coupon not eligible for this customer")
	}

	if coupon.GetStatus() == 1 {
		return "", errors.New("this coupon was used")
	}

	today := time.Now().Format("2006-01-02")
	if coupon.GetExpiredDate() < today {
		return "", errors.New("coupon expired")
	}

	return "", nil
}
