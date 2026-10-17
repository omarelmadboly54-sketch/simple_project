package repository

import "simple_project/internal/models"

type ShowTimeSeatsRepository interface {
	GenerateSeatsForShowTime(showtimeID int64) error
	GetAll(showtime_id int64) (seats []models.ShowtimeSeats,err error)
	GetSeatById(seat_id int64,showtime_id int64) (seat *models.ShowtimeSeats, err error)
	BookSeat(seat_id int64,user_id int64)error
}
