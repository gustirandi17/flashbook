package main

import (
	"flashbook/config"
	// "flashbook/constant"
	"flashbook/entity"
	"flashbook/route"

	"flashbook/controller"
	"flashbook/repository"
	"flashbook/service"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	log.Println("✅ DB initialized")

	// Auto migrate
	config.DB.AutoMigrate(
	&entity.User{},
	&entity.Service{},
	&entity.Schedule{},
	&entity.Booking{},
	&entity.Payment{},
	)
	log.Println("✅ Auto migration complete")


	// Dependency Injection
	bookingRepo := repository.NewBookingRepository()
	paymentRepo := repository.NewPaymentRepository()
	scheduleRepo := repository.NewScheduleRepository()
	serviceRepo := repository.NewServiceRepository()
	userRepo := repository.NewUserRepository()
	if err := userRepo.SeedAdminIfNotExists(); err != nil {
	log.Fatal("Failed to seed admin:", err)
}

	

	bookingService := service.NewBookingService(bookingRepo)
	paymentService := service.NewPaymentService(paymentRepo, bookingRepo)
	scheduleService := service.NewScheduleService(scheduleRepo)
	serviceService := service.NewServiceService(serviceRepo)
	reportService := service.NewReportService()
	authService := service.NewAuthService(userRepo)

	authController := controller.NewAuthController(authService)
	bookingController := controller.NewBookingController(bookingService)
	paymentController := controller.NewPaymentController(paymentService)
	scheduleController := controller.NewScheduleController(scheduleService)
	serviceController := controller.NewServiceController(serviceService)
	reportController := controller.NewReportController(reportService)

	// Setup Router
	r := gin.Default()
	route.RegisterRoutes(
		r,
		authController,
		bookingController,
		paymentController,
		scheduleController,
		serviceController,
		reportController,
	)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}


