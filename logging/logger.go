package logging

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.StandardLogger()

func init() {
	if config.LogLevel == "" {
		config.LogLevel = "debug"
	}

	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}
