package controller

import (
	// "flashbook/constant"
	"flashbook/entity"
	"flashbook/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	ServiceService service.ServiceService
}

func NewServiceController(svc service.ServiceService) *ServiceController {
	return &ServiceController{
		ServiceService: svc,
	}
}

func (sc *ServiceController) Create(c *gin.Context) {
	// if c.GetString("role") != constant.RoleAdmin {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
	// 	return
	// }

	var input entity.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := sc.ServiceService.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (sc *ServiceController) FindAll(c *gin.Context) {
	result, err := sc.ServiceService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (sc *ServiceController) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := sc.ServiceService.FindByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (sc *ServiceController) Update(c *gin.Context) {
	// if c.GetString("role") != constant.RoleAdmin {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
	// 	return
	// }

	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input entity.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := sc.ServiceService.Update(uint(idUint), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (sc *ServiceController) Delete(c *gin.Context) {
	// if c.GetString("role") != constant.RoleAdmin {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
	// 	return
	// }

	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = sc.ServiceService.Delete(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
