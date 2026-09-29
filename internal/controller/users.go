package controller

import (
	"net/http"
	"simple_project/internal/presenters"
	"simple_project/internal/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserService service.UserService
}

type UserController interface{
	Routes(g *gin.RouterGroup)
}

func NewUserController(UserService service.UserService)UserController{
	return &userController{
		UserService:UserService,
	}
}

func(ctl *userController) Routes(g *gin.RouterGroup){
	g.POST("/users",ctl.CreateUser)

}

func(ctl *userController)CreateUser(g *gin.Context){
	req:= presenters.CreateUserRequest{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"failed to bind request",
	})
	return
	}
	toUser:=req.ToUser()

	err:=ctl.UserService.CreateUser(toUser)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusCreated,gin.H{"ok":"ok",
	})
}

