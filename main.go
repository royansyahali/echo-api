package main

import (
	"myapp/controllers"
	"myapp/database"
	"myapp/repositories"
	"myapp/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db := database.ConnectDB()
	repositoryCatalogues := repositories.NewCataloguesRepository()
	servicesCatalogues := services.NewCataloguesService(repositoryCatalogues, db)
	controllerCatalogues := controllers.NewCataloguesController(servicesCatalogues)

	repositorySchedules := repositories.NewSchedulesRepository()
	servicesSchedules := services.NewSchedulesService(repositorySchedules, db)
	controllerSchedules := controllers.NewSchedulesController(servicesSchedules)

	e.POST("/models/create", controllerCatalogues.Create)
	e.GET("/models/list/", controllerCatalogues.List)
	e.PATCH("/models/update/:id", controllerCatalogues.Update)
	e.POST("/models/schedules/create", controllerSchedules.Create)
	e.PATCH("/models/schedules/:id", controllerSchedules.Update)
	e.DELETE("/models/schedules/:id", controllerSchedules.Delete)
	e.GET("/models/schedules/:id", controllerSchedules.Find)
	e.GET("/models/:id", controllerCatalogues.Find)
	e.DELETE("/models/:id", controllerCatalogues.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
