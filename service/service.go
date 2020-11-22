package service

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain"
	"online-tests/domain/domainmodel"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) AddTest(cmd *command.TestCmd) (string, error) {
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

func (es *testsService) GetTest(testCode *string) (*response.GetTestResp, error) {
	var err error

	result, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	return response.NewGetTestResponse(result), nil
}

func (es *testsService) EditTest(cmd *command.TestCmd, testCode *string) error {
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

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
