package http

import (
	"net/http"
	"online-tests/app"
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

	return c.String(http.StatusOK, "ok")
}

func (s *server) AddTest(c echo.Context) error {

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
