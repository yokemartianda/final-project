package transaction_handler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"fmt"
	"net/http"
)

func (tr *TransactionHandler) StoreDataTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestTransaction
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)

	if errDecode != nil {
		fmt.Println(errDecode)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		TransactionID: 0,
		CustomerID:    "",
		ProductID:     0,
		Quantity:      0,
		Revenue:       0,
		CouponID:      "",
		PurchaseDate:  "",
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}
	errInsert := tr.repoTransaction.InsertDataTransaction(tr.ctx, transaction)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}
	w.WriteHeader(200)
	fmt.Fprint(w, "SUCCES INSERT DATA")
}
