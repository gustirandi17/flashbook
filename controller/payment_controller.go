package controller

import (
	"flashbook/config"
	"flashbook/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var input struct {
		BookingID   uint    `json:"booking_id" binding:"required"`
		Method      string  `json:"method" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
		PaymentDate string  `json:"payment_date"`
		ProofImage  string  `json:"proof_image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validasi method pembayaran
	validMethods := map[string]bool{
		"transfer": true,
		"e-wallet": true,
		"VA":       true,
		"qris":     true,
	}
	if !validMethods[input.Method] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment method"})
		return
	}

	// validasi booking harus milik user ini
	var booking entity.Booking
	if err := config.DB.First(&booking, input.BookingID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found"})
		return
	}
	if booking.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not your booking"})
		return
	}

	// validasi proof image
	if input.ProofImage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Proof image is required"})
		return
	}

	payment := entity.Payment{
		BookingID:   input.BookingID,
		Method:      input.Method,
		Amount:      input.Amount,
		PaymentDate: input.PaymentDate,
		ProofImage:  input.ProofImage,
		Status:      "waiting",
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func GetMyPayments(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var payments []entity.Payment
	if err := config.DB.
		Joins("JOIN bookings ON bookings.id = payments.booking_id").
		Where("bookings.user_id = ?", uint(userID.(float64))).
		Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func GetAllPayments(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
		return
	}
	var payments []entity.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func UpdatePaymentStatus(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	var payment entity.Payment

	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validStatus := map[string]bool{"paid": true, "rejected": true}
	if !validStatus[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}



	payment.Status = input.Status
	config.DB.Save(&payment)

	c.JSON(http.StatusOK, payment,)

	// Jika status payment menjadi 'paid', ubah juga status booking menjadi 'confirmed'
	if input.Status == "paid" {
		var booking entity.Booking
		if err := config.DB.First(&booking, payment.BookingID).Error; err == nil {
			booking.Status = "confirmed"
			config.DB.Save(&booking)
		}
	}
}

func UpdatePayment(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	id := c.Param("id")

	var payment entity.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Hanya user pemilik payment yang bisa update
	var booking entity.Booking
	if err := config.DB.First(&booking, payment.BookingID).Error; err != nil || booking.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this payment"})
		return
	}

	// Hanya bisa update jika status saat ini "rejected"
	if payment.Status != "rejected" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only rejected payments can be updated"})
		return
	}

	var input struct {
		Method      string  `json:"method"`
		Amount      float64 `json:"amount"`
		PaymentDate string  `json:"payment_date"`
		ProofImage  string  `json:"proof_image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields & set status ulang ke waiting
	payment.Method = input.Method
	payment.Amount = input.Amount
	payment.PaymentDate = input.PaymentDate
	payment.ProofImage = input.ProofImage
	payment.Status = "waiting"

	config.DB.Save(&payment)
	c.JSON(http.StatusOK, payment)
}
