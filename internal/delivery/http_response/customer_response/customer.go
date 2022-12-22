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
	Status *Status           `json:"status"`
	Data   *ResponseCustomer `json:"data"`
}
type CustomResponseCollection struct {
	Status *Status             `json:"status"`
	Data   []*ResponseCustomer `json:"data"`
}

type ResponseCustomer struct {
	CustomerID  string            `json:"customer_id"`
	Name        string            `json:"name"`
	Alamat      string            `json:"alamat"`
	PhoneNumber string            `json:"phone_number"`
	CreatedTime string            `json:"created_time"`
	Coupon      []*ResponseCoupon `json:"coupon"`
}

type ResponseCoupon struct {
	CouponID    string `json:"coupon_id"`
	Types       string `json:"types"`
	ExpiredDate string `json:"expired_date"`
	CustomerID  string `json:"customer_id"`
	Status      int    `json:"status"`
}

func MapResponseCustomer(dataCustomer *entity.Customer, code int, message string) ([]byte, error) {

	listCoupon := make([]*ResponseCoupon, 0)
	for _, dataCoupon := range dataCustomer.GetDataCoupon() {
		responseCoupon := &ResponseCoupon{
			CouponID:    dataCoupon.GetCouponID(),
			Types:       dataCoupon.GetTypes(),
			ExpiredDate: dataCoupon.GetExpiredDate(),
			CustomerID:  dataCoupon.GetCustomerID(),
			Status:      dataCoupon.GetStatus(),
		}
		listCoupon = append(listCoupon, responseCoupon)
	}

	customer := &ResponseCustomer{
		CustomerID:  dataCustomer.GetCustomerID(),
		Name:        dataCustomer.GetName(),
		Alamat:      dataCustomer.GetAlamat(),
		PhoneNumber: dataCustomer.GetPhoneNumber(),
		CreatedTime: dataCustomer.GetCreatedTime(),
		Coupon:      listCoupon,
	}

	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Mesaage: message,
		},
		Data: customer,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

func MapResponseListCustomer(listCustomer []*entity.Customer, code int, message string) ([]byte, error) {
	lisResp := make([]*ResponseCustomer, 0)
	for _, data := range listCustomer {
		rsp := &ResponseCustomer{
			CustomerID:  data.GetCustomerID(),
			Name:        data.GetName(),
			Alamat:      data.GetAlamat(),
			PhoneNumber: data.GetPhoneNumber(),
			CreatedTime: data.GetCreatedTime(),
		}
		lisResp = append(lisResp, rsp)
	}
	httpResponse := &CustomResponseCollection{
		Status: &Status{
			Code:    code,
			Mesaage: message,
		},
		Data: lisResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
