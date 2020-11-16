package http

import (
	"net/http"
	"online-tests/app"
	"online-tests/delivery/command"
	"online-tests/domain"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseResult struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.GET("/test/:id", handler.GetTest)
	e.POST("/test", handler.AddTest)
}

func (s *server) GetTest(c echo.Context) error {
	var err error

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func (s *server) AddTest(c echo.Context) error {
	var err error
	cmd := command.AddTestCmd{}

	err = c.Bind(&cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = s.TestsService.AddTest(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.String(http.StatusOK, "ok")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound, gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
