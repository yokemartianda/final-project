package http_request

type RequestCoupon struct {
	CouponID    string `json:"coupon_id"`
	Types       string `json:"types"`
	ExpiredDate string `json:"expired_date"`
	CustomerID  string `json:"customer_id"`
	Status      int    `json:"status"`
}
