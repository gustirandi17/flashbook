package service

import (
	"errors"
	"flashbook/entity"
	"flashbook/repository"
)

type ScheduleService interface {
	CreateSchedule(schedule entity.Schedule) (*entity.Schedule, error)
	GetAllSchedules() ([]entity.Schedule, error)
	GetScheduleByID(id uint) (*entity.Schedule, error)
	UpdateSchedule(id uint, input entity.Schedule) (*entity.Schedule, error)
	DeleteSchedule(id uint) error
}

type scheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(r repository.ScheduleRepository) ScheduleService {
	return &scheduleService{repo: r}
}

func (s *scheduleService) CreateSchedule(schedule entity.Schedule) (*entity.Schedule, error) {
	return s.repo.Create(schedule)
}

func (s *scheduleService) GetAllSchedules() ([]entity.Schedule, error) {
	return s.repo.FindAll()
}

func (s *scheduleService) GetScheduleByID(id uint) (*entity.Schedule, error) {
	return s.repo.FindByID(id)
}

func (s *scheduleService) UpdateSchedule(id uint, input entity.Schedule) (*entity.Schedule, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("schedule not found")
	}

	schedule.Date = input.Date
	schedule.TimeSlot = input.TimeSlot
	schedule.ServiceID = input.ServiceID
	schedule.IsBooked = input.IsBooked

	return s.repo.Update(schedule)
}

func (s *scheduleService) DeleteSchedule(id uint) error {
	return s.repo.Delete(id)
}