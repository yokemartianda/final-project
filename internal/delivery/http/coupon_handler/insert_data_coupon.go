package coupon_handler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"fmt"
	"net/http"
)

func (c *CouponHandler) StoreDataCoupon(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCoupon
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(req)

	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	coupon, err := entity.NewCoupon(entity.DTOCoupon{
		CouponID:    req.CouponID,
		CustomerID:  req.CustomerID,
		Discount:    req.Discount,
		ExpiredDate: req.ExpiredDate,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Build Data"))
		return
	}

	errInsert := c.repoCoupon.InsertDataCoupon(c.ctx, coupon)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "SUKSES INSERT DATA")

}

//
