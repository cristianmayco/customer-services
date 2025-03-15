package main

import (
	"fmt"

	"github.com/cristianmayco/customer-service/models"
)

func main() {

	service := models.NewService(1, "Service 1", "Service 1 Description", "Active")
	customer := models.NewCustomer(1, "Customer 1", "Customer 1 Identity", "Active")
	customerService := models.NewCustomerService(1, 1, 1, "Active", "2021-01-01", "2021-01-01", "Token Service")

	fmt.Println(service)
	fmt.Println(customer)
	fmt.Println(customerService)

}
