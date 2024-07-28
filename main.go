package main

import (
	"github.com/isaacmirandacampos/gocourses/config"
	"github.com/isaacmirandacampos/gocourses/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	logger.Info("Starting the application...")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing the database: %v", err)
		return
	}
	logger.Info("Database initialized successfully")
	router.Initialize()
}
