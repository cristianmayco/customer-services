package models

import (
	"fmt"
	"time"
)

type Customer struct {
	ID               int    `json:"id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	CustomerIdentity string `json:"customer_identity" validate:"required"`
	Status           string `json:"status" validate:"required"`
	CreateAt         string `json:"create_at" validate:"required"`
	UpdateAt         string `json:"update_at" validate:"required"`
}

func NewCustomer(id int, name, customerIdentity, status string) *Customer {
	currentTime := time.Now().Format("2006-01-02 15:04:05") // Formata a data e hora
	customer := &Customer{
		ID:               id,
		Name:             name,
		CustomerIdentity: customerIdentity,
		Status:           status,
		CreateAt:         currentTime,
		UpdateAt:         currentTime,
	}
	// Log para depuração
	fmt.Printf("NewCustomer created: %+v\n", customer)
	return customer
}
