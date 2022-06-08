package service

import (
	"online-tester-service/delivery/commands"
	"online-tester-service/delivery/responses"
	"online-tester-service/domain"
	"online-tester-service/domain/domainmodel"
	"online-tester-service/utils"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) AddTest(cmd *commands.AddEditTestCmd) (*string, error) {
	var err error

	cmd.Test.TestCode = utils.RandomCode(8)
	cmd.Test.Password, err = utils.HashPassword(cmd.Test.Password)
	if err != nil {
		return nil, err
	}

	test, err := domainmodel.NewTestModel(cmd)
	if err != nil {
		return nil, err
	}

	err = es.testsRepo.Create(test)
	if err != nil {
		return nil, err
	}

	return &test.TestCode, nil
}

func (es *testsService) GetTest(testCode *string, cmd *commands.AuthorizeTestCmd) (*responses.TestModel, error) {
	pwd, err := es.testsRepo.GetTestPasswordHashByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(cmd.Test.Password, *pwd) {
		return nil, domain.ErrUnauthorized
	}

	result, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	return responses.NewTestModelResp(result), nil
}

func (es *testsService) EditTest(testCode *string, cmd *commands.AddEditTestCmd) error {
	pwd, err := es.testsRepo.GetTestPasswordHashByTestCode(testCode)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(cmd.Test.Password, *pwd) {
		return domain.ErrUnauthorized
	}

	test, err := domainmodel.NewEditTestModel(cmd)
	if err != nil {
		return err
	}

	err = es.testsRepo.EditTestByTestCode(test, testCode)
	if err != nil {
		return err
	}

	return nil
}

func (es *testsService) DeleteTest(testCode *string, cmd *commands.AuthorizeTestCmd) error {
	pwd, err := es.testsRepo.GetTestPasswordHashByTestCode(testCode)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(cmd.Test.Password, *pwd) {
		return domain.ErrUnauthorized
	}

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

	tm.PrepareTest()
	tm.ShuffleTest()

	rm := domainmodel.NewResultModel(cmd, &tm.ID, &tm.TestTime)
	err = es.testsRepo.AddResult(tm, rm)

	return responses.NewTestSolveModelResp(tm, &rm.ID, &rm.CreatedAt), nil
}

func (es *testsService) FinishTest(testCode *string, resultUUID *string, cmd *commands.FinishTestCmd) (*float32, error) {

	tm, err := es.testsRepo.GetByTestCode(testCode)
	if err != nil {
		return nil, err
	}

	var (
		numOfAnswers uint
		partialScore uint
		score        float32
	)

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
		return nil, err
	}

	return &finalScore, nil
}

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
