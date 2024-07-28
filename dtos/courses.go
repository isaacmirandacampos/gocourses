package dtos

import "time"

type CourseResponse struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Level           string    `json:"level"`
	Description     string    `json:"description"`
	DurationInHours int       `json:"duration_in_hours"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
