package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New(io io.Writer) *Logger {
	log := logrus.New()
	log.SetOutput(io)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: log}
}
