package service

import (
	"errors"
	"simple_project/internal/presenters"
	"simple_project/internal/repository"
	"simple_project/token"
	"time"
)

type showtimeseatService struct {
	showtimeseatsRepo repository.ShowTimeSeatsRepo
	tokenMaker token.Maker
	tokenDuration time.Duration
}

type ShowTimeSeatService interface{
	Get_Seats(showtime_id int64) ([]presenters.SeatResponse,error)
	Book_seat(seat_id int64,user_id int64)error
}


func NewShowTimeSeatService(showtimeseatsrepo repository.ShowTimeSeatsRepo,tokenMaker token.Maker,tokenDuration time.Duration) ShowTimeSeatService{
	return &showtimeseatService{
		showtimeseatsRepo: showtimeseatsrepo,
		tokenMaker: tokenMaker,
		tokenDuration: tokenDuration,
	}
}

func(s *showtimeseatService)Get_Seats(showtime_id int64) ([]presenters.SeatResponse,error){
	
	seats,err:=s.showtimeseatsRepo.Get_Seats(showtime_id)
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

func(s *showtimeseatService)Book_seat(seat_id int64,user_id int64)error{
	 err:=s.showtimeseatsRepo.Book_seat(seat_id,user_id)
	 if err!=nil{
		if err.Error()=="seat is already booked"{
			return errors.New("Conflict:this seat is no longer available")
		}
		return err
	 }
	 return nil
}
