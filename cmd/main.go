package main

import (
	"cristianmayco/customer-services/handlers"
	"cristianmayco/customer-services/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)

	// Rotas para o Customer
	customerGroup := r.Group("/customers")
	customerGroup.Use(middleware.JWTMiddleware())
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
