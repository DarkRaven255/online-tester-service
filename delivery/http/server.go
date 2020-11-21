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

type ResponseTestCode struct {
	TestCode string `json:"testCode"`
}

type ResponseStatus struct {
	Status string `json:"status"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.POST("/test", handler.AddTest)
	e.GET("/test/:code", handler.GetTest)
	e.DELETE("/test/:code", handler.DeleteTest)
}

func (s *server) AddTest(c echo.Context) error {
	var err error
	cmd := command.AddTestCmd{}

	err = c.Bind(&cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	testCode, err := s.TestsService.AddTest(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseTestCode{TestCode: testCode})
}

func (s *server) GetTest(c echo.Context) error {
	var (
		err      error
		testCode string
	)

	testCode = c.Param("code")

	resp, err := s.TestsService.GetTest(testCode)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) DeleteTest(c echo.Context) error {
	var (
		err      error
		testCode string
	)

	testCode = c.Param("code")

	err = s.TestsService.DeleteTest(testCode)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseStatus{Status: "ok"})
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
