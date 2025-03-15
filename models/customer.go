package models

type Customer struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	CustomerIdentity string `json:"customer_identity"`
	Status           string `json:"status"`
}

func NewCustomer(id int, name, customerIdentity, status string) *Customer {
	return &Customer{
		ID:               id,
		Name:             name,
		CustomerIdentity: customerIdentity,
		Status:           status,
	}
}
