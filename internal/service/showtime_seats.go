package service

import (
	"errors"
	"simple_project/internal/apperrors"
	"simple_project/internal/models"
	"simple_project/internal/presenters"
	"simple_project/internal/repository"
)

type showtimeseatService struct {
	showtimeseatsRepo repository.ShowTimeSeatsRepository
}

type ShowTimeSeatService interface{
	GetAll(showtime_id int64) ([]presenters.SeatResponse,error)
	BookSeat(seat_id int64,user_id int64)error
	GetSeatById(seat_id int64,showtime_id int64) (seat *models.ShowtimeSeats, err error)
}


func NewShowTimeSeatService(showtimeseatsrepo repository.ShowTimeSeatsRepository) ShowTimeSeatService{
	return &showtimeseatService{
		showtimeseatsRepo: showtimeseatsrepo,
		
	}
}

func(s *showtimeseatService)GetAll(showtime_id int64) ([]presenters.SeatResponse,error){
	
	seats,err:=s.showtimeseatsRepo.GetAll(showtime_id)
	if err!=nil{
		return nil,err
	}

	var response []presenters.SeatResponse
	for _,seat:=range seats{
		response = append(response, presenters.SeatResponse{
			SeatNumber: seat.SeatNumber,
			Status: seat.Status,
		})
	}
	return response,nil
}

func(s *showtimeseatService)GetSeatById(seat_id int64,showtime_id int64) (seat *models.ShowtimeSeats, err error){
	return s.showtimeseatsRepo.GetSeatById(seat_id,showtime_id)
}




func(s *showtimeseatService)BookSeat(seat_id int64,user_id int64)error{
	 err:=s.showtimeseatsRepo.BookSeat(seat_id,user_id)
	 if err!=nil{
		if err.Error()==apperrors.ErrSeatAlreadyBooked.Error(){
			return errors.New("Conflict:this seat is no longer available")
		}
		return err
	 }
	 return nil
}
