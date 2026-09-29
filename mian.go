package main

import (
	"log"
	"simple_project/configuration"
)

func main() {

	db, err := configuration.InitPostgres()
	if err!=nil{
		log.Fatal("failed to connect to DataBase",err)
	}
	log.Println("Database Connection Successfully")
	_=db

}

