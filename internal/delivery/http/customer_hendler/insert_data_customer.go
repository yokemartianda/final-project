package customer_hendler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"fmt"
	"net/http"
	"time"
)

func (c *CustomerHandler) StoreDataCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCustomer
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)

	if errDecode != nil {
		fmt.Println(errDecode)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	var date = time.Now()

	customer, err := entity.NewCustomer(entity.DTOCustomer{
		Name:        req.Name,
		Alamat:      req.Alamat,
		PhoneNumber: req.PhoneNumber,
		CreatedTime: date.Format("2006-01-02"),
	})
	customer.SetUniqCustomerID()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	errInsert := c.repoCustomer.InsertDataCustomer(c.ctx, customer)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "SUCCES INSERT DATA")

}
