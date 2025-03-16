package handlers

import (
	"cristianmayco/customer-services/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /customers - Retorna todos os clientes
func GetAllCustomers(c *gin.Context) {
	customers, err := models.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// GET /customers/:id - Retorna um cliente pelo ID
func GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	customer, err := models.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// POST /customers - Cria um novo cliente
func CreateCustomer(c *gin.Context) {
	var newCustomer models.Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Preenche os campos de data
	currentTime := time.Now()
	newCustomer.CreateAt = currentTime
	newCustomer.UpdateAt = currentTime

	if err := newCustomer.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save customer"})
		return
	}

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

	updatedCustomer.ID = id
	updatedCustomer.UpdateAt = time.Now()

	if err := updatedCustomer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, updatedCustomer)
}

// DELETE /customers/:id - Remove um cliente pelo ID
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteCustomer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
