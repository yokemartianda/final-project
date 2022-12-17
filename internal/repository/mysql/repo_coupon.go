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

	insertQuery := "INSERT INTO coupon (coupon_id, discount, expired_date)" +
		"VALUES(?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCoupon.GetCouponID(), dataCoupon.GetDiscount(), dataCoupon.GetExpiredDate())

	if errMysql != nil {
		return errMysql
	}
	return nil
}
