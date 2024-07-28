package config

import (
	"os"

	"github.com/isaacmirandacampos/gocourses/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	err := createSqliteDatabaseIfNotExist(logger)
	if err != nil {
		logger.Errorf("Error creating the database: %v", err)
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing the database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Course{})
	if err != nil {
		logger.Errorf("Sqlite auto migration error: %v", err)
		return nil, err
	}
	return db, nil
}

func createSqliteDatabaseIfNotExist(logger *Logger) error {
	dbPath := "./db/main.db"
	_, error := os.Stat(dbPath)
	if os.IsNotExist(error) {
		logger.Info("Database file does not exist, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
