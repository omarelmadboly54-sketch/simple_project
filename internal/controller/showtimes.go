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

type showtimeController struct {
	showtimeService service.ShowTimeService
	tokenMaker		token.Maker
}

type ShowTimeController interface{
	Routes(g *gin.RouterGroup)
}

func NewShowtimeController(showtimeService service.ShowTimeService,tokenMaker token.Maker)ShowTimeController{
	return &showtimeController{
		showtimeService: showtimeService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *showtimeController)Routes(g *gin.RouterGroup){
	g.POST("/showtimes",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Create_showtime)
	g.GET("/showtimes",ctl.Get_showtimes)
	g.GET("/showtimes/:id",ctl.Get_Showtime)
	g.DELETE("/showtimes/:id",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Delete_showtime)
	g.PATCH("/showtimes/:id",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Update_showtime)
}

func(ctl *showtimeController)Create_showtime(g *gin.Context){
	req:=presenters.CreateShowTimeReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	to_showtime:=req.ToShowTime()
	err:=ctl.showtimeService.Create_ShowTime(&to_showtime)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusCreated,gin.H{"ok":"ok"})
}

func(ctl *showtimeController)Get_showtimes(g *gin.Context){
	
	showtimes,err:=ctl.showtimeService.Get_showtimes()
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,showtimes)
}


func(ctl *showtimeController)Get_Showtime(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!= nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	showtime,err:=ctl.showtimeService.Get_showtime(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":err.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,&showtime)
}


func(ctl *showtimeController)Delete_showtime(g  *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	err = ctl.showtimeService.Delete_showtime(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":err.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}


func(ctl *showtimeController)Update_showtime(g *gin.Context){
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

	err=ctl.showtimeService.Update_showtime(id,&req)
	if err!=nil{
		if err.Error() == "record not found"{
			g.JSON(http.StatusNotFound,gin.H{"error":"showtime not found",
		})
		return
		}
		if err.Error()== "hall is already booked at this time slot"{
			g.JSON(http.StatusConflict,gin.H{"error":err.Error(),
		})
		return
		}

		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}
