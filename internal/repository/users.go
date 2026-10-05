package repository

import "simple_project/internal/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	Get_User_By_Email(email string)(user models.User,err error)
	Get_Users()(users []models.User,err error)
	Delete_user(id int64)error
	Get_User(id int64)(user models.User,err error)
	
}

