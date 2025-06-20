package controller

import (
	// "flashbook/constant"
	"flashbook/entity"
	"flashbook/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	ScheduleService service.ScheduleService
}

func NewScheduleController(svc service.ScheduleService) *ScheduleController {
	return &ScheduleController{ScheduleService: svc}
}

func (sc *ScheduleController) CreateSchedule(c *gin.Context) {

	var input entity.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := sc.ScheduleService.CreateSchedule(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (sc *ScheduleController) GetAllSchedules(c *gin.Context) {
	schedules, err := sc.ScheduleService.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedules)
}

func (sc *ScheduleController) GetScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	schedule, err := sc.ScheduleService.GetScheduleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func (sc *ScheduleController) UpdateSchedule(c *gin.Context) {
	

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input entity.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := sc.ScheduleService.UpdateSchedule(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (sc *ScheduleController) DeleteSchedule(c *gin.Context) {
	

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := sc.ScheduleService.DeleteSchedule(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}
