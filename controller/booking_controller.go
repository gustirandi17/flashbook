package controller

import (
	// "flashbook/constant"
	"flashbook/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	BookingService service.BookingService
}

func NewBookingController(s service.BookingService) *BookingController {
	return &BookingController{BookingService: s}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	var input struct {
		ScheduleID uint   `json:"schedule_id" binding:"required"`
		Notes      string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, err := bc.BookingService.CreateBooking(userID, input.ScheduleID, input.Notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func (bc *BookingController) GetMyBookings(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	bookings, err := bc.BookingService.GetMyBookings(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (bc *BookingController) GetAllBookings(c *gin.Context) {
	// role := c.GetString("role")
	// if role != constant.RoleAdmin {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
	// 	return
	// }
	bookings, err := bc.BookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}