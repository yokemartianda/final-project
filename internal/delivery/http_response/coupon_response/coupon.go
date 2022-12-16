package coupon_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CostumResponseSingle struct {
	Status *Status             `json:"stastus"`
	Data   *ResponseCouponJson `json:"data"`
}

type CostumeResponseCollection struct {
	Status *Status               `json:"status"`
	Data   []*ResponseCouponJson `json:"data"`
}

type ResponseCouponJson struct {
	CouponID    string              `json:"coupon_id"`
	CustomerID  string              `json:"customer_id"`
	Discount    int                 `json:"discount"`
	ExpiredDate string              `json:"expired_date"`
	Customer    *DetailCustomerJson `json:"customer,omitempty"`
}

type DetailCustomerJson struct {
	CustomerID  string `json:"customer_id"`
	Name        string `json:"name"`
	Alamat      string `json:"alamat"`
	PhoneNumber string `json:"phone_number"`
	CreatedTime string `json:"created_date"`
}

func MapResponseListCoupon(dataCoupon []*entity.Coupon, code int, message string) ([]byte, error) {
	listResp := make([]*ResponseCouponJson, 0)
	for _, data := range dataCoupon {
		resp := &ResponseCouponJson{
			CouponID:    data.GetCouponID(),
			CustomerID:  data.GetCustomerID(),
			Discount:    data.GetDiscount(),
			ExpiredDate: data.GetCouponID(),
			Customer: &DetailCustomerJson{
				CustomerID:  data.GetDataCustomer().GetCustomerID(),
				Name:        data.GetDataCustomer().GetName(),
				Alamat:      data.GetDataCustomer().GetAlamat(),
				PhoneNumber: data.GetDataCustomer().GetPhoneNumber(),
				CreatedTime: data.GetDataCustomer().GetCreatedTime(),
			},
		}

		listResp = append(listResp, resp)
	}
	httpResponse := &CostumeResponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil

}

func MapResponseCoupon(dataCoupon *entity.Coupon, code int, message string) ([]byte, error) {
	var resp *ResponseCouponJson
	if dataCoupon != nil {
		resp = &ResponseCouponJson{
			CouponID:    dataCoupon.GetCouponID(),
			CustomerID:  dataCoupon.GetCustomerID(),
			Discount:    dataCoupon.GetDiscount(),
			ExpiredDate: dataCoupon.GetExpiredDate(),
		}
		if dataCoupon.GetDataCustomer() != nil {
			resp.Customer = &DetailCustomerJson{
				CustomerID:  dataCoupon.GetDataCustomer().GetCustomerID(),
				Name:        dataCoupon.GetDataCustomer().GetName(),
				Alamat:      dataCoupon.GetDataCustomer().GetAlamat(),
				PhoneNumber: dataCoupon.GetDataCustomer().GetPhoneNumber(),
				CreatedTime: dataCoupon.GetDataCustomer().GetCreatedTime(),
			}
		}
	}
	httpResponse := &CostumResponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: resp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
