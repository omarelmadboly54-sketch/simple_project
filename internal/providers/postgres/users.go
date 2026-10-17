package postgres

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}


func NewUserRepo(db *gorm.DB) repository.UserRepository{
	return &userRepo{
		db:db,
	}
} 

func (r *userRepo)CreateUser(user *models.User) error{
	return r.db.Create(user).Error	
}

func (r *userRepo) GetUserByEmail(email string)(user models.User,err error){
	err=r.db.Where("email=?",email).First(&user).Error
	return user,err
}

func (r *userRepo) GetAll()(users []models.User,err error){
	if err:=r.db.Find(&users).Error;err!=nil{
		return nil,err
	}
	return users,nil
}

func (r *userRepo) GetUser(id int64)(user*models.User,err error){
	err = r.db.Where("id=?",id).First(&user).Error
	return user,err
}


 func (r *userRepo)DeleteUser(id int64)error{
 	result:=r.db.Delete(&models.User{},id)
 	if result.Error!=nil{
 		return result.Error
 	}
 	return nil
 }

  func (r *userRepo)UpdateUserRole(id int64)(user *models.User,err error){
	// var user models.User
	result:=r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?",id).Update("role","admin")
	if result.Error !=nil{
		return nil,result.Error
	}
	return user,nil
  }












