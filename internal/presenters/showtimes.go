package presenters

import (
	"simple_project/internal/models"
	"time"
)

type CreateShowTimeReq struct {
	MovieID  	int64 		`json:"movie_id" binding:"required"`
	ShowTime 	time.Time	`json:"show_time"  binding:"required"`
	Price	 	int			`json:"price" binding:"required"`
	HallNumber	int			`json:"hall_number" binding:"required"`
}

type UpdateShowTimeReq struct{
	ShowTime		*time.Time		`json:"show_time"`
	HallNumber		*int			`json:"hall_number"`
}


func(c *CreateShowTimeReq)ToShowTime()models.ShowTime{
	return models.ShowTime{
		MovieID: c.MovieID,
		ShowTime: c.ShowTime,
		Price: c.Price,
		HallNumber: c.HallNumber,
	}
}


