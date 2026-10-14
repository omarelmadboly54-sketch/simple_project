package repository

import "simple_project/internal/models"

type ShowTimeSeatsRepo interface {
	GenerateSeatsForShowtime(showtimeID int64) error
	Get_Seats(showtime_id int64) (seats []models.ShowtimeSeats,err error)
	Book_seat(seat_id int64,user_id int64)error
}
