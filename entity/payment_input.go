package entity

type PaymentInput struct {
	BookingID   uint    `json:"booking_id" binding:"required"`
	Method      string  `json:"method" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	PaymentDate string  `json:"payment_date"`
	ProofImage  string  `json:"proof_image"`
}