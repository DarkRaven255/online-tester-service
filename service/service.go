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

func (es *testsService) FinishTest(testCode *string, resultUUID *string, cmd *commands.FinishTestCmd) (score float32, err error) {

	tm, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return 0.0, err
	}

	for _, questionsBase := range tm.Questions {
		for _, questionsAnswered := range cmd.Test.Questions {
			if questionsBase.ID == questionsAnswered.ID {
				partialScore := 0.0
				for _, answersBase := range questionsBase.Answers {
					for _, answersAnswered := range questionsAnswered.Answers {
						if answersBase.ID == answersAnswered.ID && answersBase.Correct == answersAnswered.Checked {
							partialScore += 1.0
						}
					}
				}
				if partialScore == 4 {
					score++
				}
			}
		}
	}

	finalScore := (score / float32(tm.NumTestOfQuestions)) * 100

	err = es.testsRepo.UpdateResult(resultUUID, &finalScore)

	if err != nil {
		return 0.0, err
	}

	return finalScore, nil
}

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
