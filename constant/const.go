package constant

const (
	// Role
	RoleAdmin    = "admin"
	RoleCustomer = "customer"

	// Booking Status
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
	StatusCompleted = "completed"

	// Payment Status
	PaymentWaiting  = "waiting"
	PaymentPaid     = "paid"
	PaymentRejected = "rejected"
)

var ValidPaymentMethods = map[string]bool{
	"transfer": true,
	"e-wallet": true,
	"VA":       true,
	"qris":     true,
}

func IsValidPaymentMethod(method string) bool {
	return ValidPaymentMethods[method]
}

var ValidPaymentStatus = map[string]bool{
	PaymentPaid:     true,
	PaymentRejected: true,
}

func IsValidPaymentStatus(status string) bool {
	return ValidPaymentStatus[status]
}