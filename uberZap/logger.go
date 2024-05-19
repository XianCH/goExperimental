package main

import "go.uber.org/zap"

func loggerTest() {
	baseLogger, _ := zap.NewDevelopment()
	defer baseLogger.Sync()
	logger := baseLogger.Sugar()

	logger.Infof("Check order. id: %d, name: %v", 123, "Fruit")
	logger.Infow("Check order.", "id", 123, "name", "Fruit")

	produLog, _ := zap.NewProduction()
	defer produLog.Sync()
	logger = produLog.Sugar()

	logger.Infof("Check order. id: %d, name: %v", 123, "Fruit")
	logger.Infow("Check order.", "id", 123, "name", "Fruit")

	sugerLoger, _ := zap.NewProduction()
	defer sugerLoger.Sync()
	logger = sugerLoger.Sugar()

	logger.Infof("Check order. id: %d, name: %v", 123, "Fruit")
	logger.Infow("Check order.", "id", 123, "name", "Fruit")

}
