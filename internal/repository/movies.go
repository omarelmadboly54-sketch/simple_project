package repository

import "simple_project/internal/models"

type MovieRepository interface {
	CreateMovie(movie *models.Movie) error
	GetAll(genre string,page int,limit int)(movies []models.Movie,err error)
	GetMovieByID(id int64)(movie *models.Movie,err error)
	DeleteMovie(id int64)error
}


