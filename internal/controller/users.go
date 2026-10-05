package controller

import (
	"errors"
	"net/http"
	"simple_project/internal/middleware"
	"simple_project/internal/presenters"
	"simple_project/internal/service"
	"simple_project/token"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	UserService service.UserService
	tokenMaker	token.Maker
}

type UserController interface{
	Routes(g *gin.RouterGroup)
}

func NewUserController(UserService service.UserService,tokenMaker token.Maker)UserController{
	return &userController{
		UserService:UserService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *userController) Routes(g *gin.RouterGroup){
	g.POST("/users",ctl.CreateUser)
	g.POST("/login",ctl.Login)
	
	g.GET("/users",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Get_Users)
	g.GET("/users/:id",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Get_User)
	g.DELETE("/users/:id",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Delete_User)

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

func(ctl *userController)Get_Users(g *gin.Context){
	users,err:=ctl.UserService.Get_Users()
	if err !=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	g.JSON(http.StatusOK,users)

}

func(ctl *userController)Get_User(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"invalid user id",
	})
	return
	}
	user,err:=ctl.UserService.Get_User(id)
	if err!= nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":"user not found"})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	
	g.JSON(http.StatusOK,user)

}

func(ctl *userController)Delete_User(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"invalid user id"})
	return
	}
	err=ctl.UserService.Delete_User(id)
	if err!= nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":"user not found"})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}


	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}




func(ctl *userController)Login(g *gin.Context){
	req:=presenters.LoginRequest{}
	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
		return
	}
	access_token,err:=ctl.UserService.Login(req.Email,req.Password)
	if err!=nil{
		g.JSON(http.StatusUnauthorized,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{"access_token":access_token})
}




