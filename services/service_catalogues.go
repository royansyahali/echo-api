package services

import (
	"myapp/model"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB, err error) error {
	if err != nil {
		errRollback := tx.Rollback().Error
		if errRollback != nil {
			return errRollback
		}
		return err
	} else {
		errCommit := tx.Commit().Error
		if errCommit != nil {
			return errCommit
		}
	}
	return nil
}

type CataloguesService interface {
	Create(c echo.Context, cata model.CataloguesRequest) error
	Find(c echo.Context, id int) (model.CataloguesResponse, error)
	Delete(c echo.Context, id int) error
	List(c echo.Context) ([]model.CataloguesResponse, error)
	Update(c echo.Context, id int, cata model.CataloguesRequest) error
}

type CataloguesServiceImp struct {
	CataloguesRepository repositories.CataloguesRepository
	DB                   *gorm.DB
}

func NewCataloguesService(r repositories.CataloguesRepository, db *gorm.DB) CataloguesService {
	return &CataloguesServiceImp{
		CataloguesRepository: r,
		DB:                   db,
	}
}

func (service *CataloguesServiceImp) Create(c echo.Context, cata model.CataloguesRequest) error {
	tx := service.DB.Begin()
	catalog := model.Catalogues{
		ModelName:   cata.ModelName,
		ImageUrl:    cata.ImageUrl,
		Description: cata.Description,
	}
	err := service.CataloguesRepository.Create(c, tx, catalog)
	result := CommitOrRollback(tx, err)
	return result
}

func (service *CataloguesServiceImp) Find(c echo.Context, id int) (model.CataloguesResponse, error) {
	tx := service.DB.Begin()
	cata, err := service.CataloguesRepository.Find(c, tx, id)
	result := CommitOrRollback(tx, err)
	return cata, result
}

func (service *CataloguesServiceImp) List(c echo.Context) ([]model.CataloguesResponse, error) {
	tx := service.DB.Begin()
	catas, err := service.CataloguesRepository.List(c, tx)
	result := CommitOrRollback(tx, err)
	return catas, result
}

func (service *CataloguesServiceImp) Update(c echo.Context, id int, cata model.CataloguesRequest) error {
	tx := service.DB.Begin()
	catalog := model.Catalogues{
		ModelName:   cata.ModelName,
		ImageUrl:    cata.ImageUrl,
		Description: cata.Description,
	}
	err := service.CataloguesRepository.Update(c, tx, id, catalog)
	result := CommitOrRollback(tx, err)
	return result
}

func (service *CataloguesServiceImp) Delete(c echo.Context, id int) error {
	tx := service.DB.Begin()
	err := service.CataloguesRepository.Delete(c, tx, id)
	result := CommitOrRollback(tx, err)
	return result
}
