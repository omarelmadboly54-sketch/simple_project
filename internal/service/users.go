package service

import (
	"errors"
	"fmt"
	"simple_project/internal/models"
	"simple_project/internal/repository"
	"simple_project/internal/util"
	"simple_project/token"
	"time"
)



type userService struct {
	UserRepo 		repository.UserRepo
	tokenMaker		token.Maker
	tokenDuration	time.Duration
}

type UserService interface{
	CreateUser(user models.User)error
	Login(email string,password string)(string,error)
}


func NewUserService(UserRepo repository.UserRepo,tokenMaker token.Maker,tokenDuration time.Duration)UserService{
	return &userService{
		UserRepo:UserRepo,
		tokenMaker:tokenMaker,
		tokenDuration:tokenDuration,
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

func(s *userService)Login(email string,password string)(string,error){
	//we need first to check if he/she is a user or not
	user,err:=s.UserRepo.Get_User_By_Email(email)
	if err!=nil{
		return "",fmt.Errorf("invalid credentials")
	}

	//so if he/she is a user we need to check password 
	err=util.VerfiyPassword(user.Password,password)
	if err!=nil{
		return "",fmt.Errorf("invalid credentials")
	}
	// after checking the user,email we need to create tokens

	accessToken,err:=s.tokenMaker.CreateToken(user.Email,user.Role,s.tokenDuration)
	if err!=nil{
		return "",fmt.Errorf("failed to generate tokens")
	}
	return accessToken,nil
}



