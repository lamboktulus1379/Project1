package Repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Services/Products"
)

type ProductRepository interface {
	GetProducts(pagination *Models.Pagination) ([]Models.Product, error)
	GetAProduct(id string) (Models.Product, error)
	CreateAProduct(product Models.Product) (Models.Product, error)
	UpdateAProduct(product Models.Product, id string) (Models.Product, error)
	DeleteAProduct(product Models.Product, id string) error
	ReduceAmount(id string) (Models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (repository *productRepository) GetProducts(pagination *Models.Pagination) ([]Models.Product, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	products := []Models.Product{}

	err := repository.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (repository *productRepository) GetAProduct(id string) (Models.Product, error) {
	product := Models.Product{}

	err := repository.db.Where("id = ?", id).First(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (repository *productRepository) CreateAProduct(product Models.Product) (Models.Product, error) {
	tx := repository.db.Begin()

	product.Status = Models.ACTIVE
	err := tx.Create(&product).Error

	if err != nil {
		fmt.Println("Error: ", err)
		tx.Rollback()
		return product, err
	}
	tx.Commit()

	return product, nil
}

func (repository *productRepository) UpdateAProduct(productInput Models.Product, id string) (Models.Product, error) {
	var product Models.Product

	productType := productInput.ProductType

	switch productType {
	case "BASIC":
		productType = Products.BASIC
	case "STANDARD":
		productType = Products.STANDARD
	case "PREMIUM":
		productType = Products.PREMIUM
	default:
		return product, errors.New("Product type not found!")
	}

	repository.db.Where("id = ?", id).Find(&product)

	product.Name = productInput.Name
	product.Amount = productInput.Amount
	product.Price = productInput.Price
	product.ProductType = productInput.ProductType
	product.Status = productInput.Status

	err := repository.db.Save(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *productRepository) DeleteAProduct(product Models.Product, id string) error {
	err := repository.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occurred")
		return err
	}
	return nil
}

func (repository *productRepository) ReduceAmount(id string) (Models.Product, error) {
	var product Models.Product

	repository.db.Where("id = ?", id).Find(&product)

	product.Amount = product.Amount - 1
	result, err := repository.UpdateAProduct(product, id)

	if err != nil {
		return result, err
	}
	return result, nil
}
