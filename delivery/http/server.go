package http

import (
	"net/http"
	"online-tester-service/app"
	"online-tester-service/delivery/commands"
	"online-tester-service/domain"

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

type ResponsePercentResult struct {
	PercentResult float32 `json:"percentResult"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.POST("/test", handler.AddTest)
	e.POST("/test/get/:code", handler.GetTest)
	e.PATCH("/test/:code", handler.EditTest)
	e.DELETE("/test/:code", handler.DeleteTest)

	e.POST("/test/start/:code", handler.StartTest)
	e.POST("/test/save/:code/:resultUUID", handler.FinishTest)
}

func (s *server) AddTest(c echo.Context) error {
	var cmd commands.AddEditTestCmd

	err := c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	testCode, err := s.TestsService.AddTest(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseTestCode{TestCode: *testCode})
}

func (s *server) GetTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
		cmd      commands.AuthorizeTestCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	resp, err := s.TestsService.GetTest(&testCode, &cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) EditTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
		cmd      commands.AddEditTestCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	err = s.TestsService.EditTest(&testCode, &cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) DeleteTest(c echo.Context) error {
	var (
		err      error
		testCode = c.Param("code")
		cmd      commands.AuthorizeTestCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	err = s.TestsService.DeleteTest(&testCode, &cmd)

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

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	resp, err := s.TestsService.StartTest(&testCode, &cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) FinishTest(c echo.Context) error {
	var (
		err        error
		testCode   = c.Param("code")
		resultUUID = c.Param("resultUUID")
		cmd        commands.FinishTestCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	res, err := s.TestsService.FinishTest(&testCode, &resultUUID, &cmd)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponsePercentResult{PercentResult: *res})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound, domain.ErrRecordNotFound, gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
