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
