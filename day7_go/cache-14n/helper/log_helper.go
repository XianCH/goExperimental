package helper

import (
	"fmt"
	"log"
	"os"
)

type LogHelper struct {
	logger  *log.Logger
	logFile *os.File
}

func NewLogHelper(filePath string) (*LogHelper, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not open log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	return &LogHelper{
		logger:  logger,
		logFile: file,
	}, nil
}

func (lh *LogHelper) Info(message string) {
	lh.logger.SetPrefix("INFO: ")
	lh.logger.Println(message)
}

func (lh *LogHelper) Warn(message string) {
	lh.logger.SetPrefix("WARN: ")
	lh.logger.Println(message)
}

func (lh *LogHelper) Error(message string) {
	lh.logger.SetPrefix("ERROR: ")
	lh.logger.Println(message)
}

func (lh *LogHelper) Debug(message string) {
	lh.logger.SetPrefix("DEBUG: ")
	lh.logger.Println(message)
}

func (lh *LogHelper) Close() error {
	if err := lh.logFile.Close(); err != nil {
		return fmt.Errorf("could not close log file: %v", err)
	}
	return nil
}
