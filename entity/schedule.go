package entity

import "time"

type Schedule struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ServiceID  uint      `json:"service_id"`
	Date       string    `json:"date"`       // Format: YYYY-MM-DD
	TimeSlot   string    `json:"time_slot"`  // Format: HH:MM:SS
	IsBooked   bool      `json:"is_booked"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
