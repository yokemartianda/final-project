package coupon_handler

import (
	"final-project/internal/delivery/http_response/coupon_response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *CouponHandler) GetCouponByIdCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := r.URL.Query()

	dataCoupon, err := c.repoCoupon.GetCouponByIdCustomer(c.ctx, vars["idcustomer"])
	fmt.Println(query.Get("include"))
	if query.Get("include") == "customer" {
		dataCustomer, errCust := c.repoCustomer.GetCustomerByID(c.ctx, dataCoupon.GetCustomerID())
		fmt.Println(dataCustomer)
		if errCust != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("CUSTOMER TIDAK DI TEMUKAN"))
			return
		}

		dataCoupon = dataCoupon.AddDataCustomer(dataCustomer)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := coupon_response.MapResponseCoupon(dataCoupon, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
		return
	}
	w.WriteHeader(200)
	w.Write(response)

}
