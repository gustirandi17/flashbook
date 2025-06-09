package controller

import (
	"flashbook/config"
	"flashbook/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSchedule(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input entity.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func GetAllSchedules(c *gin.Context) {
	var schedules []entity.Schedule
	if err := config.DB.Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func GetScheduleByID(c *gin.Context) {
	id := c.Param("id")
	var schedule entity.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func UpdateSchedule(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	var schedule entity.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	var input entity.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&schedule).Updates(input)

	c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	var schedule entity.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	config.DB.Delete(&schedule)

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted"})
}
