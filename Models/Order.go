package Models

import (
	"time"
)

type Order struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"userId"`
	Price     float64    `json:"price"`
	ProductID uint       `json:"productId"`
	Status    string     `json:"status"`
	CreatedBy string     `json:"createdBy" default:"System"`
	UpdatedBy string     `json:"updatedBy" default:"nil"`
	CreatedAt time.Time  `json:"createdAt"`
	PaidAt    *time.Time `json:"paidAt" default:"nil"`
}

func (b *Order) TableName() string {
	return "order"
}
