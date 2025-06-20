package repository

import (
	"flashbook/config"
	"flashbook/entity"
)

type ScheduleRepository interface {
	Create(schedule entity.Schedule) (*entity.Schedule, error)
	FindAll() ([]entity.Schedule, error)
	FindByID(id uint) (*entity.Schedule, error)
	Update(schedule *entity.Schedule) (*entity.Schedule, error)
	Delete(id uint) error
}

type scheduleRepository struct{}

func NewScheduleRepository() ScheduleRepository {
	return &scheduleRepository{}
}

func (r *scheduleRepository) Create(schedule entity.Schedule) (*entity.Schedule, error) {
	if err := config.DB.Create(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule
	if err := config.DB.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) FindByID(id uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	if err := config.DB.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) Update(schedule *entity.Schedule) (*entity.Schedule, error) {
	if err := config.DB.Save(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *scheduleRepository) Delete(id uint) error {
	if err := config.DB.Delete(&entity.Schedule{}, id).Error; err != nil {
		return err
	}
	return nil
}
