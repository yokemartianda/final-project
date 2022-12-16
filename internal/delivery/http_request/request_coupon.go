package http_request

type RequestCoupon struct {
	CouponID    string `json:"coupon_id"`
	CustomerID  string `json:"customer_id"`
	Discount    int    `json:"discount"`
	ExpiredDate string `json:"expired_date"`
}
