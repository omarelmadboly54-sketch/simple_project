package configuration

import (
	"log"
	"simple_project/internal/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB,error){
	config,err:=util.LoadConfig(".")
	if err!=nil{
		log.Fatal("failed to load config:",err)
	}
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err!=nil{
		return nil,err
	}
	return db,nil
}

