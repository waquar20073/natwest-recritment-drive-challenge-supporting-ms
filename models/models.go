package models

import "gorm.io/gorm"

// Product represents the structure of an inventory product
type Product struct {
	gorm.Model
	Id       int     `gorm:"type:char(36);primaryKey" json:"id"` // UUID as product ID
	Name     string  `json:"name"`                               // Product name
	Quantity int     `json:"quantity"`                           // Available quantity
	Price    float64 `json:"price"`                              // Price of the product
	IsActive bool    `json:"isActive"`                           // Product availability status
}
