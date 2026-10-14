package postgres

import (
	"errors"
	"fmt"
	"simple_project/internal/models"
	"simple_project/internal/repository"

	"gorm.io/gorm"
)

type showtimeseatsRepo struct {
	db *gorm.DB
}


func NewShowTimeSeatRepo(db *gorm.DB) repository.ShowTimeSeatsRepo{
	return &showtimeseatsRepo{
		db: db,
	}
}

func(r *showtimeseatsRepo)GenerateSeatsForShowtime(showtimeID int64)error{
	rows:=[]string{"A","B","C","D","E"}
	cols:=10

	var seats []models.ShowtimeSeats

	for _,rowLetter:=range rows{
		for j:=1;j<=cols;j++{
			seat_number:=fmt.Sprintf("%s%d",rowLetter,j)
			seats = append(seats, models.ShowtimeSeats{
				ShowtimeID: showtimeID,
				SeatNumber: seat_number,
				Status: "available",
			})

		}
	}
	if err:=r.db.Create(&seats).Error;err!=nil{
		return err
	}
	return nil
}


func(r *showtimeseatsRepo)Get_Seats(showtime_id int64) (seats []models.ShowtimeSeats,err error){
	result:=r.db.Where("showtime_id",showtime_id).Find(&seats)

		if result.Error!=nil{
			return nil,err
		}
		return seats,nil
}


func(r *showtimeseatsRepo)Book_seat(seat_id int64,user_id int64)error{
	result:=r.db.Model(&models.ShowtimeSeats{}).
	Where("id=? AND status =?",seat_id,"available").
	Updates(map[string]interface{}{
		"status":"booked",
		"user_id":user_id,
	})

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return errors.New("seat is already booked")
	}
	return nil

}