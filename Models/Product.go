package Models

import "time"

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Amount      uint64    `json:"amount"`
	ProductType string    `json:"productType" default:"BASIC"`
	Status      string    `json:"status" default:"ACTIVE"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (b *Product) TableName() string {
	return "product"
}
