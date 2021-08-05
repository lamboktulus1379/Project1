package Repositories

import (
	"fmt"

	"gorm.io/gorm"
	"mygra.tech/project1/Models"
)

type TodoRepository interface {
	GetTodos() ([]Models.Todo, error)
	GetATodo(id string) (Models.Todo, error)
	CreateATodo(todo Models.Todo) (Models.Todo, error)
	UpdateATodo(todo Models.Todo, id string) (Models.Todo, error)
	DeleteATodo(todo Models.Todo, id string) error
}

type todoRepository struct {
	db *gorm.DB
}

func InitTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (repository *todoRepository) GetTodos() ([]Models.Todo, error) {
	todos := []Models.Todo{}

	err := repository.db.Find(&todos).Error

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (repository *todoRepository) GetATodo(id string) (Models.Todo, error) {
	todo := Models.Todo{}

	err := repository.db.Where("id = ?", id).First(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository *todoRepository) CreateATodo(todo Models.Todo) (Models.Todo, error) {
	tx := repository.db.Begin()

	err := tx.Create(&todo).Error

	if err != nil {
		fmt.Println("Error: ", err)
		tx.Rollback()
		return todo, err
	}
	tx.Commit()

	return todo, nil
}

func (repository *todoRepository) UpdateATodo(todoInput Models.Todo, id string) (Models.Todo, error) {
	var todo Models.Todo

	repository.db.Where("id = ?", id).Find(&todo)
	todo.Title = todoInput.Title
	todo.Description = todoInput.Description

	err := repository.db.Save(&todo).Error

	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (repository *todoRepository) DeleteATodo(todo Models.Todo, id string) error {
	err := repository.db.Where("id = ?", id).Delete(&todo).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occurred")
		return err
	}
	return nil
}
