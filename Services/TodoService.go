package Services

import (
	"fmt"

	"mygra.tech/project1/Models"
	"mygra.tech/project1/Repositories"
)

type TodoService interface {
	GetTodos(pagination *Models.Pagination) ([]Models.Todo, error)
	GetATodo(id string) (Models.Todo, error)
	CreateATodo(todo Models.Todo) (Models.Todo, error)
	UpdateATodo(todo Models.Todo, id string) (Models.Todo, error)
	DeleteATodo(todo Models.Todo, id string) error
}

type todoService struct {
	repository Repositories.TodoRepository
}

func InitTodoService(repository Repositories.TodoRepository) *todoService {
	return &todoService{repository}
}

func (service *todoService) GetTodos(pagination *Models.Pagination) ([]Models.Todo, error) {
	result, err := service.repository.GetTodos(pagination)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *todoService) GetATodo(id string) (Models.Todo, error) {
	result, err := service.repository.GetATodo(id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *todoService) CreateATodo(todo Models.Todo) (Models.Todo, error) {
	result, err := service.repository.CreateATodo(todo)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *todoService) UpdateATodo(todo Models.Todo, id string) (Models.Todo, error) {
	result, err := service.repository.UpdateATodo(todo, id)

	fmt.Println(todo)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *todoService) DeleteATodo(todo Models.Todo, id string) error {
	err := service.repository.DeleteATodo(todo, id)

	if err != nil {
		return err
	}
	return nil
}
