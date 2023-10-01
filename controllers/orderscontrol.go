package controllers

import (
	"Assignment-2/controllers/responsejson"
	"Assignment-2/database"
	"Assignment-2/models"
	"Assignment-2/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var request models.Orders
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := models.Orders{
		CustomerName: request.CustomerName,
		OrderedAt:    request.OrderedAt,
	}

	for _, itemData := range request.Item {
		newItem := models.Items{
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
		}
		newOrder.Item = append(newOrder.Item, newItem)
	}

	if err := repository.CreateOrder(&newOrder, database.GetDB()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the order"})
		return
	}

	orderResponse := responsejson.OrderResponse{
		ID:           newOrder.ID,
		CreatedAt:    newOrder.CreatedAt,
		CustomerName: newOrder.CustomerName,
		OrderedAt:    newOrder.OrderedAt,
		Items:        newOrder.Item,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": orderResponse})
}

func GetAllOrders(c *gin.Context) {
	db := database.GetDB()
	// Check if db is a valid database connection
	if database.GetDB() == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not available"})
		return
	}

	// Fetch all orders using the repository function
	orders, err := repository.GetAllOrders(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
	// Get the order ID from the request URL
	orderIDStr := c.Param("orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Fetch the existing order from the database
	existingOrder, err := repository.GetOrderByID(orderID, database.GetDB())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the order"})
		return
	}

	// Bind the request JSON to an updatedOrder struct
	var updatedOrder models.Orders
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the order fields as needed
	existingOrder.CustomerName = updatedOrder.CustomerName
	existingOrder.OrderedAt = updatedOrder.OrderedAt

	// Save the updated order to the database
	if err := repository.UpdateOrder(existingOrder, database.GetDB()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the order"})
		return
	}

	c.JSON(http.StatusOK, existingOrder)
}

func UpdateItem(c *gin.Context) {
	// Get the item ID from the request URL
	itemIDStr := c.Param("Id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Fetch the existing item from the database
	existingItem, err := repository.GetItemByID(itemID, database.GetDB())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the item"})
		return
	}

	// Bind the request JSON to an updatedItem struct
	var updatedItem models.Items
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the item fields as needed
	existingItem.ItemCode = updatedItem.ItemCode
	existingItem.Description = updatedItem.Description
	existingItem.Quantity = updatedItem.Quantity

	// Save the updated item to the database
	if err := repository.UpdateItem(itemID, &existingItem, database.GetDB()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the item"})
		return
	}

	c.JSON(http.StatusOK, existingItem)
}

func DeleteOrder(c *gin.Context) {
	// Get the order ID from the request URL
	orderIDStr := c.Param("orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Delete the order from the database using the database connection
	if err := repository.DeleteOrder(orderID, database.GetDB()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
