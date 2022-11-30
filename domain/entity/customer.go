package entity

import (
	"errors"
	"time"
)

type Customer struct {
	userID      int
	name        string
	alamat      string
	phoneNumber string
	createTime  time.Time
}

type DTOCustomer struct {
	UserID      int
	Name        string
	Alamat      string
	PhoneNumber string
	CreatedTime string
}

func NewCustomer(dto DTOCustomer) (*Customer, error) {
	if dto.Name == "" {
		return nil, errors.New("NAMA CANNOT BE EMPTY")
	}
	if dto.Alamat == "" {
		return nil, errors.New("ALAMAT CANNOT BE EMPTY")
	}
	if dto.PhoneNumber == "" {
		return nil, errors.New("PHONE NUMBER CANNOT BE EMPTY")
	}
	if dto.CreatedTime == "" {
		return nil, errors.New("CREATED TIME CANNOT BE EMPTY")
	}
	strCreatedTime, _ := time.Parse("2006-01-02", dto.CreatedTime)

	customer := &Customer{
		userID:      dto.UserID,
		name:        dto.Name,
		alamat:      dto.Alamat,
		phoneNumber: dto.PhoneNumber,
		createTime:  strCreatedTime,
	}
	return customer, nil
}

func (c *Customer) GetUserID() int {
	return c.userID
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetAlamat() string {
	return c.alamat
}

func (c *Customer) GetPhoneNumber() string {
	return c.phoneNumber
}

func (c *Customer) GetCreatedTime() string {
	return c.createTime.Format("2006-01-02")
}
