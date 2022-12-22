package mysql

import (
	"context"
	"database/sql"
	"final-project/domain/entity"
	"time"
)

type CouponMysqlInteractor struct {
	db *sql.DB
}

// GetCouponById implements repository.InterfaceRepoCoupon
func (*CouponMysqlInteractor) GetCouponById(ctx context.Context, id_coupon string) (*entity.Coupon, error) {
	panic("unimplemented")
}

func NewCouponMysql(db *sql.DB) *CouponMysqlInteractor {
	return &CouponMysqlInteractor{
		db: db,
	}
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

func (c *CouponMysqlInteractor) GetCouponByCustomerId(ctx context.Context, customer_id string) ([]*entity.Coupon, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT coupon_id, types, expired_date, status, customer_id FROM coupon WHERE customer_id = ?"
	rows, errMysql := c.db.QueryContext(ctx, sqlQuery, customer_id)
	if errMysql != nil {
		return nil, errMysql
	}

	couponCollection := make([]*entity.Coupon, 0)
	for rows.Next() {
		var (
			coupon_id    string
			types        string
			expired_date string
			status       int
			customer_id  string
		)

		errCoupon := rows.Scan(&coupon_id, &types, &expired_date, &status, &customer_id)

		if errCoupon != nil {
			return nil, errCoupon
		}
		coupon, errNewCoupon := entity.NewCoupon(entity.DTOCoupon{
			CouponID:    coupon_id,
			Types:       types,
			ExpiredDate: expired_date,
			CustomerID:  customer_id,
			Status:      status,
		})
		if errNewCoupon != nil {
			return nil, errNewCoupon
		}
		couponCollection = append(couponCollection, coupon)
	}
	defer rows.Close()

	return couponCollection, nil
}
