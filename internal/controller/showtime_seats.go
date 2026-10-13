package controller

import (
	"net/http"
	"simple_project/internal/middleware"
	"simple_project/internal/service"
	"simple_project/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

type showtimeseatsController struct {
	showtimeseatsService service.ShowTimeSeatService
	tokenMaker token.Maker
}

type ShowTimeSeatsController interface{
	Routes(g *gin.RouterGroup)
}

func NewShowTimeSeatContoller (showtimeseatsService service.ShowTimeSeatService,tokenMaker token.Maker) ShowTimeSeatsController{
	return &showtimeseatsController{
		showtimeseatsService: showtimeseatsService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *showtimeseatsController)Routes(g *gin.RouterGroup){
	g.GET("/showtimes/:id/seats",middleware.Authmiddleware(ctl.tokenMaker),ctl.Get_Seats)
}


func(ctl *showtimeseatsController)Get_Seats(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	seats,err:=ctl.showtimeseatsService.Get_Seats(id)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}	

	g.JSON(http.StatusOK,gin.H{"seats":seats})
}
