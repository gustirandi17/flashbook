package route

import (
	"flashbook/controller"
	"flashbook/middleware"

	"github.com/gin-gonic/gin"
	"flashbook/constant"
)

func RegisterRoutes(
	r *gin.Engine,
	authController *controller.AuthController,
	bookingController *controller.BookingController,
	paymentController *controller.PaymentController,
	scheduleController *controller.ScheduleController,
	serviceController *controller.ServiceController,
	reportController *controller.ReportController,
) {
	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Protected routes (require JWT)
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())

	// Public access inside protected route
	protected.GET("/protected", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		role, _ := c.Get("role")
		c.JSON(200, gin.H{
			"message": "Access granted",
			"user_id": userID,
			"role":    role,
		})
	})

	// ================== USER ROLE ==================

	// Bookings (customer access)
	bookingGroup := protected.Group("/bookings")
	{
		bookingGroup.POST("", bookingController.CreateBooking)
		bookingGroup.GET("/my", bookingController.GetMyBookings)
	}

	// Payments (customer access)
	paymentGroup := protected.Group("/payments")
	{
		paymentGroup.POST("", paymentController.CreatePayment)
		paymentGroup.GET("/my", paymentController.GetMyPayments)
		paymentGroup.PUT("/:id", paymentController.UpdatePayment)
	}

	// Schedules (open for all roles)
	scheduleGroup := protected.Group("/schedules")
	{
		scheduleGroup.GET("", scheduleController.GetAllSchedules)
		scheduleGroup.GET("/:id", scheduleController.GetScheduleByID)
	}

	// Services (open for all roles)
	serviceGroup := protected.Group("/services")
	{
		serviceGroup.GET("", serviceController.FindAll)
		serviceGroup.GET("/:id", serviceController.FindByID)
	}

	// ================== ADMIN ROLE ==================

	// Admin-only routes using RBAC middleware
	admin := protected.Group("/")
	admin.Use(middleware.RBAC(constant.RoleAdmin))
	{
		// Service management
		admin.POST("/services", serviceController.Create)
		admin.PUT("/services/:id", serviceController.Update)
		admin.DELETE("/services/:id", serviceController.Delete)

		// Schedule management
		admin.POST("/schedules", scheduleController.CreateSchedule)
		admin.PUT("/schedules/:id", scheduleController.UpdateSchedule)
		admin.DELETE("/schedules/:id", scheduleController.DeleteSchedule)

		// Booking & Payment management
		admin.GET("/bookings", bookingController.GetAllBookings)
		admin.GET("/payments", paymentController.GetAllPayments)
		admin.PUT("/payments/:id/status", paymentController.UpdatePaymentStatus)

		// Reports
		admin.GET("/reports", reportController.GetReport)
	}
}