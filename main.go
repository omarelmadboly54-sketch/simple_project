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

	tokenMaker,err:=token.NewJWTMaker(config.SecretKey)
	if err!=nil{
		log.Fatal("failed to generate token maker:",err)
	}
	r:=gin.Default()

//users:
	userRepo:=postgres.NewUserRepo(db)
	tokenDuration:=time.Duration(config.AccessTokenMinutes)*time.Minute
	userService:=service.NewUserService(userRepo,tokenMaker,tokenDuration)
	userController:=controller.NewUserController(userService,tokenMaker)
	userController.RegisterRoutes(r.Group("/api"))
//movies:
	movieRepo:=postgres.NewMovieRepo(db)
	movieService:=service.NewMovieService(movieRepo)
	movieController:=controller.NewMovieController(movieService,tokenMaker)
	movieController.RegisterRoutes(r.Group("/api"))
//showtime:
	seatRepo:=postgres.NewShowTimeSeatRepo(db)
	showtimeRepo:=postgres.NewShowTimeRepo(db)
	showtimeService:=service.NewShowTimeService(showtimeRepo,seatRepo)
	showtimeController:=controller.NewShowtimeController(showtimeService,tokenMaker)
	showtimeController.RegisterRoutes(r.Group("/api"))

//showtime_seats:
	seatService:=service.NewShowTimeSeatService(seatRepo)
	seatController:=controller.NewShowTimeSeatContoller(seatService,tokenMaker)
	seatController.RegisterRoutes(r.Group("/api"))

	if err:=r.Run(config.ServerAddress);err!=nil{
		log.Fatalf("failed to start server:%v",err)
	}
}




