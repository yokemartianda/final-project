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

// GetCouponByIdCustomer implements repository.InterfaceRepoCoupon
func (*CouponMysqlInteractor) GetCouponByIdCustomer(ctx context.Context, idCustomer string) (*entity.Coupon, error) {
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

	insertQuery := "INSERT INTO coupon (customer_ID, discount, expired_date)" +
		"VALUES(?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCoupon.GetCouponID(), dataCoupon.GetDiscount(), dataCoupon.GetExpiredDate())

	if errMysql != nil {
		return errMysql
	}
	return nil
}

//
