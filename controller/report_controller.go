package controller

import (
	"flashbook/config"
	"flashbook/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var totalBookings int64
	var totalPaid int64
	var totalIncome float64

	// Total semua booking
	config.DB.Model(&entity.Booking{}).Count(&totalBookings)

	// Total pembayaran berstatus 'paid'
	config.DB.Model(&entity.Payment{}).
		Where("status = ?", "paid").
		Count(&totalPaid)

	// Total pendapatan dari pembayaran 'paid'
	config.DB.Model(&entity.Payment{}).
		Select("SUM(amount)").Where("status = ?", "paid").
		Scan(&totalIncome)

	// Top layanan (berdasarkan jumlah booking)
	type TopService struct {
		ServiceID   uint   `json:"service_id"`
		Name        string `json:"name"`
		TotalBooked int64  `json:"total_booked"`
	}
	var topServices []TopService

	config.DB.Table("bookings").
		Select("services.id as service_id, services.name, COUNT(bookings.id) as total_booked").
		Joins("JOIN schedules ON bookings.schedule_id = schedules.id").
		Joins("JOIN services ON schedules.service_id = services.id").
		Group("services.id").
		Order("total_booked DESC").
		Limit(3).
		Scan(&topServices)

	c.JSON(http.StatusOK, gin.H{
		"total_bookings": totalBookings,
		"total_paid":     totalPaid,
		"total_income":   totalIncome,
		"top_services":   topServices,
	})
}
