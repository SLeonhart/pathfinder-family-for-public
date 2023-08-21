package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type myFormatter struct {
	logrus.TextFormatter
}

var (
	logger *logrus.Logger
)

// Init is inital loging level
func Init(logLvl string) {
	logger = createLogger(logLvl)
}

func createLogger(logLvl string) *logrus.Logger {
	if logLvl == "" {
		// os.Setenv("LogLevel", "debug")
		logLvl = os.Getenv("LogLevel")
	}

	level, err := logrus.ParseLevel(logLvl)
	if err != nil {
		level = logrus.InfoLevel
	}

	return &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
		Formatter: &myFormatter{
			logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "2006-01-02 15:04:05.000",
				ForceColors:            true,
				DisableLevelTruncation: true,
			},
		},
	}
}
