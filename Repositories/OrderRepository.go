package Repositories

import (
	"fmt"

	"gorm.io/gorm"
	"mygra.tech/project1/Models"
)

type OrderRepository interface {
	GetOrders(pagination *Models.Pagination) ([]Models.Order, error)
	GetAOrder(id string) (Models.Order, error)
	CreateAOrder(order Models.Order) (Models.Order, error)
	UpdateAOrder(order Models.Order, id string) (Models.Order, error)
	DeleteAOrder(order Models.Order, id string) error
}

type orderRepository struct {
	db *gorm.DB
}

func InitOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (repository *orderRepository) GetOrders(pagination *Models.Pagination) ([]Models.Order, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	orders := []Models.Order{}

	err := repository.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&orders).Error

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (repository *orderRepository) GetAOrder(id string) (Models.Order, error) {
	order := Models.Order{}

	err := repository.db.Where("id = ?", id).First(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (repository *orderRepository) CreateAOrder(order Models.Order) (Models.Order, error) {
	tx := repository.db.Begin()

	err := tx.Create(&order).Error

	if err != nil {
		fmt.Println("Error: ", err)
		tx.Rollback()
		return order, err
	}
	tx.Commit()

	return order, nil
}

func (repository *orderRepository) UpdateAOrder(orderInput Models.Order, id string) (Models.Order, error) {
	var order Models.Order

	repository.db.Where("id = ?", id).Find(&order)
	order.Status = orderInput.Status
	order.Price = orderInput.Price
	order.ProductID = orderInput.ProductID

	err := repository.db.Save(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (repository *orderRepository) DeleteAOrder(order Models.Order, id string) error {
	err := repository.db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occurred")
		return err
	}
	return nil
}
