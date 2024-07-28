package handler

type CreateCourseRequest struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Level           string `json:"level"`
	Description     string `json:"description"`
	DurationInHours int    `json:"duration_in_hours"`
}
