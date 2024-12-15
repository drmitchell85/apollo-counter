package controllers

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/repository"
)

type UserController interface {
	CreateUser(req models.NewUserRequest) error
	GetUserByEmail(email string) (models.GetUserByEmailResponse, error)
}

type userController struct {
	userRepo repository.UserRepository
}

// our constructor for our UserController
func NewUserController(repo repository.UserRepository) UserController {
	return &userController{
		userRepo: repo,
	}
}

func (c *userController) CreateUser(req models.NewUserRequest) error {
	user := models.User{
		Username:    req.Username,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Phonenumber: req.Phonenumber,
		Email:       req.Email,
		Password:    req.Password,
	}

	err := c.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) GetUserByEmail(email string) (models.GetUserByEmailResponse, error) {

	user, err := c.userRepo.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}
