package models

import "time"

type User struct {
	ID        int64  	`gorm:"primarykey" json:"id"`
	Name      string 	`gorm:"not null" json:"name"`
	Email     string 	`gorm:"unique;not null" json:"email"`
	Password  string 	`gorm:"not null" json:"password"`
	Role      string 	`gorm:"deafult:'user'" json:"role"`
	CreatedAt time.Time  `gorm:"autocreatetime" json:"created_at"`
}
