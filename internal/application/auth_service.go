package application

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guiddodg/go-jwt/http/request"
	"github.com/guiddodg/go-jwt/inicializer"
	"github.com/guiddodg/go-jwt/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Login(body request.AuthRequest) (string, error)
	Register(body request.AuthRequest) error
}

type authService struct{}

func (s *authService) Login(body request.AuthRequest) (string, error) {
	var user model.User
	inicializer.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	token, tokenErr := s.getToken(int(user.ID))
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}
func (s *authService) Register(body request.AuthRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash the password: %s", err)
	}

	user, userErr := model.NewUser(body.Email, string(hash))
	if userErr != nil {
		return fmt.Errorf("failed to create the user: %s", userErr)
	}

	dbUser := inicializer.DB.Create(user)
	if dbUser.Error != nil {
		return fmt.Errorf("failed to create the user: %s", dbUser.Error)
	}

	return nil
}

func (s *authService) getToken(sub int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewAuthService() AuthService {
	return &authService{}
}
