package Repositories

import (
	"fmt"

	"gorm.io/gorm"
	"mygra.tech/project1/Models"
)

type UserRepository interface {
	GetUsers(pagination *Models.Pagination) ([]Models.User, error)
	GetAUser(id string) (Models.User, error)
	CreateAUser(user Models.User) (Models.User, error)
	UpdateAUser(user Models.User, id string) (Models.User, error)
	DeleteAUser(user Models.User, id string) error
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repository *userRepository) GetUsers(pagination *Models.Pagination) ([]Models.User, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	users := []Models.User{}

	err := repository.db.Preload("Todo").Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func (repository *userRepository) GetAUser(id string) (Models.User, error) {
	user := Models.User{}

	err := repository.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) CreateAUser(user Models.User) (Models.User, error) {
	tx := repository.db.Begin()

	err := tx.Create(&user).Error

	if err != nil {
		fmt.Println("Error: ", err)
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, nil
}

func (repository *userRepository) UpdateAUser(userInput Models.User, id string) (Models.User, error) {
	var user Models.User

	repository.db.Where("id = ?", id).Find(&user)
	user.Username = userInput.Username
	user.Email = userInput.Email
	user.PhoneNumber = userInput.PhoneNumber

	err := repository.db.Save(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *userRepository) DeleteAUser(user Models.User, id string) error {
	err := repository.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occurred")
		return err
	}
	return nil
}
