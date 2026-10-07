package postgres

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"
	"time"

	"gorm.io/gorm"
)

type showtimeRepo struct {
	db *gorm.DB
}

func NewShowTimeRepo(db *gorm.DB)repository.ShowTime{
	return &showtimeRepo{
		db:db,
	}
}

func(r *showtimeRepo)Create_ShowTime(showtime *models.ShowTime) error{
	result:=r.db.Create(&showtime)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

func(r *showtimeRepo)Get_showtimes()(showtimes []models.ShowTime,err error){
	result:=r.db.Preload("Movie").Find(&showtimes)
	if result.Error!=nil{
		return nil,result.Error
	}
	return showtimes,nil
}

func(r *showtimeRepo)Get_showtime(id int64)(showtime *models.ShowTime,err error){
	result:=r.db.Where("id=?",id).Preload("Movie").First(&showtime)
	if result.Error!=nil{
		return nil,result.Error
	}
	return showtime,nil
}

func(r *showtimeRepo)Delete_showtime (id int64)error{
	result:=r.db.Delete(&models.ShowTime{},id)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

func(r *showtimeRepo)CheckHallConflict(hallnumber int,showtime time.Time,exludeID int64)(bool,error){
	var count int64

	err:=r.db.Model(&models.ShowTime{}).Where("hall_number=? AND show_time=? AND id != ?",hallnumber,showtime,exludeID).Count(&count).Error
	if err!=nil{
		return false,err
	}
	return count >0,nil
}

func(r *showtimeRepo)Update_showtime(showtime *models.ShowTime)error{
	return r.db.Save(showtime).Error
}



