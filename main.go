package main

import (
	"flashbook/config"
	// "flashbook/controller"
	"flashbook/entity"
	"flashbook/route"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()


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