package main

import (
	"log"
	"simple_project/configuration"
	"simple_project/internal/controller"
	"simple_project/internal/providers/postgres"
	"simple_project/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db,config, err := configuration.InitPostgres()
	if err!=nil{
		log.Fatal("failed to connect to DataBase",err)
	}
	log.Println("Database Connection Successfully")

	r:=gin.Default()
	
	userRepo:=postgres.NewUserRepo(db)
	userService:=service.NewUserService(userRepo)
	userController:=controller.NewUserController(userService)
	userController.Routes(r.Group("/api"))





	if err:=r.Run(config.ServerAddress);err!=nil{
		log.Fatalf("failed to start server:%v",err)
	}
}



