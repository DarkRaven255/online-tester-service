package http

import (
	"net/http"
	"online-tests/app"
	"online-tests/delivery/commands"
	"online-tests/domain"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseTestCode struct {
	TestCode string `json:"testCode"`
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
	e.PATCH("/test/:code", handler.EditTest)
	e.DELETE("/test/:code", handler.DeleteTest)

	e.GET("/test/check/:code", handler.CheckIsTest)
	e.POST("/test/start/:code", handler.StartTest)
	// e.POST("/test/save/:code", handler.AddTestSolve)
}

func (s *server) AddTest(c echo.Context) error {
	var (
		err      error
		testCode string
		cmd      commands.TestCmd
	)

	err = c.Bind(&cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	testCode, err = s.TestsService.AddTest(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseTestCode{TestCode: testCode})
}

func (s *server) GetTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
	)

	resp, err := s.TestsService.GetTest(&testCode)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) EditTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
		cmd      commands.TestCmd
	)

	err = c.Bind(&cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	err = s.TestsService.EditTest(&cmd, &testCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) DeleteTest(c echo.Context) error {
	var (
		err      error
		testCode string
	)

	testCode = c.Param("code")

	err = s.TestsService.DeleteTest(&testCode)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) CheckIsTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
	)

	_, err = s.TestsService.GetTest(&testCode)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) StartTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
		cmd      commands.StartTestCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	resp, createdAt, resultUUID, err := s.TestsService.StartTest(&testCode, &cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "resultUUID"
	cookie.Value = *resultUUID
	cookie.Expires = createdAt.Add(75 * time.Minute) //TODO: add option to change time and fix for timezone to use UTC everywhere except cookies
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Domain = "web.app"

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, resp)
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
