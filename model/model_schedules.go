package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	Id           int            `json:"id"`
	ModelId      int            `json:"model_id"`
	ScheduleDate time.Time      `json:"schedule_date"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type SchedulesResponse struct {
	Id             int       `json:"id"`
	ModelId        int       `json:"model_id"`
	ScheduleDateDB time.Time `json:"-"`
	ScheduleDate   string    `json:"schedule_date"`
}
