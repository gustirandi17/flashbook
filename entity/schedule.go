package entity

import (
	"time"
)

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ServiceID uint      `gorm:"not null;index" json:"service_id"`    // foreign key ke service
	Date      string    `gorm:"type:text;not null" json:"date"`      // Format: YYYY-MM-DD
	TimeSlot  string    `gorm:"type:text;not null" json:"time_slot"` // Format: HH:MM:SS
	IsBooked  bool      `gorm:"default:false" json:"is_booked"`      // default false
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
