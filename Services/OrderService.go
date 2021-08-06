package Services

import (
	"fmt"

	"mygra.tech/project1/Models"
	"mygra.tech/project1/Repositories"
)

type OrderService interface {
	GetOrders(pagination *Models.Pagination) ([]Models.Order, error)
	GetAOrder(id string) (Models.Order, error)
	CreateAOrder(order Models.Order) (Models.Order, error)
	UpdateAOrder(order Models.Order, id string) (Models.Order, error)
	DeleteAOrder(order Models.Order, id string) error
}

type orderService struct {
	repository        Repositories.OrderRepository
	productRepository Repositories.ProductRepository
}

func InitOrderService(repository Repositories.OrderRepository, productRepository Repositories.ProductRepository) *orderService {
	return &orderService{repository, productRepository}
}

func (service *orderService) GetOrders(pagination *Models.Pagination) ([]Models.Order, error) {
	result, err := service.repository.GetOrders(pagination)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *orderService) GetAOrder(id string) (Models.Order, error) {
	result, err := service.repository.GetAOrder(id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *orderService) CreateAOrder(order Models.Order) (Models.Order, error) {
	result, err := service.repository.CreateAOrder(order)

	if err != nil {
		return result, err
	}
	resultProductAmount, errReductProductAmount := service.productRepository.ReduceAmount(fmt.Sprint(order.ProductID))

	if errReductProductAmount != nil {
		fmt.Println(resultProductAmount)
		return result, errReductProductAmount
	}

	return result, nil
}

func (service *orderService) UpdateAOrder(order Models.Order, id string) (Models.Order, error) {
	result, err := service.repository.UpdateAOrder(order, id)

	fmt.Println(order)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *orderService) DeleteAOrder(order Models.Order, id string) error {
	err := service.repository.DeleteAOrder(order, id)

	if err != nil {
		return err
	}
	return nil
}
