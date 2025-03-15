package models

import "time"

type CustomerService struct {
	ID           int    `json:"id" validate:"required"`
	CustomerID   int    `json:"customer_id" validate:"required"`
	ServiceID    int    `json:"service_id" validate:"required"`
	Status       string `json:"status" validate:"required"`
	CreateAt     string `json:"create_at" validate:"required"`
	UpdateAt     string `json:"update_at" validate:"required"`
	TokenService string `json:"token_service" validate:"required"`
}

func NewCustomerService(id, customerID, serviceID int, status, tokenService string) *CustomerService {
	return &CustomerService{
		ID:           id,
		CustomerID:   customerID,
		ServiceID:    serviceID,
		Status:       status,
		CreateAt:     time.Now().UTC().Format(time.RFC3339),
		UpdateAt:     time.Now().UTC().Format(time.RFC3339),
		TokenService: tokenService,
	}
}
