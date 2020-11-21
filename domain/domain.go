package domain

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain/domainmodel"
)

type TestsService interface {
	AddTest(addTestCmd *command.AddTestCmd) (string, error)
	GetTest(testCode string) (*response.GetTestResp, error)
}

type TestsRepository interface {
	Create(test *domainmodel.Test) error
	Delete(test *domainmodel.Test) error
	GetByID(testCode string) (*domainmodel.Test, error)
}
