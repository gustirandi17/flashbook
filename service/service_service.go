package service

import (
	"errors"
	"flashbook/entity"
	"flashbook/repository"
)

type ServiceService interface {
	Create(input entity.Service) (*entity.Service, error)
	FindAll() ([]entity.Service, error)
	FindByID(id uint) (*entity.Service, error)
	Update(id uint, input entity.Service) (*entity.Service, error)
	Delete(id uint) error
}

type serviceService struct {
	repo repository.ServiceRepository
}

func NewServiceService(repo repository.ServiceRepository) ServiceService {
	return &serviceService{repo: repo}
}

func (s *serviceService) Create(input entity.Service) (*entity.Service, error) {
	return s.repo.Create(&input)
}

func (s *serviceService) FindAll() ([]entity.Service, error) {
	return s.repo.FindAll()
}

func (s *serviceService) FindByID(id uint) (*entity.Service, error) {
	return s.repo.FindByID(id)
}

func (s *serviceService) Update(id uint, input entity.Service) (*entity.Service, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("service not found")
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.Price = input.Price
	existing.DurationMinutes = input.DurationMinutes

	return s.repo.Update(existing)
}

func (s *serviceService) Delete(id uint) error {
	return s.repo.Delete(id)
}
