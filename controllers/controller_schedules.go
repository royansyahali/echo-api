package controllers

import (
	"myapp/model"
	"myapp/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SchedulesController interface {
	Create(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type SchedulesControllerImp struct {
	SchedulesService services.SchedulesService
}

func NewSchedulesController(service services.SchedulesService) SchedulesController {
	return &SchedulesControllerImp{
		SchedulesService: service,
	}
}

func (controller *SchedulesControllerImp) Create(c echo.Context) error {
	sch := new(model.SchedulesRequest)
	response := Response{}
	if err := c.Bind(sch); err != nil {
		response.Message = "Error: Schedules's field is required"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(sch); err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	err := controller.SchedulesService.Create(c, *sch)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "Schedule created successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}

func (controller *SchedulesControllerImp) Update(c echo.Context) error {
	response := Response{}
	id := c.Param("id")
	if id == "" {
		response.Message = "Error: Query parameter is required"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	param, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "Error: Query parameter must be number"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	sch := new(model.SchedulesRequest)
	if err := c.Bind(sch); err != nil {
		response.Message = "Error: Schedule's field is required"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(sch); err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	err = controller.SchedulesService.Update(c, param, *sch)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "Schedule updated successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}

func (controller *SchedulesControllerImp) Find(c echo.Context) error {
	response := Response{}
	id := c.Param("id")
	if id == "" {
		response.Message = "Error: Query parameter is required"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	param, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "Error: Query parameter must be number"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	cata, err := controller.SchedulesService.Find(c, param)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(cata) == 0 {
		response.Message = "Schedule not in database"
		response.Status = "Success"
		return c.JSON(http.StatusCreated, response)
	}
	return c.JSON(http.StatusCreated, cata)
}

func (controller *SchedulesControllerImp) Delete(c echo.Context) error {
	response := Response{}
	id := c.Param("id")
	if id == "" {
		response.Message = "Error: Query parameter is required"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	param, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "Error: Query parameter must be number"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	err = controller.SchedulesService.Delete(c, param)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Schedule deleted successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}
