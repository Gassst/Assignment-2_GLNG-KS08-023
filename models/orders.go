package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	CustomerName string    `gorm:"not null;type:varchar(190)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	// CreatedAt    time.Time `json:"created_at"`
	// UpdatedAt    time.Time `json:"updated_at"`
	Item []Items `gorm:"foreignKey:OrderID" json:"Items"`
}
