package controllers

import (
	"myapp/model"
	"myapp/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string
	Message string
}

type CataloguesController interface {
	Create(c echo.Context) error
	Find(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CataloguesControllerImp struct {
	CataloguesService services.CataloguesService
}

func NewCataloguesController(service services.CataloguesService) CataloguesController {
	return &CataloguesControllerImp{
		CataloguesService: service,
	}
}

func (controller *CataloguesControllerImp) Create(c echo.Context) error {
	catalog := new(model.CataloguesRequest)
	response := Response{}
	if err := c.Bind(catalog); err != nil {
		response.Message = "Error: Model'Ms field not match"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(catalog); err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	err := controller.CataloguesService.Create(c, *catalog)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "Model created successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}

func (controller *CataloguesControllerImp) Update(c echo.Context) error {
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
	catalog := new(model.CataloguesRequest)
	if err := c.Bind(catalog); err != nil {
		response.Message = "Error: Model'Ms field not match"
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(catalog); err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	err = controller.CataloguesService.Update(c, param, *catalog)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "Model updated successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}

func (controller *CataloguesControllerImp) List(c echo.Context) error {
	response := Response{}
	catas, err := controller.CataloguesService.List(c)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusCreated, catas)
}

func (controller *CataloguesControllerImp) Find(c echo.Context) error {
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

	cata, err := controller.CataloguesService.Find(c, param)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusCreated, cata)
}

func (controller *CataloguesControllerImp) Delete(c echo.Context) error {
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

	err = controller.CataloguesService.Delete(c, param)
	if err != nil {
		response.Message = err.Error()
		response.Status = "Failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Model deleted successfully"
	response.Status = "Success"
	return c.JSON(http.StatusCreated, response)
}
