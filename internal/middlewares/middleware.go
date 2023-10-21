package middlewares

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Logger(fx echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log := logrus.WithContext(context.Background())
		request := ctx.Request()
		log.WithFields(logrus.Fields{
			"Host":   request.Host,
			"URI":    request.RequestURI,
			"Method": request.Method,
		}).Info()

		return fx(ctx)
	}
}
