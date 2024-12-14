package controllers

import "database/sql"

type UserController interface {
	CreateUser()
}

type userController struct {
	db *sql.DB
}

// our constructor for our UserController
func NewUserController(db *sql.DB) UserController {
	return &userController{
		db: db,
	}
}

func (c *userController) CreateUser() {

}
