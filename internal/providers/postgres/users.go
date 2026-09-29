package postgres

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}


func NewUserRepo(db *gorm.DB) repository.UserRepo{
	return &userRepo{
		db:db,
	}
} 

func (r *userRepo)CreateUser(user *models.User) error{
	return r.db.Create(&user).Error	
}

func (r *userRepo) Get_User_By_Email(email string)(user models.User,err error){
	err=r.db.Where("email=?",email).First(&user).Error
	return user,err
}


