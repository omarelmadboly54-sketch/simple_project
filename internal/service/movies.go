package service

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"
	"simple_project/token"
	"time"
)

type movieService struct {
	movieRepo 		repository.MovieRepo
	tokenMaker		token.Maker
	tokenDuration	time.Duration
}


type MovieService interface{
	CreateMovie(movie *models.Movie)error
	Get_Movies(genre string,page int,limit int)(movies []models.Movie,err error)
	Get_Movie_By_ID(id int64)(movie *models.Movie,err error)
	Delete_Movie(id int64)error
}

func NewMovieService(movieRepo repository.MovieRepo,tokenMaker token.Maker,tokenDuration time.Duration) MovieService{
	return &movieService{
		movieRepo: movieRepo,
		tokenMaker: tokenMaker,
		tokenDuration: tokenDuration,
	}
}


func(s *movieService)CreateMovie(movie *models.Movie)error{
	return s.movieRepo.CreateMovie(movie)
}

func(s *movieService)Get_Movies(genre string,page int,limit int)(movies []models.Movie,err error){
	return s.movieRepo.Get_Movies(genre,page,limit)
}


func(s *movieService)Get_Movie_By_ID(id int64)(movie *models.Movie,err error){
	return s.movieRepo.Get_Movie_By_ID(id)
}

func(s *movieService)Delete_Movie(id int64)error{
	movie,err:=s.movieRepo.Get_Movie_By_ID(id)
	if err!=nil{
		return err		
	}
	if err =s.movieRepo.Delete_Movie(movie.ID);err!=nil{
		return err
	}
	return nil
}

