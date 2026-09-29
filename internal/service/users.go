package service

import (
	"errors"
	"simple_project/internal/models"
	"simple_project/internal/repository"
	"simple_project/internal/util"
)



type userService struct {
	UserRepo repository.UserRepo
}

type UserService interface{
	CreateUser(user models.User)error
}


func NewUserService(UserRepo repository.UserRepo)UserService{
	return &userService{
		UserRepo:UserRepo,
	}
}

func(s *userService) CreateUser(user models.User)error{
	_,err:=s.UserRepo.Get_User_By_Email(user.Email)
	if err ==nil{
		return errors.New("user already exist")	
	}
	hashedpassword,err:=util.HashPassword(user.Password)
	if err!=nil{
		return err
	}

	user.Password=hashedpassword
	return s.UserRepo.CreateUser(&user)

}



