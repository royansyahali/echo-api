package model

type CataloguesRequest struct {
	ModelName   string `json:"model_name" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type SchedulesRequest struct {
	ModelId      int    `json:"model_id" validate:"required"`
	ScheduleDate string `json:"schedule_date" validate:"required"`
}
