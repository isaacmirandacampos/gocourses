package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name            string
	Slug            string
	Level           string
	Description     string
	DurationInHours int
}
