package entity

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `gorm:"unique" json:"booking_id"`
	Method      string    `gorm:"type:enum('transfer','e-wallet','VA','qris')" json:"method"`
	Amount      float64   `json:"amount"`
	Status      string    `gorm:"type:enum('waiting','paid','rejected')" json:"status"`
	PaymentDate string    `json:"payment_date"`
	ProofImage  string    `gorm:"type:text;not null" json:"proof_image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
