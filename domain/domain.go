package domain

import (
	"online-tester-service/delivery/commands"
	"online-tester-service/delivery/responses"
	"online-tester-service/domain/domainmodel"
)

type TestsService interface {
	AddTest(addTestCmd *commands.AddEditTestCmd) (*string, error)
	GetTest(testCode *string, authorizeTestCmd *commands.AuthorizeTestCmd) (*responses.TestModel, error)
	EditTest(testCode *string, addTestCmd *commands.AddEditTestCmd) error
	DeleteTest(testCode *string, authorizeTestCmd *commands.AuthorizeTestCmd) error
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
