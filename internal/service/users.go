package service

import (
	"fmt"
	"simple_project/internal/apperrors"
	"simple_project/internal/models"
	"simple_project/internal/repository"
	"simple_project/internal/util"
	"simple_project/token"
	"time"
)



type userService struct {
	UserRepo 		repository.UserRepository
	tokenMaker		token.Maker
	tokenDuration	time.Duration
}

type UserService interface{
	CreateUser(user *models.User)error
	Login(email string,password string)(string,error)
	GetAll()(users []models.User,err error)
	GetUser(id int64)(user *models.User,err error)
	DeleteUser(id int64)error
	UpdateUserRole(id int64)(user *models.User ,err error)
}


func NewUserService(UserRepo repository.UserRepository,tokenMaker token.Maker,tokenDuration time.Duration)UserService{
	return &userService{
		UserRepo:UserRepo,
		tokenMaker:tokenMaker,
		tokenDuration:tokenDuration,
	}
}


func(s *userService) CreateUser(user *models.User)error{
	_,err:=s.UserRepo.GetUserByEmail(user.Email)
	if err ==nil{
		return apperrors.ErrUserAlreadyExists	
	}
	hashedpassword,err:=util.HashPassword(user.Password)
	if err!=nil{
		return err
	}

	user.Password=hashedpassword
	return s.UserRepo.CreateUser(user)
}

func(s *userService)GetAll()(users []models.User,err error){
	//there is no logic here so we will return the users 
	return s.UserRepo.GetAll()
}

func(s *userService)GetUser(id int64)(user *models.User,err error){
	return s.UserRepo.GetUser(id)
}

func(s *userService)DeleteUser(id int64)error{
 	user,err:=s.UserRepo.GetUser(id)
	if err!=nil{
		return err
	}
	if err:=s.UserRepo.DeleteUser(user.ID);err!=nil{
		return err
	}
	return nil	
 }

 func(s *userService)UpdateUserRole(id int64)(user *models.User ,err error){
	// i nned to check if he/she is a user or not
	user,err=s.UserRepo.GetUser(id)
	if err!=nil{
		return nil,err
	}
	return s.UserRepo.UpdateUserRole(user.ID)
 }



func(s *userService)Login(email string,password string)(string,error){
	//we need first to check if he/she is a user or not
	user,err:=s.UserRepo.GetUserByEmail(email)
	if err!=nil{
		return "",apperrors.ErrInvalidCredentials
	}

	//so if he/she is a user we need to check password 
	err=util.VerifyPassword(user.Password,password)
	if err!=nil{
		return "",apperrors.ErrInvalidCredentials
	}
	// after checking the user,email we need to create tokens

	accessToken,err:=s.tokenMaker.CreateToken(user.ID,user.Email,user.Role,s.tokenDuration)
	if err!=nil{
		return "",fmt.Errorf("failed to generate tokens")
	}
	return accessToken,nil
}



