package entity

import "time"

type Service struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Price           float64   `gorm:"not null" json:"price"`
	DurationMinutes int       `gorm:"not null" json:"duration_minutes"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}