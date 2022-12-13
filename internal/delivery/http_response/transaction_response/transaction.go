package transaction_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type ResponseTransactionJson struct {
	TransactionID    string                      `json:"transaction_id"`
	CustomerID       string                      `json:"customer_id"`
	Revenue          int                         `json:"revenue"`
	CouponID         string                      `json:"coupon_id"`
	PurchaseDate     string                      `json:"purchase_date"`
	TransactionItems []*ResponseTransactionItems `json:"transaction_items"`
}

type ResponseTransactionItems struct {
	TransactionID string `json:"transaction_id"`
	CriteriaID    int    `json:"criteria_id"`
	RevenueItem   int    `json:"revenue_item"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomReponseSingle struct {
	Status *Status                  `json:"status"`
	Data   *ResponseTransactionJson `json:"data"`
}

type CustomReponseCollection struct {
	Status *Status                    `json:"status"`
	Data   []*ResponseTransactionJson `json:"data"`
}

func MapResponseTransaction(dataTransaction *entity.Transaction, code int, message string) ([]byte, error) {
	var resp *ResponseTransactionJson
	if dataTransaction != nil {
		listItems := make([]*ResponseTransactionItems, 0)

		for _, item := range dataTransaction.GetTransactionItems() {
			transactionItem := &ResponseTransactionItems{
				TransactionID: item.GetTransactionID(),
				CriteriaID:    item.GetCriteriaID(),
				RevenueItem:   item.GetRevenueItem(),
			}

			listItems = append(listItems, transactionItem)
		}

		resp = &ResponseTransactionJson{
			TransactionID:    dataTransaction.GetTransactionID(),
			CustomerID:       dataTransaction.GetCustomerID(),
			Revenue:          dataTransaction.GetRevenue(),
			CouponID:         dataTransaction.GetCouponID(),
			PurchaseDate:     dataTransaction.GetPurchaseDate(),
			TransactionItems: listItems,
		}
	}

	httpResponse := &CustomReponseSingle{
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
