package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

type LogLevel int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var logLevelName = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type CustomLogger struct {
	logFile *os.File
}

func NewCustomLogger(logFilePath string) (*CustomLogger, error) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &CustomLogger{logFile: file}, nil
}

func (c *CustomLogger) logWriteStack(level LogLevel, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	// funcName := runtime.FuncForPC(pc).Name()

	message := fmt.Sprintf("[%s] [%s] [%s:%d] ", time.Now().Format("2006-01-02 15:04:05"), logLevelName[level], file, line)
	message += fmt.Sprintf(format, args...)
	log.Println(message)

	if c.logFile != nil {
		fmt.Fprintln(c.logFile, message)
	}
}

func (c *CustomLogger) Debug(format string, args ...interface{}) {
	c.logWriteStack(DEBUG, format, args...)
}

func (c *CustomLogger) Info(format string, args ...interface{}) {
	c.logWriteStack(INFO, format, args...)
}

func (c *CustomLogger) Warn(format string, args ...interface{}) {
	c.logWriteStack(WARN, format, args...)
}

func (c *CustomLogger) Error(format string, args ...interface{}) {
	c.logWriteStack(ERROR, format, args...)
}

func (c *CustomLogger) Close() {
	if c.logFile != nil {
		c.logFile.Close()
	}
}

func main() {
	logFilePath := "example.log"

	// Create a custom logger
	logger, err := NewCustomLogger(logFilePath)
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}
	defer logger.Close()
	genTime()
}

func genTime() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
	// 生成 Unix 时间戳（秒级）
	unixTimestamp := currentTime.Unix()
	fmt.Println("Unix Timestamp (seconds):", unixTimestamp)
}
