package repositories

import (
	"errors"
	"myapp/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SchedulesRepository interface {
	Create(c echo.Context, tx *gorm.DB, sch model.Schedules) error
	Find(c echo.Context, tx *gorm.DB, id int) ([]model.SchedulesResponse, error)
	Update(c echo.Context, tx *gorm.DB, id int, sch model.Schedules) error
	Delete(c echo.Context, tx *gorm.DB, id int) error
}

type SchedulesRepositoryImpl struct {
}

func NewSchedulesRepository() SchedulesRepository {
	return &SchedulesRepositoryImpl{}
}

func (repository *SchedulesRepositoryImpl) Create(c echo.Context, tx *gorm.DB, sch model.Schedules) error {
	err := tx.Create(&model.Schedules{ModelId: sch.ModelId, ScheduleDate: sch.ScheduleDate}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *SchedulesRepositoryImpl) Find(c echo.Context, tx *gorm.DB, id int) ([]model.SchedulesResponse, error) {
	schedule := model.SchedulesResponse{}
	sch := []model.SchedulesResponse{}
	// rows, err := tx.Raw("SELECT id, model_id, schedule_date FROM schedules WHERE model_id = ?", id).Rows()
	rows, err := tx.Table("catalogues").Select(`schedules.id,schedules.model_id, schedules.schedule_date`).Joins("right join schedules on schedules.model_id = catalogues.id").
		Where("schedules.deleted_at IS NULL").Where("catalogues.deleted_at IS NULL").Where("catalogues.id = ?", id).Order("schedules.id asc").Rows()

	if err != nil {
		return sch, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&schedule.Id, &schedule.ModelId, &schedule.ScheduleDateDB)
		schedule.ScheduleDate = schedule.ScheduleDateDB.Format("2006-01-02")
		sch = append(sch, schedule)
	}
	return sch, nil
}

func (repository *SchedulesRepositoryImpl) Update(c echo.Context, tx *gorm.DB, id int, sch model.Schedules) error {
	schedule := model.Schedules{}
	tx.Raw("SELECT id FROM schedules WHERE id = ?", id).Scan(&schedule.Id)
	if schedule.Id == 0 {
		return errors.New("record not found in database")
	}
	tx.Model(&model.Schedules{}).Where("id = ?", id).Updates(sch)
	return nil
}

func (repository *SchedulesRepositoryImpl) Delete(c echo.Context, tx *gorm.DB, id int) error {
	schedule := model.Schedules{}
	tx.Raw("SELECT id FROM schedules WHERE id = ?", id).Scan(&schedule.Id)
	if schedule.Id == 0 {
		return errors.New("record not found in database")
	}
	tx.Where("id = ?", id).Delete(&model.Schedules{})
	return nil
}
