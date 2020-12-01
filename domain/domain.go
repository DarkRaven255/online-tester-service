package domain

import (
	"online-tests/delivery/commands"
	"online-tests/delivery/responses"
	"online-tests/domain/domainmodel"
	"time"
)

type TestsService interface {
	AddTest(addTestCmd *commands.TestCmd) (string, error)
	GetTest(testCode *string) (*responses.TestModel, error)
	EditTest(addTestCmd *commands.TestCmd, testCode *string) error
	DeleteTest(testCode *string) error
	StartTest(testCode *string, cmd *commands.StartTestCmd) (*responses.TestSolveModel, *time.Time, *string, error)
}

type TestsRepository interface {
	Create(test *domainmodel.Test) error
	GetByTestCode(testCode *string) (*domainmodel.Test, error)
	EditTestByTestCode(test *domainmodel.Test, testCode *string) error
	Delete(testCode *string) error
	AddResult(entry *domainmodel.Test, result *domainmodel.Result) error
}
