package service

import (
	"errors"
	"flashbook/constant"
	"flashbook/entity"
	"flashbook/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input entity.RegisterInput) error
	Login(input entity.LoginInput) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(input entity.RegisterInput) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     constant.RoleCustomer,
	}
	return s.userRepo.Create(&user)
}

func (s *authService) Login(input entity.LoginInput) (string, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
