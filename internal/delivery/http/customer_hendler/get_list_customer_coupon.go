package customer_hendler

import (
	"context"
	"final-project/internal/delivery/http_response/customer_response"
	"net/http"
)

func (c *CustomerHandler) GetListCustomerCoupon(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		include = r.URL.Query().Get("include")
	)

	listCustomer, err := c.customerUsecase.GetListCustomerCoupon(ctx, include)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := customer_response.MapResponseListCustomer(listCustomer, 200, "Succes")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Mapping data"))
	}
	w.WriteHeader(200)
	w.Write(response)
}
