package uberzap

import (
	"testing"

	"go.uber.org/zap"
)

func TestLoggerTest(t *testing.T) {
	baseLogger, _ := zap.NewDevelopment()
	defer baseLogger.Sync()

	// Test loggerTest function with NewDevelopment logger
	loggerTest1(baseLogger)

	produLogger, _ := zap.NewProduction()
	defer produLogger.Sync()

	// Test loggerTest function with NewProduction logger
	loggerTest1(produLogger)

	sugarLogger, _ := zap.NewProduction()
	defer sugarLogger.Sync()

	// Test loggerTest function with NewProduction logger with Sugar
	loggerTest1(sugarLogger)
}

// loggerTest is the function under test
func loggerTest1(logger *zap.Logger) {
	logger.Sugar().Infof("Check order. id: %d, name: %v", 123, "Fruit")
	logger.Sugar().Infow("Check order.", "id", 123, "name", "Fruit")
}
