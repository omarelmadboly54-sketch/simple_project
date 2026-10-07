package models

import "time"

type ShowTime struct {
	ID       		int64 		`gorm:"primarykey" json:"id"`
	MovieID  		int64 		`gorm:"not null" json:"movie_id"`
	Movie			Movie		`gorm:"foreignkey:MovieID;constraint:OnDelete:CASCADE;" json:"movie,omitempty"`
	ShowTime 		time.Time	`gorm:"type:timestamp;not null" json:"show_time"`
	Price	 		int			`gorm:"not null" json:"price"`
	HallNumber		int			`gorm:"not null" json:"hall_number"`
	CreatedAt		time.Time	`gorm:"autoCreateTime" json:"-"`
	UpdatedAt		time.Time	`gorm:"autoUpdateTime" json:"-"`
}




