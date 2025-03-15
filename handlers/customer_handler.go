package handlers

import (
	"cristianmayco/customer-services/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Simula um banco de dados em memória
var customers = []models.Customer{}

// GET /customers - Retorna todos os clientes
func GetAllCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)
}

// GET /customers/:id - Retorna um cliente pelo ID
func GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, customer := range customers {
		if customer.ID == id {
			c.JSON(http.StatusOK, customer)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

// POST /customers - Cria um novo cliente
func CreateCustomer(c *gin.Context) {
	var newCustomer models.Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer.ID = len(customers) + 1
	customers = append(customers, newCustomer)
	c.JSON(http.StatusCreated, newCustomer)
}

// PUT /customers/:id - Atualiza um cliente existente
func UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedCustomer models.Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			customers[i].ID = id // Garante que o ID não seja alterado
			c.JSON(http.StatusOK, customers[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

// DELETE /customers/:id - Remove um cliente pelo ID
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}
