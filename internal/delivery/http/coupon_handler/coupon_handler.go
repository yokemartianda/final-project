package coupon_handler

import (
	"context"
	"final-project/domain/repository"
)

type CouponHandler struct {
	ctx          context.Context
	repoCoupon   repository.InterfaceRepoCoupon
	repoCustomer repository.InterfaceRepoCustomer
}

func NewCouponHandler(ctx context.Context, repoCoupon repository.InterfaceRepoCoupon, repoCustomer repository.InterfaceRepoCustomer) *CouponHandler {
	return &CouponHandler{
		ctx:          ctx,
		repoCoupon:   repoCoupon,
		repoCustomer: repoCustomer,
	}
}
