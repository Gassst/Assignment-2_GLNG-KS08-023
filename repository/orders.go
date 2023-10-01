package repository

import (
	"Assignment-2/models"

	"gorm.io/gorm"
)

func GetAllOrders(db *gorm.DB) ([]models.Orders, error) {
	var orders []models.Orders
	result := db.Preload("Item").Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return orders, result.Error
		}
	}

}

func CreateOrder(input *models.Orders, db *gorm.DB) error {
	result := db.Debug().Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetOrderByID(id int, db *gorm.DB) (models.Orders, error) {
	var order models.Orders
	err := db.Preload("Item").First(&order, id).Error
	if err != nil {
		return models.Orders{}, err
	}
	return order, err
}

func UpdateOrder(order models.Orders, db *gorm.DB) error {
	result := db.Save(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetItemByID(id int, db *gorm.DB) (models.Items, error) {
	var item models.Items
	err := db.First(&item, id).Error
	if err != nil {
		return models.Items{}, err
	}
	return item, err
}

func UpdateItem(id int, updatedItem *models.Items, db *gorm.DB) error {
	// Fetch the existing item
	existingItem, err := GetItemByID(id, db)
	if err != nil {
		return err
	}

	// Update the item fields
	existingItem.ItemCode = updatedItem.ItemCode
	existingItem.Description = updatedItem.Description
	existingItem.Quantity = updatedItem.Quantity

	// Save the updated item
	result := db.Save(&existingItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Now you have GetItemByID to fetch an item by its ID and UpdateItem to update an item based on its ID. These functions are designed to be used in your UpdateItem controller for handling PATCH and PUT requests.

func DeleteOrder(id int, db *gorm.DB) error {
	var order models.Orders

	// Find the order by ID and delete it
	del := db.Delete(&order, id)

	if del.Error != nil {
		return del.Error
	}

	return nil
}
