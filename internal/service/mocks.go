package service

import (
	"context"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/integration/customer"

	"github.com/stretchr/testify/mock"
)

type MockCustomerService struct {
	mock.Mock
}

func (c *MockCustomerService) Get(_ context.Context, login canonical.Login) (canonical.Login, error) {
	args := c.Called(login)
	return args.Get(0).(canonical.Login), args.Error(1)
}

type MockService struct {
	mock.Mock
	customer.CustomerService
}

func (c *MockService) Login(_ context.Context, login canonical.Login) (string, error) {
	args := c.Called(login)
	return args.Get(0).(string), args.Error(1)
}

func (c *MockService) Bypass(_ context.Context) (string, error) {
	return "fakebypasstoken", nil
}
