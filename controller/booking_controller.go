package controller

import (
	"flashbook/config"
	"flashbook/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var input struct {
		ScheduleID uint   `json:"schedule_id" binding:"required"`
		Notes      string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah schedule tersedia
	var schedule entity.Schedule
	if err := config.DB.First(&schedule, input.ScheduleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Schedule not found"})
		return
	}
	if schedule.IsBooked {
		c.JSON(http.StatusConflict, gin.H{"error": "Schedule already booked"})
		return
	}

	// buat booking baru
	booking := entity.Booking{
		UserID:     uint(userID.(float64)),
		ScheduleID: input.ScheduleID,
		Status:     "pending",
		Notes:      input.Notes,
	}
	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// tandai schedule sebagai booked
	schedule.IsBooked = true
	config.DB.Save(&schedule)

	c.JSON(http.StatusCreated, booking)
}

func GetMyBookings(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var bookings []entity.Booking
	if err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func GetAllBookings(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
		return
	}
	var bookings []entity.Booking
	if err := config.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
