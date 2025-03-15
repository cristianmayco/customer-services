package models

type Service struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

func NewService(id int, name, description, status string) *Service {
	return &Service{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
	}
}
