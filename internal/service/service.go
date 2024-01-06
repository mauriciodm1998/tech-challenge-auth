package service

import (
	"context"
	"fmt"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/integration/customer"
	"tech-challenge-auth/internal/security"
	"tech-challenge-auth/internal/token"

	"github.com/sirupsen/logrus"
)

type LoginService interface {
	Login(context.Context, canonical.Login) (string, error)
	Bypass(context.Context) (string, error)
}

type loginService struct {
	customer.CustomerService
}

func NewLoginService() LoginService {
	return &loginService{
		customer.New(),
	}
}

func (u *loginService) Login(ctx context.Context, customer canonical.Login) (string, error) {
	customerBase, err := u.CustomerService.Get(ctx, customer)
	if err != nil {
		err = fmt.Errorf("error getting customer by email: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	if err = security.CheckPassword(customerBase.Password, customer.Password); err != nil {
		err = fmt.Errorf("error checking password: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	token, _ := token.GenerateToken(customerBase.Document)

	return token, nil
}

func (u *loginService) Bypass(_ context.Context) (string, error) {
	return token.GenerateToken("")
}
