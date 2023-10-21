package main

import (
	"tech-challenge-auth/internal/channels/rest"
	"tech-challenge-auth/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	if err := rest.New().Start(); err != nil {
		logrus.Panic()
	}
}
