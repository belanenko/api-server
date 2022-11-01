package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logrus.New(),
	}
}

func (l *Logger) SetLogLevel(level string) error {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		l.WithFields(logrus.Fields{
			"log level": level,
		}).Error(err)
		return err
	}

	l.SetLevel(logrusLevel)
	l.WithFields(logrus.Fields{
		"log level": level,
	}).Info("set log level")
	return nil
}

func (l *Logger) SetProductionFormatter() {
	l.SetFormatter(&logrus.JSONFormatter{})
}
