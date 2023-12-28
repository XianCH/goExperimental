// package main
// import (
// 	"fmt"
// 	"log"
// 	"runtime"
// 	"sync"
// 	"time"
// )
//
// type LogLevel int
//
// const (
// 	DEBUG LogLevel = iota
// 	INFO
// 	WARNING
// 	ERROR
// )
//
// var logLevelNames = map[LogLevel]string{
// 	DEBUG:   "DEBUG",
// 	INFO:    "INFO",
// 	WARNING: "WARNING",
// 	ERROR:   "ERROR",
// }
//
// // LogEntry 是扩展的日志条目，包含额外的信息
// type LogEntry struct {
// 	Date     time.Time
// 	FuncName string
// 	FilePath string
// 	LogLevel LogLevel
// 	Message  string
// }
//
// // CustomLogger 是扩展的自定义日志结构体
// type CustomLogger struct {
// 	mu       sync.Mutex
// 	cond     *sync.Cond
// 	logs     []LogEntry
// 	logFile  *log.Logger
// 	logLevel LogLevel
// }
//
// // NewCustomLogger 创建一个扩展的自定义日志记录器
// func NewCustomLogger(logFilePath string) *CustomLogger {
// 	logger := log.New(logFilePath, "", log.LstdFlags|log.Lshortfile)
// 	return &CustomLogger{
// 		cond:     sync.NewCond(&sync.Mutex{}),
// 		logs:     make([]LogEntry, 0),
// 		logFile:  logger,
// 		logLevel: INFO, // 默认日志级别为 INFO
// 	}
// }
//
// // logWithStack 是一个辅助函数，用于获取调用栈信息并添加到日志消息中
// func (c *CustomLogger) logWithStack(level LogLevel, format string, args ...interface{}) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
//
// 	pc, file, line, _ := runtime.Caller(2)
// 	funcName := runtime.FuncForPC(pc).Name()
//
// 	message := fmt.Sprintf("[%s] %s [%s:%d] ", time.Now().Format("2006-01-02 15:04:05"), logLevelNames[level], funcName, line)
// 	message += fmt.Sprintf(format, args...)
//
// 	// 输出到终端
// 	log.Println(message)
//
// 	// 记录日志
// 	entry := LogEntry{
// 		Date:     time.Now(),
// 		FuncName: funcName,
// 		FilePath: file,
// 		LogLevel: level,
// 		Message:  message,
// 	}
// 	c.logs = append(c.logs, entry)
//
// 	// 发送通知，通知监控工具有新的日志输出
// 	c.cond.Broadcast()
// }
//
// // Debug 记录 DEBUG 级别的日志
// func (c *CustomLogger) Debug(format string, args ...interface{}) {
// 	if c.logLevel <= DEBUG {
// 		c.logWithStack(DEBUG, format, args...)
// 	}
// }
//
// // Info 记录 INFO 级别的日志
// func (c *CustomLogger) Info(format string, args ...interface{}) {
// 	if c.logLevel <= INFO {
// 		c.logWithStack(INFO, format, args...)
// 	}
// }
//
// // Warning 记录 WARNING 级别的日志
// func (c *CustomLogger) Warning(format string, args ...interface{}) {
// 	if c.logLevel <= WARNING {
// 		c.logWithStack(WARNING, format, args...)
// 	}
// }
//
// // Error 记录 ERROR 级别的日志
// func (c *CustomLogger) Error(format string, args ...interface{}) {
// 	if c.logLevel <= ERROR {
// 		c.logWithStack(ERROR, format, args...)
// 	}
// }
//
// // SetLogLevel 设置日志级别
// func (c *CustomLogger) SetLogLevel(level LogLevel) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.logLevel = level
// }
//
// // GetLogs 获取当前日志
// func (c *CustomLogger) GetLogs() []LogEntry {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.logs
// }
//
// // LogToFile 将日志写入指定地址
// func (c *CustomLogger) LogToFile(logFilePath string) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
//
// 	// 实现日志写入到文件的逻辑
// 	// ...
//
// 	return nil
// }
//
// // EnableLogMonitoring 启用日志监控
// func (c *CustomLogger) EnableLogMonitoring() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	go c.logMonitor()
// }
//
// // logMonitor 是监控日志输出的协程
// func (c *CustomLogger) logMonitor() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
//
// 	for {
// 		c.cond.Wait() // 等待新的日志输出
// 		logs := c.logs
// 		c.logs = make([]LogEntry, 0)
//
// 		// 输出到终端
// 		for _, entry := range logs {
// 			fmt.Printf("[Monitor] Log entry received: %s\n", entry.Message)
// 		}
// 	}
// }
//
// // DisableLogMonitoring 禁用日志监控
// func (c *CustomLogger) DisableLogMonitoring() {
// 	// 实现禁用日志监控的逻辑
// 	// ...
// }
