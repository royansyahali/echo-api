package model

import (
	"time"

	"gorm.io/gorm"
)

type Catalogues struct {
	Id          int            `json:"id"`
	ModelName   string         `json:"model_name"`
	ImageUrl    string         `json:"image_url"`
	Description string         `json:"description"`
	Schedules   []Schedules    `json:"schedules" gorm:"foreignKey:ModelId"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type CataloguesResponse struct {
	Id          int                 `json:"id"`
	ModelName   string              `json:"model_name"`
	ImageUrl    string              `json:"image_url"`
	Description string              `json:"description"`
	Schedules   []SchedulesResponse `json:"schedules" gorm:"foreignKey:ModelId"`
}
