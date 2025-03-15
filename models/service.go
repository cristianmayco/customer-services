package models

import (
	"time"
)

type Service struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
	CreateAt    string `json:"create_at" validate:"required"`
	UpdateAt    string `json:"update_at" validate:"required"`
}

func NewService(id int, name, description, status string) *Service {
	return &Service{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
		CreateAt:    time.Now().UTC().Format(time.RFC3339),
		UpdateAt:    time.Now().UTC().Format(time.RFC3339),
	}
}
