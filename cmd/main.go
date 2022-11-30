package main

import (
	"context"
	"final-project/domain/entity"
	"final-project/domain/repository"
	databasesql "final-project/internal/config/database/mysql"
	"final-project/internal/repository/mysql"
	"fmt"
)

var (
	connectionDatabase      = databasesql.InitMysqlDB()
	customerRepositoryMysql = mysql.NewCustomerMysql(connectionDatabase)
)

type CustomerLogicFactoryHandler struct {
	customerRepository repository.InterfaceRepoCustomer
}

func NewCostumerLogicFactoryHandler(repoCustomerImplementation repository.InterfaceRepoCustomer) *CustomerLogicFactoryHandler {
	return &CustomerLogicFactoryHandler{customerRepository: repoCustomerImplementation}
}

func InsertDataCustomer(ctx context.Context) *entity.Customer {
	BuildFirstCustomer := entity.DTOCustomer{
		UserID:      1,
		Name:        "Divo",
		Alamat:      "Tanah Abang",
		PhoneNumber: "081234567890",
		CreatedTime: "2022-11-11",
	}

	FirstCustomer, errCheckDomainCustomer := entity.NewCustomer(BuildFirstCustomer)
	if errCheckDomainCustomer != nil {
		fmt.Println("GAGAL CREATE CUSTOMER KARENA WRONG DOMAIN")
		panic(errCheckDomainCustomer)
	}
	fmt.Println("--->Proces Store Data To DB")
	handlerRepo := NewCostumerLogicFactoryHandler(customerRepositoryMysql)
	errStoreRepo := handlerRepo.customerRepository.InsertDataCustomer(ctx, FirstCustomer)
	if errStoreRepo != nil {
		fmt.Println("GAGAL CREATE CUSTOMER ADA KESALAHAN DALAM PENYIMPANAN")
		panic(errStoreRepo)
	}
	return FirstCustomer
}

func main() {
	ctx := context.Background()
	InsertDataCustomer(ctx)
}
