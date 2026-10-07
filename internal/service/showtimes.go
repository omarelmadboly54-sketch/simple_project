package service

import (
	"errors"
	"simple_project/internal/models"
	"simple_project/internal/presenters"
	"simple_project/internal/repository"
	"simple_project/token"
	"time"
)

type showtimeService struct {
	showtimeRepo 	repository.ShowTime
	tokenMaker	 	token.Maker
	tokenDuration	time.Duration
}

type ShowTimeService interface{
	Create_ShowTime(showtime *models.ShowTime) error
	Get_showtimes()(showtimes []models.ShowTime,err error)
	Get_showtime(id int64)(showtime *models.ShowTime,err error)
	Delete_showtime (id int64)error
	Update_showtime(id int64,req *presenters.UpdateShowTimeReq)error
}


func NewShowTimeService(showtimeRepo repository.ShowTime,tokenMaker token.Maker,tokenDuration time.Duration)ShowTimeService{
	return &showtimeService{
		showtimeRepo: showtimeRepo,
		tokenMaker: tokenMaker,
		tokenDuration: tokenDuration,
	}
}

func(s *showtimeService)Create_ShowTime(showtime *models.ShowTime) error{
	return s.showtimeRepo.Create_ShowTime(showtime)

}

func(s *showtimeService)Get_showtimes()(showtimes []models.ShowTime,err error){
	return s.showtimeRepo.Get_showtimes()
}

func(s *showtimeService)Get_showtime(id int64)(showtime *models.ShowTime,err error){
	return s.showtimeRepo.Get_showtime(id)
}


func(s *showtimeService)Delete_showtime (id int64)error{
	showtime,err:=s.showtimeRepo.Get_showtime(id)
	if err!=nil{
		return err
	}

	if err:=s.showtimeRepo.Delete_showtime(showtime.ID);err!=nil{
		return err
	}
	return nil
}


func(s *showtimeService)Update_showtime(id int64,req *presenters.UpdateShowTimeReq)error{
	
	existingShowtime,err:=s.showtimeRepo.Get_showtime(id)
	if err!=nil{
		return err
	}
	
	targetHall:=existingShowtime.HallNumber
	if req.HallNumber!=nil{
		targetHall=*req.HallNumber
	}

	targetTime:=existingShowtime.ShowTime
	if req.ShowTime!=nil{
		targetTime=*req.ShowTime
	}

	hasconflict,err:=s.showtimeRepo.CheckHallConflict(targetHall,targetTime,id)
	if err!=nil{
		return err
	}
	if hasconflict{
		return errors.New("hall is already booked at this time slot")
	}

	if req.HallNumber!=nil{
		existingShowtime.HallNumber=*req.HallNumber
	}
	if req.ShowTime!=nil{
		existingShowtime.ShowTime=*req.ShowTime
	}

	return s.showtimeRepo.Update_showtime(existingShowtime)
}
