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

type movieController struct {
	movieService service.MovieService
	tokenMaker	 token.Maker
}

type MovieController interface{
	RegisterRoutes(g *gin.RouterGroup)
}

func NewMovieController(movieService service.MovieService,tokenMaker token.Maker)MovieController{
	return &movieController{
		movieService: movieService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *movieController)RegisterRoutes(g *gin.RouterGroup){
	g.POST("/movies",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.CreateMovie)
	g.DELETE("/movies/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.DeleteMovie)
	g.GET("/movies",ctl.GetAll)
	g.GET("/movies/:id",ctl.GetMovieByID)
}


func(ctl *movieController)CreateMovie(g *gin.Context){
	req:=presenters.CreateMovieReq{}
	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	toMovie:=req.ToMovie()

	err:=ctl.movieService.CreateMovie(&toMovie)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusCreated,gin.H{"ok":"ok"})
}


func(ctl *movieController)GetAll(g *gin.Context){
	genre:=g.Query("genre")
	page,_:=strconv.Atoi(g.DefaultQuery("page","1"))
	if page<1{
		page = 1
	}
	limit,_:=strconv.Atoi(g.DefaultQuery("limit","5"))
	if limit <1{
		limit=5
	}

	movies,err:=ctl.movieService.GetAll(genre,page,limit)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,movies)
}

func(ctl *movieController)GetMovieByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	movie,err:=ctl.movieService.GetMovieByID(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrMovieNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,&movie)
}

func(ctl *movieController)DeleteMovie(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	err=ctl.movieService.DeleteMovie(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrMovieNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{"ok":"ok"})
}


