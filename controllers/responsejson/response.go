package responsejson

import (
	"Assignment-2/models"
	"time"
)

type OrderResponse struct {
	ID           uint           `json:"ID"`
	CreatedAt    time.Time      `json:"CreatedAt"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    time.Time      `json:"ordered_at"`
	Items        []models.Items `json:"Items"`
}
