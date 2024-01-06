package service

import (
	"context"
	"fmt"
	"tech-challenge-auth/internal/canonical"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type tests struct {
	suite.Suite
	customerMock *MockCustomerService

	svc LoginService
}

func Test(t *testing.T) {
	suite.Run(t, new(tests))
}

func (s *tests) SetupSuite() {
	s.customerMock = new(MockCustomerService)

	s.svc = &loginService{
		CustomerService: s.customerMock,
	}
}

func (t *tests) TestLogin() {
	testCases := map[string]struct {
		inputLogin       canonical.Login
		customerResponse canonical.Login
		expectedError    error
		returnedError    error
	}{
		"sucessCase": {
			inputLogin: canonical.Login{
				Document: "446.842.868-60",
				Email:    "mauriciodmpires1@gmail.com",
				Password: "pass123",
			},
			customerResponse: canonical.Login{
				Document: "446.842.868-60",
				Email:    "mauriciodmpires1@gmail.com",
				Password: "$2a$10$vUqsinANCO86qwXn.MeP1ORE.DJ5bwCL/Zf33iL.S/pnxrtHSqYvG",
			},
		},
		"integration error": {
			inputLogin: canonical.Login{
				Document: "446.842.868-12",
				Email:    "mauriciodmpires1@gmail.com",
				Password: "pass123",
			},
			customerResponse: canonical.Login{
				Document: "446.842.868-12",
			},
			expectedError: fmt.Errorf("error getting customer by email: %w", fmt.Errorf("error generting token")),
			returnedError: fmt.Errorf("error generting token"),
		},
		"wrong password": {
			inputLogin: canonical.Login{
				Document: "446.842.868-56",
				Email:    "mauriciodmpires1@gmail.com",
				Password: "pass133",
			},
			customerResponse: canonical.Login{
				Document: "446.842.868-56",
				Email:    "mauriciodmpires1@gmail.com",
				Password: "$2a$10$vUqsinANCO86qwXn.MeP1ORE.DJ5bwCL/Zf33iL.S/pnxrtHSqYvG",
			},
			expectedError: fmt.Errorf("error checking password: %w", fmt.Errorf("crypto/bcrypt: hashedPassword is not the hash of the given password")),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func() {
			t.customerMock.On("Get", mock.MatchedBy(func(l canonical.Login) bool {
				return l.Document == tc.inputLogin.Document
			})).Return(tc.customerResponse, tc.returnedError).Times(1)

			_, err := t.svc.Login(context.Background(), tc.inputLogin)

			t.Equal(err, tc.expectedError)
		})
	}
}

func (t *tests) TestBypass() {
	result, _ := t.svc.Bypass(context.Background())

	t.NotNil(result)
}
