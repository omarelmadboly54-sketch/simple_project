package main

import (
	"log"
	"simple_project/configuration"
	"simple_project/internal/controller"
	"simple_project/internal/providers/postgres"
	"simple_project/internal/service"
	"simple_project/token"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	db,config, err := configuration.InitPostgres()
	if err!=nil{
		log.Fatal("failed to connect to DataBase",err)
	}
	log.Println("Database Connection Successfully")

	tokenMaker,err:=token.NewJwtMaker(config.SecretKey)
	if err!=nil{
		log.Fatal("cannot create token maker:",err)
	}
	


	r:=gin.Default()
	
	userRepo:=postgres.NewUserRepo(db)
	tokenDuration:=time.Duration(config.AccessTokenMinutes)*time.Minute
	userService:=service.NewUserService(userRepo,tokenMaker,tokenDuration)
	userController:=controller.NewUserController(userService)
	userController.Routes(r.Group("/api"))


	
	if err:=r.Run(config.ServerAddress);err!=nil{
		log.Fatalf("failed to start server:%v",err)
	}
}



