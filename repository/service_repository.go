package repository

import (
	"flashbook/config"
	"flashbook/entity"
)

type ServiceRepository interface {
	Create(service *entity.Service) (*entity.Service, error)
	FindAll() ([]entity.Service, error)
	FindByID(id uint) (*entity.Service, error)
	Update(service *entity.Service) (*entity.Service, error)
	Delete(id uint) error
}

type serviceRepo struct{}

func NewServiceRepository() ServiceRepository {
	return &serviceRepo{}
}

func (r *serviceRepo) Create(service *entity.Service) (*entity.Service, error) {
	if err := config.DB.Create(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func (r *serviceRepo) FindAll() ([]entity.Service, error) {
	var services []entity.Service
	err := config.DB.Find(&services).Error
	return services, err
}

func (r *serviceRepo) FindByID(id uint) (*entity.Service, error) {
	var service entity.Service
	err := config.DB.First(&service, id).Error
	return &service, err
}

func (r *serviceRepo) Update(service *entity.Service) (*entity.Service, error) {
	err := config.DB.Save(service).Error
	return service, err
}

func (r *serviceRepo) Delete(id uint) error {
	return config.DB.Delete(&entity.Service{}, id).Error
}
