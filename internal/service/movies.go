package service

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"
)

type movieService struct {
	movieRepo 		repository.MovieRepository
}


type MovieService interface{
	CreateMovie(movie *models.Movie)error
	GetAll(genre string,page int,limit int)(movies []models.Movie,err error)
	GetMovieByID(id int64)(movie *models.Movie,err error)
	DeleteMovie(id int64)error
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService{
	return &movieService{
		movieRepo: movieRepo,
	
	}
}


func(s *movieService)CreateMovie(movie *models.Movie)error{
	return s.movieRepo.CreateMovie(movie)
}

func(s *movieService)GetAll(genre string,page int,limit int)(movies []models.Movie,err error){
	return s.movieRepo.GetAll(genre,page,limit)
}


func(s *movieService)GetMovieByID(id int64)(movie *models.Movie,err error){
	return s.movieRepo.GetMovieByID(id)
}

func(s *movieService)DeleteMovie(id int64)error{
	movie,err:=s.movieRepo.GetMovieByID(id)
	if err!=nil{
		return err		
	}
	if err =s.movieRepo.DeleteMovie(movie.ID);err!=nil{
		return err
	}
	return nil
}

