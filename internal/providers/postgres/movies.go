package postgres

import (
	"simple_project/internal/models"
	"simple_project/internal/repository"

	"gorm.io/gorm"
)

type movieRepo struct {
	db *gorm.DB
}


func NewMovieRepo(db *gorm.DB)repository.MovieRepo{
	return&movieRepo{
		db: db,
	}
}


func(r *movieRepo)CreateMovie(movie *models.Movie)error{
	result:=r.db.Create(&movie)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

func(r *movieRepo)Get_Movies(genre string,page int,limit int)(movies []models.Movie,err error){
	query:=r.db.Model(&movies)

	if genre!=""{
		query = query.Where("genre=?",genre)
	}
	offset:=(page-1)*limit
	if err:=query.Limit(limit).Offset(offset).Find(&movies).Error;err!=nil{
		return nil,err
	}
	return movies,nil
}

func(r *movieRepo)Get_Movie_By_ID(id int64)(movie *models.Movie,err error){
	result:=r.db.Where("id=?",id).First(&movie)
	if result.Error!=nil{
		return nil,result.Error
	}
	return movie,nil
}

func(r *movieRepo)Delete_Movie(id int64)error{
	result:=r.db.Delete(&models.Movie{},id)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}






