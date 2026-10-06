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

type movieController struct {
	movieService service.MovieService
	tokenMaker	 token.Maker
}

type MovieController interface{
	Routes(g *gin.RouterGroup)
}

func NewMovieController(movieService service.MovieService,tokenMaker token.Maker)MovieController{
	return &movieController{
		movieService: movieService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *movieController)Routes(g *gin.RouterGroup){
	g.POST("/movies",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Create_Movie)
	g.DELETE("/movies/:id",middleware.Authmiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.Delete_Movie)
	g.GET("/movies",middleware.Authmiddleware(ctl.tokenMaker),ctl.Get_Movies)
	g.GET("/movies/:id",middleware.Authmiddleware(ctl.tokenMaker),ctl.Get_Movie_By_ID)
}


func(ctl *movieController)Create_Movie(g *gin.Context){
	req:=presenters.CreateMovieReq{}
	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	to_movie:=req.ToMovie()

	err:=ctl.movieService.CreateMovie(&to_movie)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusCreated,gin.H{"ok":"ok"})
}


func(ctl *movieController)Get_Movies(g *gin.Context){
	genre:=g.Query("genre")
	page,_:=strconv.Atoi(g.DefaultQuery("page","1"))
	if page<1{
		page = 1
	}
	limit,_:=strconv.Atoi(g.DefaultQuery("limit","5"))
	if limit <1{
		limit=5
	}

	movies,err:=ctl.movieService.Get_Movies(genre,page,limit)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,movies)
}

func(ctl *movieController)Get_Movie_By_ID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	movie,err:=ctl.movieService.Get_Movie_By_ID(id)
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
	g.JSON(http.StatusOK,&movie)
}

func(ctl *movieController)Delete_Movie(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	err=ctl.movieService.Delete_Movie(id)
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


