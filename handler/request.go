package handler

import "github.com/isaacmirandacampos/gocourses/helpers"

type CreateCourseRequest struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Level           string `json:"level"`
	Description     string `json:"description"`
	DurationInHours int    `json:"duration_in_hours"`
}

func (r *CreateCourseRequest) CreateCourseValidator() error {
	if r.Name == "" {
		return helpers.ErrParamIsRequired("name", "string")
	}
	if r.Slug == "" {
		return helpers.ErrParamIsRequired("slug", "string")
	}
	if r.Level == "" {
		return helpers.ErrParamIsRequired("level", "string")
	}
	if r.Description == "" {
		return helpers.ErrParamIsRequired("description", "string")
	}
	if r.DurationInHours <= 0 {
		return helpers.ErrParamIsRequired("duration_in_hours", "int")
	}
	return nil
}