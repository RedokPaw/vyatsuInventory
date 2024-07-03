package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"net/http"
	"strconv"
	_ "vyatsuAPIInventory/docs"
	"vyatsuAPIInventory/vyatsuAPI/models"
	"vyatsuAPIInventory/vyatsuAPI/repository"
)

// @title Vyatsu
// @host localhost:8080
// @version 1.0
// @BasePath /
func main() {
	e := echo.New()
	e.GET("/employees", getEmployees)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// @Summary Get List Of Employees
// @Tags get
// @Description get list
// @Produce json
// @Success 200
func getEmployees(e echo.Context) error {
	db, err := repository.New("postgres://tgbot:pass@localhost:5432/vyatsu")
	offset, err := strconv.Atoi(e.QueryParam("offset"))
	limit, err := strconv.Atoi(e.QueryParam("limit"))
	position := e.QueryParam("position")
	var employees []models.Employee
	if limit == 0 {
		limit = 10
	}
	if position == "" {
		employees, err = db.GetEmployees(offset, limit)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		employees, err = db.GetEmployeesWithFilterByPosition(offset, limit, position)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return e.JSON(http.StatusOK, employees)

}
