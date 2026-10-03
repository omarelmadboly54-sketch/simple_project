package service

import (
	"errors"
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
		tokenMaker: tokenMaker,
		tokenDuration: tokenDuration,
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

func(s *userService) Login(email string,password string)(string,error){
	user,err:=s.UserRepo.Get_User_By_Email(email)
	if err!=nil{
		return "",errors.New("invalid email or password")
	}
	err = util.VerfiyPassword(user.Password,password)
	if err!=nil{
		return "",errors.New("invalid email or password")
	}
	accessToken,err:=s.tokenMaker.CreateToken(user.Email,user.Role,s.tokenDuration)
	if err!=nil{
		return "",errors.New("failed to generate token")
	}
	return accessToken,nil
}

