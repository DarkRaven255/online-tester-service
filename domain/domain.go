package domain

import (
	"online-tests/delivery/commands"
	"online-tests/delivery/responses"
	"online-tests/domain/domainmodel"
)

type TestsService interface {
	AddTest(addTestCmd *commands.AddEditTestCmd) (*string, error)
	GetTest(getTestCmd *commands.GetTestCmd) (*responses.TestModel, error)
	EditTest(addTestCmd *commands.AddEditTestCmd, testCode *string) error
	DeleteTest(testCode *string) error
	StartTest(testCode *string, cmd *commands.StartTestCmd) (*responses.TestSolveModel, error)
	FinishTest(testCode *string, resultUUID *string, cmd *commands.FinishTestCmd) (*float32, error)
}

type TestsRepository interface {
	Create(test *domainmodel.Test) error
	GetByTestCode(testCode *string) (*domainmodel.Test, error)
	EditTestByTestCode(test *domainmodel.Test, testCode *string) error
	Delete(testCode *string) error
	AddResult(entry *domainmodel.Test, result *domainmodel.Result) error
	UpdateResult(resultUUID *string, finalScore *float32) error
	GetTestPasswordHashByTestCode(testCode *string) (*string, error)
}
