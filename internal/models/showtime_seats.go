package models

import "time"

type ShowtimeSeats struct {
	ID         int64  		`gorm:"primarykey" json:"id"`
	ShowtimeID int64  		`gorm:"foreignkey:ShowtimeID;constraint:OnDelete:CASCADE;" json:"showtime_id"`
	SeatNumber string 		`gorm:"not null" json:"seat_number"`
	Status     string 		`gorm:"not null" json:"status"`
	UserID     *int64  		`gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"user_id"`
	CreatedAt  time.Time	`gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time	`gorm:"autoUpdateTime" json:"-"`	 
}
