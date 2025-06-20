package repository

import (
	"flashbook/config"
	"flashbook/entity"
	"golang.org/x/crypto/bcrypt"
	"flashbook/constant"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	SeedAdminIfNotExists() error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *entity.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SeedAdminIfNotExists() error {
	var count int64
	config.DB.Model(&entity.User{}).Where("role = ?", constant.RoleAdmin).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := entity.User{
			Name:     "admin",
			Email:    "admin1@gmail.com",
			Password: string(hashedPassword),
			Role:     constant.RoleAdmin,
		}
		return config.DB.Create(&admin).Error
	}
	return nil
}
