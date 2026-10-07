package repository

import (
	"simple_project/internal/models"
	"time"
)

type ShowTime interface {
	Create_ShowTime(showtime *models.ShowTime) error
	Get_showtimes()(showtimes []models.ShowTime,err error)
	Get_showtime(id int64)(showtime *models.ShowTime,err error)
	Delete_showtime (id int64)error
	CheckHallConflict(hallnumber int,showtime time.Time,exludeID int64)(bool,error)
	Update_showtime(showtime *models.ShowTime)error
}




