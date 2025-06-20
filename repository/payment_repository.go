package repository

import (
	"flashbook/config"
	"flashbook/entity"
)

type PaymentRepository interface {
	FindByID(id uint) (*entity.Payment, error)
	FindAll() ([]entity.Payment, error)
	FindByUserID(userID uint) ([]entity.Payment, error)
	Create(payment *entity.Payment) (*entity.Payment, error)
	Save(payment *entity.Payment) (*entity.Payment, error)
}

type paymentRepo struct{}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) FindByID(id uint) (*entity.Payment, error) {
	var payment entity.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepo) FindAll() ([]entity.Payment, error) {
	var payments []entity.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepo) FindByUserID(userID uint) ([]entity.Payment, error) {
	var payments []entity.Payment
	if err := config.DB.
		Joins("JOIN bookings ON bookings.id = payments.booking_id").
		Where("bookings.user_id = ?", userID).
		Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepo) Create(payment *entity.Payment) (*entity.Payment, error) {
	if err := config.DB.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepo) Save(payment *entity.Payment) (*entity.Payment, error) {
	if err := config.DB.Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}
