package main

import (
	"flashbook/config"
	"flashbook/entity"
	"flashbook/route"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.InitDB()
	seedAdmin()

	//Auto migrate all entities
	err := config.DB.AutoMigrate(
		&entity.User{},
		&entity.Service{},
		&entity.Schedule{},
		&entity.Booking{},
		&entity.Payment{},
	)

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	r := gin.Default()
	route.RegisterRoutes(r)

	r.Run()
}

func seedAdmin() {
	var count int64
	config.DB.Model(&entity.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := entity.User{
			Name:     "admin",
			Email:    "admin1@gmail.com",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		config.DB.Create(&admin)
	}
}
