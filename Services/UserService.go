package Services

import (
	"fmt"

	"mygra.tech/project1/Models"
	"mygra.tech/project1/Repositories"
)

type UserService interface {
	GetUsers(pagination *Models.Pagination) ([]Models.User, error)
	GetAUser(id string) (Models.User, error)
	CreateAUser(user Models.User) (Models.User, error)
	UpdateAUser(user Models.User, id string) (Models.User, error)
	DeleteAUser(user Models.User, id string) error
}

type userService struct {
	repository Repositories.UserRepository
}

func InitUserService(repository Repositories.UserRepository) *userService {
	return &userService{repository}
}

func (service *userService) GetUsers(pagination *Models.Pagination) ([]Models.User, error) {
	result, err := service.repository.GetUsers(pagination)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *userService) GetAUser(id string) (Models.User, error) {
	result, err := service.repository.GetAUser(id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *userService) CreateAUser(user Models.User) (Models.User, error) {
	result, err := service.repository.CreateAUser(user)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *userService) UpdateAUser(user Models.User, id string) (Models.User, error) {
	result, err := service.repository.UpdateAUser(user, id)

	fmt.Println(user)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *userService) DeleteAUser(user Models.User, id string) error {
	err := service.repository.DeleteAUser(user, id)

	if err != nil {
		return err
	}
	return nil
}
