package service

import (
	"errors"
	"flashbook/constant"
	"flashbook/entity"
	"flashbook/repository"
)

type PaymentService interface {
	CreatePayment(userID uint, input entity.PaymentInput) (*entity.Payment, error)
	GetPaymentsByUser(userID uint) ([]entity.Payment, error)
	GetAllPayments() ([]entity.Payment, error)
	UpdatePaymentStatus(id uint, status string) (*entity.Payment, error)
	UpdatePayment(id uint, updated entity.Payment, userID uint) (*entity.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
	bookingRepo repository.BookingRepository
}

func NewPaymentService(pRepo repository.PaymentRepository, bRepo repository.BookingRepository) PaymentService {
	return &paymentService{
		paymentRepo: pRepo,
		bookingRepo: bRepo,
	}
}

func (s *paymentService) CreatePayment(userID uint, input entity.PaymentInput) (*entity.Payment, error) {
	if !constant.IsValidPaymentMethod(input.Method) {
		return nil, errors.New("invalid payment method")
	}

	booking, err := s.bookingRepo.FindByID(input.BookingID)
	if err != nil {
		return nil, errors.New("booking not found")
	}

	if booking.UserID != userID {
		return nil, errors.New("not your booking")
	}

	if input.ProofImage == "" {
		return nil, errors.New("proof image is required")
	}

	payment := entity.Payment{
		BookingID:   input.BookingID,
		Method:      input.Method,
		Amount:      input.Amount,
		PaymentDate: input.PaymentDate,
		ProofImage:  input.ProofImage,
		Status:      constant.PaymentWaiting,
	}

	return s.paymentRepo.Create(&payment)
}

func (s *paymentService) GetPaymentsByUser(userID uint) ([]entity.Payment, error) {
	return s.paymentRepo.FindByUserID(userID)
}

func (s *paymentService) GetAllPayments() ([]entity.Payment, error) {
	return s.paymentRepo.FindAll()
}

func (s *paymentService) UpdatePaymentStatus(id uint, status string) (*entity.Payment, error) {
	if !constant.IsValidPaymentStatus(status) {
		return nil, errors.New("invalid payment status")
	}

	payment, err := s.paymentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	payment.Status = status
	_, err = s.paymentRepo.Save(payment)
	if err != nil {
		return nil, err
	}

	if status == constant.PaymentPaid {
		booking, err := s.bookingRepo.FindByID(payment.BookingID)
		if err == nil {
			booking.Status = constant.StatusConfirmed
			_ = s.bookingRepo.Save(booking)
		}
	}

	return payment, nil
}

func (s *paymentService) UpdatePayment(id uint, updated entity.Payment, userID uint) (*entity.Payment, error) {
	payment, err := s.paymentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	booking, err := s.bookingRepo.FindByID(payment.BookingID)
	if err != nil || booking.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if payment.Status != constant.PaymentRejected {
		return nil, errors.New("only rejected payments can be updated")
	}

	payment.Method = updated.Method
	payment.Amount = updated.Amount
	payment.PaymentDate = updated.PaymentDate
	payment.ProofImage = updated.ProofImage
	payment.Status = constant.PaymentWaiting

	_, err = s.paymentRepo.Save(payment)
	return payment, err
}
