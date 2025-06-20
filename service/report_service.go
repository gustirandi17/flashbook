package service

import (
	"flashbook/config"
	"flashbook/constant"
	"flashbook/entity"
)

type TopService struct {
	ServiceID   uint   `json:"service_id"`
	Name        string `json:"name"`
	TotalBooked int64  `json:"total_booked"`
}

type ReportData struct {
	TotalBookings int64        `json:"total_bookings"`
	TotalPaid     int64        `json:"total_paid"`
	TotalIncome   float64      `json:"total_income"`
	TopServices   []TopService `json:"top_services"`
}

type ReportService interface {
	GetReportData() (ReportData, error)
}

type reportService struct{}

func NewReportService() ReportService {
	return &reportService{}
}

func (s *reportService) GetReportData() (ReportData, error) {
	var totalBookings int64
	var totalPaid int64
	var totalIncome float64
	var topServices []TopService

	// Hitung total bookings
	if err := config.DB.Model(&entity.Booking{}).Count(&totalBookings).Error; err != nil {
		return ReportData{}, err
	}

	// Hitung total pembayaran paid
	if err := config.DB.Model(&entity.Payment{}).
		Where("status = ?", constant.PaymentPaid).
		Count(&totalPaid).Error; err != nil {
		return ReportData{}, err
	}

	// Hitung total income
	if err := config.DB.Model(&entity.Payment{}).
		Select("SUM(amount)").Where("status = ?", constant.PaymentPaid).
		Scan(&totalIncome).Error; err != nil {
		return ReportData{}, err
	}

	// Top 3 services berdasarkan booking
	if err := config.DB.Table("bookings").
		Select("services.id as service_id, services.name, COUNT(bookings.id) as total_booked").
		Joins("JOIN schedules ON bookings.schedule_id = schedules.id").
		Joins("JOIN services ON schedules.service_id = services.id").
		Group("services.id").
		Order("total_booked DESC").
		Limit(3).
		Scan(&topServices).Error; err != nil {
		return ReportData{}, err
	}

	return ReportData{
		TotalBookings: totalBookings,
		TotalPaid:     totalPaid,
		TotalIncome:   totalIncome,
		TopServices:   topServices,
	}, nil
}
