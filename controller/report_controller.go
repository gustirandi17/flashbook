package controller

import (
	// "flashbook/constant"
	"flashbook/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	ReportService service.ReportService
}

func NewReportController(reportService service.ReportService) *ReportController {
	return &ReportController{ReportService: reportService}
}

func (rc *ReportController) GetReport(c *gin.Context) {
	// if c.GetString("role") != constant.RoleAdmin {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 	return
	// }

	report, err := rc.ReportService.GetReportData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
