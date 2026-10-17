package service

import (
	"simple_project/internal/apperrors"
	"simple_project/internal/models"
	"simple_project/internal/presenters"
	"simple_project/internal/repository"
)

type showtimeService struct {
	showtimeRepo 	repository.ShowTimeRepository
	seatRepo		repository.ShowTimeSeatsRepository
	
}

type ShowTimeService interface{
	CreateShowTime(showtime *models.ShowTime) error
	GetAll()(showtimes []models.ShowTime,err error)
	GetShowTime(id int64)(showtime *models.ShowTime,err error)
	DeleteShowTime (id int64)error
	UpdateShowTime(id int64,req *presenters.UpdateShowTimeReq)error
}


func NewShowTimeService(showtimeRepo repository.ShowTimeRepository,seatRepo repository.ShowTimeSeatsRepository)ShowTimeService{
	return &showtimeService{
		showtimeRepo: showtimeRepo,
		seatRepo: seatRepo,
	}
}

func(s *showtimeService)CreateShowTime(showtime *models.ShowTime) error{
	if err:= s.showtimeRepo.CreateShowTime(showtime);err!=nil{
		return err
	}
	if err:=s.seatRepo.GenerateSeatsForShowTime(showtime.ID);err!=nil{
		return err
	}
	return nil

}

func(s *showtimeService)GetAll()(showtimes []models.ShowTime,err error){
	return s.showtimeRepo.GetAll()
}

func(s *showtimeService)GetShowTime(id int64)(showtime *models.ShowTime,err error){
	return s.showtimeRepo.GetShowTime(id)
}


func(s *showtimeService)DeleteShowTime (id int64)error{
	showtime,err:=s.showtimeRepo.GetShowTime(id)
	if err!=nil{
		return err
	}

	if err:=s.showtimeRepo.DeleteShowTime(showtime.ID);err!=nil{
		return err
	}
	return nil
}


func(s *showtimeService)UpdateShowTime(id int64,req *presenters.UpdateShowTimeReq)error{
	
	existingShowtime,err:=s.showtimeRepo.GetShowTime(id)
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
		return apperrors.ErrHallAlreadyBooked
	}

	if req.HallNumber!=nil{
		existingShowtime.HallNumber=*req.HallNumber
	}
	if req.ShowTime!=nil{
		existingShowtime.ShowTime=*req.ShowTime
	}

	return s.showtimeRepo.UpdateShowTime(existingShowtime)
}
