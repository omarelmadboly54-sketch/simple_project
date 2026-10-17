package controller

import (
	"errors"
	"net/http"
	"simple_project/internal/apperrors"
	"simple_project/internal/middleware"
	"simple_project/internal/presenters"
	"simple_project/internal/service"
	"simple_project/token"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type showtimeController struct {
	showtimeService service.ShowTimeService
	tokenMaker		token.Maker
}

type ShowTimeController interface{
	RegisterRoutes(g *gin.RouterGroup)
}

func NewShowtimeController(showtimeService service.ShowTimeService,tokenMaker token.Maker)ShowTimeController{
	return &showtimeController{
		showtimeService: showtimeService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *showtimeController)RegisterRoutes(g *gin.RouterGroup){
	g.POST("/showtimes",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.CreateShowTime)
	g.GET("/showtimes",ctl.GetAll)
	g.GET("/showtimes/:id",ctl.GetShowTime)
	g.DELETE("/showtimes/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.DeleteShowTime)
	g.PATCH("/showtimes/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.UpdateShowTime)
}

func(ctl *showtimeController)CreateShowTime(g *gin.Context){
	req:=presenters.CreateShowTimeReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	toShowTime:=req.ToShowTime()
	err:=ctl.showtimeService.CreateShowTime(&toShowTime)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusCreated,gin.H{"ok":"ok"})
}

func(ctl *showtimeController)GetAll(g *gin.Context){
	
	showtimes,err:=ctl.showtimeService.GetAll()
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,showtimes)
}


func(ctl *showtimeController)GetShowTime(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!= nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	showtime,err:=ctl.showtimeService.GetShowTime(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrShowtimeNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,&showtime)
}


func(ctl *showtimeController)DeleteShowTime(g  *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	err = ctl.showtimeService.DeleteShowTime(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrShowtimeNotFound.Error()})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}


func(ctl *showtimeController)UpdateShowTime(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	req:=presenters.UpdateShowTimeReq{}
	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	err=ctl.showtimeService.UpdateShowTime(id,&req)
	if err!=nil{
		if err.Error() == "record not found"{
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrShowtimeNotFound.Error()})
		return
		}
		if err.Error()== "hall is already booked at this time slot"{
			g.JSON(http.StatusConflict,gin.H{"error":apperrors.ErrHallAlreadyBooked.Error()})
		return
		}

		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}
