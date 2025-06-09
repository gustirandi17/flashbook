package entity

import "time"

type Booking struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	ScheduleID uint      `gorm:"unique" json:"schedule_id"`
	Status     string    `gorm:"type:enum('pending','confirmed','cancelled','completed')" json:"status"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
