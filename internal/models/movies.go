package models

import "time"

type Movie struct {
	ID          int64  		`gorm:"primarykey" json:"id"`
	Title       string 		`gorm:"unique;not null" json:"title"`
	Description string 		`gorm:"not null" json:"description"`
	Image       string 		`gorm:"not null" json:"image"`
	Genre       string 		`gorm:"not null" json:"genre"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}


