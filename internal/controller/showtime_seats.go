package controller

import (
	"net/http"
	"simple_project/internal/apperrors"
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
	RegisterRoutes(g *gin.RouterGroup)
}

func NewShowTimeSeatContoller (showtimeseatsService service.ShowTimeSeatService,tokenMaker token.Maker) ShowTimeSeatsController{
	return &showtimeseatsController{
		showtimeseatsService: showtimeseatsService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *showtimeseatsController)RegisterRoutes(g *gin.RouterGroup){
	g.GET("/showtimes/:id/seats",middleware.AuthMiddleware(ctl.tokenMaker),ctl.GetAll)
	g.PUT("/showtimes/:showtime_id/seats/:seat_id",middleware.AuthMiddleware(ctl.tokenMaker),ctl.BookSeat)
}


func(ctl *showtimeseatsController)GetAll(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	seats,err:=ctl.showtimeseatsService.GetAll(id)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}	

	g.JSON(http.StatusOK,gin.H{"seats":seats})
}

func(ctl *showtimeseatsController)BookSeat(g *gin.Context){
	seatIDStr:=g.Param("seat_id")
	seatID,err:=strconv.ParseInt(seatIDStr,10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid seat ID format"})
		return
	}


	showtimeIDStr:=g.Param("showtime_id")
	showtimeID,err:=strconv.ParseInt(showtimeIDStr,10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid showtime ID format"})
		return
	}

		_,err=ctl.showtimeseatsService.GetSeatById(seatID,showtimeID)
	if err!=nil{
		g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrSeatNotFound.Error()})
		return
	}

	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}
	userID:=payload.UserID

	err = ctl.showtimeseatsService.BookSeat(seatID,userID)
	if err!=nil{
		if err.Error()=="Conflict:this seat is no longer available"{
			g.JSON(http.StatusConflict,gin.H{"error":apperrors.ErrSeatAlreadyBooked.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{
		"message_booked":"seat successfully booked",
		"seat_Id":seatID,
})	
}
