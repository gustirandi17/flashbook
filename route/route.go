package route

import (
	"flashbook/controller"
	"flashbook/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/protected", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			role, _ := c.Get("role")
			c.JSON(200, gin.H{
				"message": "Access granted",
				"user_id": userID,
				"role":    role,
			})
		})

		// Service routes
		protected.POST("/services", controller.CreateService)
		protected.GET("/services", controller.GetAllServices)
		protected.GET("/services/:id", controller.GetServiceByID)
		protected.PUT("/services/:id", controller.UpdateService)
		protected.DELETE("/services/:id", controller.DeleteService)

		// Booking routes
		protected.POST("/bookings", controller.CreateBooking)
		protected.GET("/bookings/my", controller.GetMyBookings)
		protected.GET("/bookings", controller.GetAllBookings)

		// Payment routes
		protected.POST("/payments", controller.CreatePayment)
		protected.GET("/payments/my", controller.GetMyPayments)
		protected.GET("/payments", controller.GetAllPayments)

		// Schedule routes
		protected.POST("/schedules", controller.CreateSchedule)     
		protected.GET("/schedules", controller.GetAllSchedules)      
		protected.GET("/schedules/:id", controller.GetScheduleByID) 
		protected.PUT("/schedules/:id", controller.UpdateSchedule)  
		protected.DELETE("/schedules/:id", controller.DeleteSchedule) 

		// Reports routes
		protected.GET("/reports", controller.GetReport)

		// Update Payment
		protected.PUT("/payments/:id/status", controller.UpdatePaymentStatus)
		protected.PUT("/payments/:id", controller.UpdatePayment)

		
	}
}
