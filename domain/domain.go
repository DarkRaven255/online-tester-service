package domain

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain/domainmodel"
)

type TestsService interface {
	GetTest(testUUID string) (*response.GetTestResp, error)
	AddTest(addTestCmd *command.AddTestCmd) error
}

type TestsRepository interface {
	Create(test *domainmodel.Test) error
	Delete(test *domainmodel.Test) error
	GetByID(id string) (*domainmodel.Test, error)
}
