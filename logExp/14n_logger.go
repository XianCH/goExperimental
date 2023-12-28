package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type LoggerLevel int

const (
	DEBUG LoggerLevel = iota
	INFO
	WARRN
	ERROR
)

var logLevelName = map[LoggerLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARRN: "WARRN",
	ERROR: "ERROR",
}

type LogEnty struct {
	Date     time.Time
	FuncName string
	FilePath string
	Level    LoggerLevel
	Message  string
}

type CustemLogger struct {
	mu       sync.Mutex
	cond     sync.Cond
	logs     []LogEnty
	logFile  *log.Logger
	logLevel LoggerLevel
}

func NewCustemLogger(logFilePath string) *CustemLogger {
	logger := log.New(logFilePath, "", log.LstdFlags|log.Lshortfile)
	return &CustemLogger{
		cond:     sync.NewCond(&sync.Mutex{}),
		logs:     make([]LogEntry, 0),
		logFile:  logger,
		logLevel: INFO, // 默认日志级别为 INFO
	}
}
