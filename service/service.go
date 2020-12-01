package service

import (
	"online-tests/delivery/commands"
	"online-tests/delivery/responses"
	"online-tests/domain"
	"online-tests/domain/domainmodel"
	"time"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) AddTest(cmd *commands.TestCmd) (string, error) {
	var (
		err  error
		test = domainmodel.NewTestModel(cmd)
	)

	err = es.testsRepo.Create(&test)
	if err != nil {
		return "", err
	}

	return test.TestCode, nil
}

func (es *testsService) GetTest(testCode *string) (*responses.TestModel, error) {
	var err error

	result, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	return responses.NewTestModelResp(result), nil
}

func (es *testsService) EditTest(cmd *commands.TestCmd, testCode *string) error {
	var (
		err  error
		test = domainmodel.NewEditTestModel(cmd)
	)

	err = es.testsRepo.EditTestByTestCode(&test, testCode)
	if err != nil {
		return err
	}

	return nil
}

func (es *testsService) DeleteTest(testCode *string) error {
	var err error

	err = es.testsRepo.Delete(testCode)
	if err != nil {
		return err
	}

	return nil
}

func (es *testsService) StartTest(testCode *string, cmd *commands.StartTestCmd) (*responses.TestSolveModel, *time.Time, *string, error) {

	tm, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, nil, nil, err
	}

	rm := domainmodel.NewResultModel(cmd, tm.ID)
	err = es.testsRepo.AddResult(tm, rm)

	return responses.NewTestSolveModelResp(tm), &rm.CreatedAt, &rm.ResultUUID, nil
}

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
