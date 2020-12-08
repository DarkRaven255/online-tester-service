package service

import (
	"online-tests/delivery/commands"
	"online-tests/delivery/responses"
	"online-tests/domain"
	"online-tests/domain/domainmodel"
	"online-tests/utils"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) AddTest(cmd *commands.AddEditTestCmd) (string, error) {
	var err error

	cmd.Test.TestCode = utils.RandomCode(8)
	cmd.Test.Password, err = utils.HashPassword(cmd.Test.Password)
	if err != nil {
		return "", err
	}

	test := domainmodel.NewTestModel(cmd)

	err = es.testsRepo.Create(&test)
	if err != nil {
		return "", err
	}

	return test.TestCode, nil
}

func (es *testsService) GetTest(cmd *commands.GetTestCmd) (*responses.TestModel, error) {
	pwd, err := es.testsRepo.GetTestPasswordHashByTestCode(&cmd.Test.TestCode)
	if !utils.CheckPasswordHash(cmd.Test.Password, *pwd) {
		return nil, domain.ErrUnauthorized
	}

	result, err := es.testsRepo.GetByTestCode(&cmd.Test.TestCode)
	if err != nil {
		return nil, err
	}

	return responses.NewTestModelResp(result), nil
}

func (es *testsService) EditTest(cmd *commands.AddEditTestCmd, testCode *string) error {

	pwd, err := es.testsRepo.GetTestPasswordHashByTestCode(testCode)
	if !utils.CheckPasswordHash(cmd.Test.Password, *pwd) {
		return domain.ErrUnauthorized
	}

	test := domainmodel.NewEditTestModel(cmd)

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

func (es *testsService) StartTest(testCode *string, cmd *commands.StartTestCmd) (*responses.TestSolveModel, error) {

	tm, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	rm := domainmodel.NewResultModel(cmd, tm.ID, tm.TestTime)
	err = es.testsRepo.AddResult(tm, rm)

	return responses.NewTestSolveModelResp(tm, &rm.ResultUUID, &rm.CreatedAt, &rm.FinishedAt), nil
}

func (es *testsService) FinishTest(testCode *string, resultUUID *string, cmd *commands.FinishTestCmd) (score float32, err error) {

	tm, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return 0.0, err
	}

	var numOfAnswers uint
	var partialScore uint

	for _, questionsBase := range tm.Questions {
		for _, questionsAnswered := range cmd.Test.Questions {
			if questionsBase.ID == questionsAnswered.ID {
				partialScore = 0
				numOfAnswers = uint(len(questionsBase.Answers))
				for _, answersBase := range questionsBase.Answers {
					for _, answersAnswered := range questionsAnswered.Answers {

						if answersBase.ID == answersAnswered.ID && answersBase.Correct == answersAnswered.Checked {
							partialScore += 1
						}
					}
				}
				if partialScore == numOfAnswers {
					score++
				}
			}
		}
	}

	finalScore := (float32(score) / float32(tm.NumOfTestQuestions)) * 100

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
