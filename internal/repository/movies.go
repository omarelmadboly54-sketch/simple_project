package repository

import "simple_project/internal/models"

type MovieRepo interface {
	CreateMovie(movie *models.Movie) error
	Get_Movies()(movies []models.Movie,err error)
	Get_Movie_By_ID(id int64)(movie *models.Movie,err error)
	Delete_Movie(id int64)error
}


