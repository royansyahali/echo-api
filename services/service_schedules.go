package services

import (
	"myapp/model"
	"myapp/repositories"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SchedulesService interface {
	Create(c echo.Context, sch model.SchedulesRequest) error
	Find(c echo.Context, id int) ([]model.SchedulesResponse, error)
	Delete(c echo.Context, id int) error
	Update(c echo.Context, id int, sch model.SchedulesRequest) error
}

type SchedulesServiceImp struct {
	SchedulesRepository repositories.SchedulesRepository
	DB                  *gorm.DB
}

func NewSchedulesService(r repositories.SchedulesRepository, db *gorm.DB) SchedulesService {
	return &SchedulesServiceImp{
		SchedulesRepository: r,
		DB:                  db,
	}
}

func (service *SchedulesServiceImp) Create(c echo.Context, sch model.SchedulesRequest) error {
	tx := service.DB.Begin()

	dt, err := time.Parse("2006-01-02", sch.ScheduleDate)
	if err != nil {
		return err
	}

	catalog := repositories.NewCataloguesRepository()
	_, err = catalog.Find(c, tx, sch.ModelId)
	if err != nil {
		return err
	}
	schedules := model.Schedules{
		ModelId:      sch.ModelId,
		ScheduleDate: dt,
	}
	err = service.SchedulesRepository.Create(c, tx, schedules)
	result := CommitOrRollback(tx, err)
	return result
}

func (service *SchedulesServiceImp) Find(c echo.Context, id int) ([]model.SchedulesResponse, error) {
	tx := service.DB.Begin()
	sch, err := service.SchedulesRepository.Find(c, tx, id)
	result := CommitOrRollback(tx, err)
	return sch, result
}

func (service *SchedulesServiceImp) Update(c echo.Context, id int, sch model.SchedulesRequest) error {
	tx := service.DB.Begin()
	dt, err := time.Parse("2006-01-02", sch.ScheduleDate)
	if err != nil {
		return err
	}
	catalog := repositories.NewCataloguesRepository()
	_, err = catalog.Find(c, tx, sch.ModelId)
	if err != nil {
		return err
	}
	schedule := model.Schedules{
		ModelId:      sch.ModelId,
		ScheduleDate: dt,
	}
	err = service.SchedulesRepository.Update(c, tx, id, schedule)
	result := CommitOrRollback(tx, err)
	return result
}

func (service *SchedulesServiceImp) Delete(c echo.Context, id int) error {
	tx := service.DB.Begin()
	err := service.SchedulesRepository.Delete(c, tx, id)
	result := CommitOrRollback(tx, err)
	return result
}
