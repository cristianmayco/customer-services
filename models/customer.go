package models

import (
	"context"
	"cristianmayco/customer-services/infra"
	"fmt"
	"time"
)

type Customer struct {
	ID               int       `json:"id" db:"id" validate:"required"`
	Name             string    `json:"name" db:"name" validate:"required"`
	CustomerIdentity string    `json:"customer_identity" db:"customer_identity" validate:"required"`
	Status           string    `json:"status" db:"status" validate:"required"`
	CreateAt         time.Time `json:"create_at" db:"create_at" validate:"required"`
	UpdateAt         time.Time `json:"update_at" db:"update_at" validate:"required"`
}

func NewCustomer(id int, name, customerIdentity, status string) *Customer {
	currentTime := time.Now() // Use time.Time diretamente
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

func (c *Customer) Save() error {
	query := `
        INSERT INTO "customers" (name, customer_identity, status, create_at, update_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	fmt.Printf("Saving customer: %+v\n", c) // Log para depuração
	err := infra.DB.QueryRow(context.Background(), query, c.Name, c.CustomerIdentity, c.Status, c.CreateAt, c.UpdateAt).Scan(&c.ID)
	if err != nil {
		fmt.Printf("Error saving customer: %v\n", err) // Log do erro
	}
	return err
}

func GetCustomerByID(id int) (*Customer, error) {
	query := `SELECT id, name, customer_identity, status, create_at, update_at FROM customers WHERE id = $1`
	row := infra.DB.QueryRow(context.Background(), query, id)

	customer := &Customer{}
	err := row.Scan(&customer.ID, &customer.Name, &customer.CustomerIdentity, &customer.Status, &customer.CreateAt, &customer.UpdateAt)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func GetAllCustomers() ([]Customer, error) {
	query := `SELECT * FROM customers`
	fmt.Println("Executing query:", query) // Log para depuração
	rows, err := infra.DB.Query(context.Background(), query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err) // Log do erro
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.CustomerIdentity, &customer.Status, &customer.CreateAt, &customer.UpdateAt); err != nil {
			fmt.Printf("Error scanning row: %v\n", err) // Log do erro
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *Customer) Update() error {
	query := `
        UPDATE customers
        SET name = $1, customer_identity = $2, status = $3, update_at = $4
        WHERE id = $5
    `
	_, err := infra.DB.Exec(context.Background(), query, c.Name, c.CustomerIdentity, c.Status, c.UpdateAt, c.ID)
	return err
}

func DeleteCustomer(id int) error {
	query := `DELETE FROM customers WHERE id = $1`
	_, err := infra.DB.Exec(context.Background(), query, id)
	return err
}
