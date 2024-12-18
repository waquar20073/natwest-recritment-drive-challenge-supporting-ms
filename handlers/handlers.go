package handlers

import (
	"inventory-ms/db"
	"inventory-ms/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Check Product Availability
func CheckAvailability(c *gin.Context) {
	productId := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, "id = ?", productId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	response := gin.H{
		"id":          product.Id,
		"quantity":    product.Quantity,
		"price":       product.Price,
		"isAvailable": product.Quantity > 0,
	}
	c.JSON(http.StatusOK, response)
}

// Handler to create an inventory record
func CreateInventoryHandler(c *gin.Context) {
	var newInventory models.Product
	if err := c.ShouldBindJSON(&newInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&newInventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory record"})
		return
	}
	c.JSON(http.StatusCreated, newInventory)
}

// Handler to get all inventory records with pagination, filtering, and sorting
func GetAllInventoryHandler(c *gin.Context) {
	var inventories []models.Product
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")
	minQuantity := c.DefaultQuery("min_quantity", "0")
	maxPrice := c.DefaultQuery("max_price", "999999")

	query := db.DB.Where("quantity >= ? AND price <= ?", minQuantity, maxPrice)
	offset := (page - 1) * limit

	if err := query.Order(sort + " " + order).Offset(offset).Limit(limit).Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory records"})
		return
	}
	c.JSON(http.StatusOK, inventories)
}

// Handler to get inventory record by product ID
func GetInventoryByIDHandler(c *gin.Context) {
	productID := c.Param("id")
	var inventory models.Product
	if err := db.DB.First(&inventory, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory record not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// Handler to update an inventory record
func UpdateInventoryHandler(c *gin.Context) {
	productID := c.Param("id")
	var inventory models.Product
	if err := db.DB.First(&inventory, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory record not found"})
		return
	}

	var updateData models.Product
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&inventory).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory record"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// Handler to delete an inventory record
func DeleteInventoryHandler(c *gin.Context) {
	productID := c.Param("id")
	if err := db.DB.Delete(&models.Product{}, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Inventory record deleted successfully"})
}

// Handler to perform bulk deduction of inventory
func BulkDeductInventoryHandler(c *gin.Context) {
	var request struct {
		Deductions []struct {
			ProductID int `json:"id"`
			Quantity  int `json:"quantity"`
		} `json:"deductions"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, deduction := range request.Deductions {
		var inventory models.Product
		if err := db.DB.First(&inventory, "id = ?", deduction.ProductID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product ID " + strconv.Itoa(deduction.ProductID) + " not found"})
			return
		}

		if inventory.Quantity < deduction.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient inventory for Product ID " + strconv.Itoa(deduction.ProductID)})
			return
		}

		inventory.Quantity -= deduction.Quantity
		if err := db.DB.Save(&inventory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory for Product ID " + strconv.Itoa(deduction.ProductID)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory deductions completed successfully"})
}
