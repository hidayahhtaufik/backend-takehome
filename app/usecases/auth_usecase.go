package usecases

import (
	"errors"

	"app/dto"
	"app/models"
	"app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	Users repositories.UserRepository
}

func (u AuthUsecase) Register(in dto.RegisterRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{Name: in.Name, Email: in.Email, PasswordHash: string(hash)}
	return user, u.Users.Create(user)
}

func (u AuthUsecase) Verify(in dto.LoginRequest) (*models.User, error) {
	usr, err := u.Users.GetByEmail(in.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(in.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return usr, nil
}
