package repositories

import (
	"errors"
	"myapp/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CataloguesRepository interface {
	Create(c echo.Context, tx *gorm.DB, cata model.Catalogues) error
	Find(c echo.Context, tx *gorm.DB, id int) (model.CataloguesResponse, error)
	List(c echo.Context, tx *gorm.DB) ([]model.CataloguesResponse, error)
	Update(c echo.Context, tx *gorm.DB, id int, cata model.Catalogues) error
	Delete(c echo.Context, tx *gorm.DB, id int) error
}

type CataloguesRepositoryImpl struct {
}

func NewCataloguesRepository() CataloguesRepository {
	return &CataloguesRepositoryImpl{}
}

func (repository *CataloguesRepositoryImpl) Create(c echo.Context, tx *gorm.DB, cata model.Catalogues) error {
	err := tx.Create(&model.Catalogues{ModelName: cata.ModelName, ImageUrl: cata.ImageUrl, Description: cata.Description, CreatedAt: cata.CreatedAt}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *CataloguesRepositoryImpl) Find(c echo.Context, tx *gorm.DB, id int) (model.CataloguesResponse, error) {
	cata := model.CataloguesResponse{}
	sch := model.SchedulesResponse{}
	rows, err := tx.Table("catalogues").Select(`catalogues.id, catalogues.model_name, catalogues.image_url,catalogues.description, schedules.id,schedules.model_id, schedules.schedule_date`).Joins("inner join schedules on schedules.model_id = catalogues.id").Where("catalogues.id = ?", id).
		Where("catalogues.deleted_at IS NULL").Where("schedules.deleted_at IS NULL").Order("catalogues.id asc").Rows()
	if err != nil {
		return cata, err
	}
	for rows.Next() {
		rows.Scan(&cata.Id, &cata.ModelName, &cata.ImageUrl, &cata.Description, &sch.Id,
			&sch.ModelId, &sch.ScheduleDateDB,
		)
		if cata.Id != sch.ModelId {
			continue
		}
		sch.ScheduleDate = sch.ScheduleDateDB.Format("2006-01-02")
		cata.Schedules = append(cata.Schedules, sch)
	}

	if cata.Id == 0 {
		return cata, errors.New("record is not found in database")
	}
	return cata, nil
}

func (repository *CataloguesRepositoryImpl) List(c echo.Context, tx *gorm.DB) ([]model.CataloguesResponse, error) {
	cata := model.CataloguesResponse{}
	var temp = model.CataloguesResponse{}
	sch := model.SchedulesResponse{}
	catas := []model.CataloguesResponse{}

	rows, err := tx.Table("catalogues").Select(`catalogues.id,catalogues.model_name, catalogues.image_url,catalogues.description, schedules.id,schedules.model_id, schedules.schedule_date`).Joins("inner join schedules on schedules.model_id = catalogues.id").
		Where("catalogues.deleted_at IS NULL").Where("catalogues.deleted_at IS NULL").Order("catalogues.id asc").Rows()
	if err != nil {
		return catas, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&cata.Id, &cata.ModelName, &cata.ImageUrl, &cata.Description, &sch.Id,
			&sch.ModelId, &sch.ScheduleDateDB,
		)
		defer rows.Close()
		if temp.Id != 0 && temp.Id != cata.Id {
			catas = append(catas, temp)
			temp = model.CataloguesResponse{}
			cata.Schedules = nil

		}
		sch.ScheduleDate = sch.ScheduleDateDB.Format("2006-01-02")
		cata.Schedules = append(cata.Schedules, sch)
		temp = cata
	}
	catas = append(catas, temp)
	if len(catas) == 0 {
		return catas, errors.New("record is not found in database")
	}
	return catas, nil
}

func (repository *CataloguesRepositoryImpl) Update(c echo.Context, tx *gorm.DB, id int, cata model.Catalogues) error {
	catalog := model.Schedules{}
	tx.Raw("SELECT id FROM catalogues WHERE id = ?", id).Scan(&catalog.Id)
	if catalog.Id == 0 {
		return errors.New("record not found in database")
	}
	tx.Model(&model.Catalogues{}).Where("id = ?", id).Updates(cata)
	return nil
}

func (repository *CataloguesRepositoryImpl) Delete(c echo.Context, tx *gorm.DB, id int) error {
	catalog := model.Schedules{}
	tx.Raw("SELECT id FROM catalogues WHERE id = ?", id).Scan(&catalog.Id)
	if catalog.Id == 0 {
		return errors.New("record not found in database")
	}
	tx.Where("id = ?", id).Delete(&model.Catalogues{})
	tx.Where("model_id = ?", id).Delete(&model.Schedules{})
	return nil

}
