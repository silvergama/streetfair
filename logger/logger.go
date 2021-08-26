package logger

import (
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() {

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash("fair.log"),
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30, // days
		Compress:   true,
	}

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = time.RFC1123Z
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)
	log.SetOutput(lumberjackLogger)
}
