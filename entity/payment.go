package entity

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `gorm:"not null;unique" json:"booking_id"`
	Method      string    `gorm:"type:enum('transfer','e-wallet','VA','qris');not null" json:"method"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Status      string    `gorm:"type:enum('waiting','paid','rejected');default:'waiting'" json:"status"`
	PaymentDate string    `gorm:"type:text;not null" json:"payment_date"`
	ProofImage  string    `gorm:"type:text;not null" json:"proof_image"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
