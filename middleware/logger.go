package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := log.New()
	logFilePath := "log/"
	logLinkName := "log/last.log"
	logger.SetLevel(log.DebugLevel)
	logWriter, _ := rotate.New(
		logFilePath+"server_%Y%m%d.log",
		rotate.WithMaxAge(7*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
		rotate.WithLinkName(logLinkName),
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.DebugLevel: logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
		log.FatalLevel: logWriter,
		log.WarnLevel:  logWriter,
	}

	hook := lfshook.NewHook(writeMap, &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)
	return func(context *gin.Context) {

		startTime := time.Now()
		context.Next()
		spendTime := time.Since(startTime)
		spendTimeString := fmt.Sprintf("%d ms", int(math.Ceil(float64(spendTime.Milliseconds()))))
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		method := context.Request.Method
		path := context.Request.RequestURI
		userAgent := context.Request.UserAgent()
		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		entry := logger.WithFields(log.Fields{
			"HostName":  hostname,
			"status":    statusCode,
			"SpendTime": spendTimeString,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"dataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
