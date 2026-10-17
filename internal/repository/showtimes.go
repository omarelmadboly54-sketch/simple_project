package repository

import (
	"simple_project/internal/models"
	"time"
)

type ShowTimeRepository interface {
	CreateShowTime(showtime *models.ShowTime) error
	GetAll()(showtimes []models.ShowTime,err error)
	GetShowTime(id int64)(showtime *models.ShowTime,err error)
	DeleteShowTime (id int64)error
	CheckHallConflict(hallnumber int,showtime time.Time,exludeID int64)(bool,error)
	UpdateShowTime(showtime *models.ShowTime)error
}




