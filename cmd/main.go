package main

import (
	"cristianmayco/customer-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rotas para o Customer
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("/", handlers.GetAllCustomers)
		customerGroup.GET("/:id", handlers.GetCustomerByID)
		customerGroup.POST("/", handlers.CreateCustomer)
		customerGroup.PUT("/:id", handlers.UpdateCustomer)
		customerGroup.DELETE("/:id", handlers.DeleteCustomer)
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
