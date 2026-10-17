package repository

import "simple_project/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string)(user models.User,err error)
	GetAll()(users []models.User,err error)
	DeleteUser(id int64)error
	GetUser(id int64)(user *models.User,err error)
	UpdateUserRole(id int64)(*models.User , error)
}



