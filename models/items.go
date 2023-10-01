package models

import (
	// "time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	//ItemID      int    `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"not null;type:varchar(190)" json:"item_code"`
	Description string `gorm:"not null;type:varchar(190)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `gorm:"not null" json:"order_id" form:"order_id"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}
