package customer_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type Status struct {
	Code    int    `json:"code"`
	Mesaage string `json:"mesaage"`
}

type CustomReponseSingle struct {
	Status *Status          `json:"status"`
	Data   *RequestCustomer `json:"data"`
}
type CustomReponseCollection struct {
	Status *Status            `json:"status"`
	Data   []*RequestCustomer `json:"data"`
}

type RequestCustomer struct {
	CustomerID  string         `json:"customer_id"`
	Name        string         `json:"name"`
	Alamat      string         `json:"alamat"`
	PhoneNumber string         `json:"phone_number"`
	CreatedTime string         `json:"created_time"`
	Coupon      *RequestCoupon `json:"coupon"`
}

type RequestCoupon struct {
	CouponID    string `json:"coupon_id"`
	Types       string `json:"discount"`
	ExpiredDate string `json:"expired_date"`
	CustomerID  string `json:"customer_id"`
}

func MapResponseListCustomer(dataCustomer []*entity.Customer, code int, message string) ([]byte, error) {
	listResp := make([]*RequestCustomer, 0)
	for _, data := range dataCustomer {
		resp := &RequestCustomer{
			CustomerID:  data.GetCustomerID(),
			Name:        data.GetName(),
			Alamat:      data.GetAlamat(),
			PhoneNumber: data.GetPhoneNumber(),
			CreatedTime: data.GetCreatedTime(),
			Coupon: &RequestCoupon{
				CouponID:    data.GetDataCoupon().GetCouponID(),
				Types:       data.GetDataCoupon().GetTypes(),
				ExpiredDate: data.GetDataCoupon().GetExpiredDate(),
				CustomerID:  data.GetDataCoupon().GetCustomerID(),
			},
		}
		listResp = append(listResp, resp)
	}

	httpResponse := &CustomReponseCollection{
		Status: &Status{
			Code:    code,
			Mesaage: message,
		},
		Data: listResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
