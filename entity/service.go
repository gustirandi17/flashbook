package entity

import "time"

type Service struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	DurationMinutes int       `json:"duration_minutes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
