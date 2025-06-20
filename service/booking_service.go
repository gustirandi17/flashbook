package service

import (
	"errors"
	"flashbook/constant"
	"flashbook/entity"
	"flashbook/repository"
)

type BookingService interface {
	CreateBooking(userID uint, scheduleID uint, notes string) (*entity.Booking, error)
	GetMyBookings(userID uint) ([]entity.Booking, error)
	GetAllBookings() ([]entity.Booking, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(userID uint, scheduleID uint, notes string) (*entity.Booking, error) {
	schedule, err := s.repo.GetSchedule(scheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}
	if schedule.IsBooked {
		return nil, errors.New("schedule already booked")
	}

	booking := &entity.Booking{
		UserID:     userID,
		ScheduleID: scheduleID,
		Status:     constant.StatusPending,
		Notes:      notes,
	}
	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}
	if err := s.repo.MarkScheduleBooked(schedule); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) GetMyBookings(userID uint) ([]entity.Booking, error) {
	return s.repo.GetByUserID(userID)
}

func (s *bookingService) GetAllBookings() ([]entity.Booking, error) {
	return s.repo.GetAll()
}