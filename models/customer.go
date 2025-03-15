package models

import "time"

type Customer struct {
	ID               int    `json:"id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	CustomerIdentity string `json:"customer_identity" validate:"required"`
	Status           string `json:"status" validate:"required"`
	CreateAt         string `json:"create_at" validate:"required"`
	UpdateAt         string `json:"update_at" validate:"required"`
}

func NewCustomer(id int, name, customerIdentity, status string) *Customer {
	return &Customer{
		ID:               id,
		Name:             name,
		CustomerIdentity: customerIdentity,
		Status:           status,
		CreateAt:         time.Now().UTC().Format(time.RFC3339),
		UpdateAt:         time.Now().UTC().Format(time.RFC3339),
	}
}
