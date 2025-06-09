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

	// Validasi: booking harus milik user ini
	var booking entity.Booking
	if err := config.DB.First(&booking, input.BookingID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found"})
		return
	}
	if booking.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not your booking"})
		return
	}

	payment := entity.Payment{
		BookingID:   input.BookingID,
		Method:      input.Method,
		Amount:      input.Amount,
		PaymentDate: input.PaymentDate,
		ProofImage:  input.ProofImage,
		Status:      "unpaid",
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