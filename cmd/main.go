package main

import (
	"context"
	databasesql "final-project/internal/config/database/mysql"
	"final-project/internal/delivery/http/coupon_handler"
	"final-project/internal/delivery/http/customer_hendler"
	"final-project/internal/delivery/http/transaction_handler"
	"final-project/internal/repository/mysql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	connectionDatabase              = databasesql.InitMysqlDB()
	customerRepositoryMysql         = mysql.NewCustomerMysql(connectionDatabase)
	transactionRepositoryMysql      = mysql.NewTransactionMysql(connectionDatabase)
	couponRepositoryMysql           = mysql.NewCouponMysql(connectionDatabase)
	transactionItemsRepositoryMysql = mysql.NewTransactionItemsMysql(connectionDatabase)
	ctx                             = context.Background()
)

func main() {
	r := mux.NewRouter()

	handlerCustomer := customer_hendler.NewCustomerHandler(ctx, customerRepositoryMysql)
	handlerTransaction := transaction_handler.NewTransactionHandler(ctx, transactionRepositoryMysql, transactionItemsRepositoryMysql)
	handlerCoupon := coupon_handler.NewCouponHandler(ctx, couponRepositoryMysql, customerRepositoryMysql)
	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/create-customer", handlerCustomer.StoreDataCustomer).Methods(http.MethodPost)
	r.HandleFunc("/create-transaction", handlerTransaction.StoreDataTransaction).Methods(http.MethodPost)
	r.HandleFunc("/create-coupon", handlerCoupon.StoreDataCoupon).Methods(http.MethodPost)
	r.HandleFunc("/getdisccoupon/{idcustomer}", handlerCoupon.GetCouponByIdCustomer).Methods(http.MethodGet)
	http.HandleFunc("/test", ParamHandlerWithoutInput)
	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCES OK")
}
