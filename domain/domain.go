package domain

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain/domainmodel"
)

type TestsService interface {
	AddTest(addTestCmd *command.TestCmd) (string, error)
	GetTest(testCode *string) (*response.GetTestResp, error)
	EditTest(addTestCmd *command.TestCmd, testCode *string) error
	DeleteTest(testCode *string) error
}

type TestsRepository interface {
	Create(test *domainmodel.Test) error
	GetByTestCode(testCode *string) (*domainmodel.Test, error)
	EditTestByTestCode(test *domainmodel.Test, testCode *string) error
	Delete(testCode *string) error
}
