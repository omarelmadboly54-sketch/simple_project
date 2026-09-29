package repository

import "simple_project/internal/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	Get_User_By_Email(email string)(user models.User,err error)
}