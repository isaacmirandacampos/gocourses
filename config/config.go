package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// var err error
	// db, err = gorm.Open("sqlite3", "courses.db")
	// if err != nil {
	// 	return err
	// }
	// return nil
	return errors.New("fake error")
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
