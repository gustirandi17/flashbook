package controller

import (
	// "flashbook/constant"
	"flashbook/entity"
	"flashbook/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: paymentService}
}

func (pc *PaymentController) CreatePayment(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	var input entity.PaymentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := pc.PaymentService.CreatePayment(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (pc *PaymentController) GetMyPayments(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))

	payments, err := pc.PaymentService.GetPaymentsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (pc *PaymentController) GetAllPayments(c *gin.Context) {
	payments, err := pc.PaymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (pc *PaymentController) UpdatePaymentStatus(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.PaymentService.UpdatePaymentStatus(uint(idUint), input.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

func (pc *PaymentController) UpdatePayment(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input entity.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.PaymentService.UpdatePayment(uint(idUint), input, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}