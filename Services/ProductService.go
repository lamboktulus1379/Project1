package Services

import (
	"fmt"

	"gorm.io/gorm"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Repositories"
)

type ProductService interface {
	GetProducts(pagination *Models.Pagination) ([]Models.Product, error)
	GetAProduct(id string) (Models.Product, error)
	CreateAProduct(product Models.Product) (Models.Product, error)
	UpdateAProduct(product Models.Product, id string) (Models.Product, error)
	DeleteAProduct(product Models.Product, id string) error
	ReduceAmount(id string) (Models.Product, error)
	WithTrx(trxHandle *gorm.DB) *productService
}

type productService struct {
	repository Repositories.ProductRepository
}

func InitProductService(repository Repositories.ProductRepository) *productService {
	return &productService{repository}
}

func (service *productService) WithTrx(trxHandle *gorm.DB) *productService {
	service.repository = service.repository.WithTrx(trxHandle)
	return service
}

func (service *productService) GetProducts(pagination *Models.Pagination) ([]Models.Product, error) {
	result, err := service.repository.GetProducts(pagination)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *productService) GetAProduct(id string) (Models.Product, error) {
	result, err := service.repository.GetAProduct(id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *productService) CreateAProduct(product Models.Product) (Models.Product, error) {
	result, err := service.repository.CreateAProduct(product)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *productService) UpdateAProduct(product Models.Product, id string) (Models.Product, error) {
	result, err := service.repository.UpdateAProduct(product, id)

	fmt.Println(product)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *productService) DeleteAProduct(product Models.Product, id string) error {
	err := service.repository.DeleteAProduct(product, id)

	if err != nil {
		return err
	}
	return nil
}

func (service *productService) ReduceAmount(id string) (Models.Product, error) {
	result, err := service.repository.ReduceAmount(id)

	if err != nil {
		return result, err
	}
	return result, nil
}
