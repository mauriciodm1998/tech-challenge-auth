package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/service"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type restTests struct {
	suite.Suite
	serviceMock *service.MockService

	server *echo.Echo
	rest   Login
}

func Test(t *testing.T) {
	suite.Run(t, new(restTests))
}

func (t *restTests) SetupSuite() {
	t.serviceMock = new(service.MockService)

	t.server = echo.New()
	t.rest = &login{
		service: t.serviceMock,
	}
}

func (t *restTests) TestStart() {
	go func() {
		err := t.rest.Start()
		t.NoError(err) // Verifique se não há erro ao iniciar o servidor
	}()

	<-time.After(100 * time.Millisecond)
}

func (t *restTests) TestLogin() {
	testCases := map[string]struct {
		input          string
		documentKey    string
		serviceError   error
		expectedResult string
		expectedError  *echo.HTTPError
	}{
		"sucess": {
			input: `
					{
						"document": "446.842.868-60",
						"email":    "mauriciodmpires1@gmail.com",
						"password": "pass123"
					}
					`,
			documentKey: "446.842.868-60",
			expectedResult: `{"token":"thisisatoken"}
`,
		},
		"invalid input": {
			expectedError: echo.NewHTTPError(http.StatusBadRequest, Response{
				Message: fmt.Errorf("invalid data").Error(),
			}),
		},
		"service error": {
			input: `
			{
				"document": "446.842.555-60",
				"email":    "mauriciodmpires1@gmail.com",
				"password": "pass123"
			}
			`,
			documentKey:  "446.842.555-60",
			serviceError: fmt.Errorf("service error"),
			expectedError: echo.NewHTTPError(http.StatusInternalServerError, Response{
				Message: fmt.Errorf("service error").Error(),
			}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func() {
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tc.input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			context := t.server.NewContext(req, rec)

			if tc.documentKey != "" {
				t.serviceMock.On("Login", mock.MatchedBy(func(l canonical.Login) bool {
					return l.Document == tc.documentKey
				})).Return("thisisatoken", tc.serviceError).Times(1)
			}

			err := t.rest.Login(context)
			if err != nil {
				t.Equal(tc.expectedError, err)
			} else {
				t.Equal(tc.expectedResult, rec.Body.String())
				t.Nil(err)
			}
		})
	}
}

func (t *restTests) TestBypass() {
	expectedResult := `{"token":"fakebypasstoken"}
`

	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := t.server.NewContext(req, rec)

	result := t.rest.Bypass(context)

	t.Nil(result)
	t.Equal(expectedResult, rec.Body.String())
}
