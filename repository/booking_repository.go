package repository

import (
	"flashbook/config"
	"flashbook/entity"
)

type BookingRepository interface {
	Create(booking *entity.Booking) error
	GetByUserID(userID uint) ([]entity.Booking, error)
	GetAll() ([]entity.Booking, error)
	GetSchedule(scheduleID uint) (*entity.Schedule, error)
	MarkScheduleBooked(schedule *entity.Schedule) error

	FindByID(id uint) (*entity.Booking, error)
	Save(booking *entity.Booking) error
}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}

func (r *bookingRepository) Create(booking *entity.Booking) error {
	return config.DB.Create(booking).Error
}

func (r *bookingRepository) GetByUserID(userID uint) ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetAll() ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := config.DB.Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetSchedule(scheduleID uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	err := config.DB.First(&schedule, scheduleID).Error
	return &schedule, err
}

func (r *bookingRepository) MarkScheduleBooked(schedule *entity.Schedule) error {
	schedule.IsBooked = true
	return config.DB.Save(schedule).Error
}

func (r *bookingRepository) FindByID(id uint) (*entity.Booking, error) {
	var booking entity.Booking
	err := config.DB.First(&booking, id).Error
	return &booking, err
}

func (r *bookingRepository) Save(booking *entity.Booking) error {
	return config.DB.Save(booking).Error
}