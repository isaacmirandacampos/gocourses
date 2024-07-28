package handler

import "github.com/isaacmirandacampos/gocourses/helpers"

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
