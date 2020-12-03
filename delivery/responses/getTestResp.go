package responses

import (
	"online-tests/delivery/models/testmodel"
	"online-tests/domain/domainmodel"
)

type TestModel struct {
	Test *testmodel.Test `json:"test"`
}

func NewTestModelResp(test *domainmodel.Test) *TestModel {
	return &TestModel{
		Test: newTestModel(test),
	}
}

func newTestModel(domainTest *domainmodel.Test) *testmodel.Test {
	return &testmodel.Test{
		ID:                 domainTest.ID,
		Title:              domainTest.Title,
		TestCode:           domainTest.TestCode,
		NumOfTestQuestions: domainTest.NumTestOfQuestions,
		Questions:          *newTestQuestionsArray(&domainTest.Questions, domainTest.NumTestOfQuestions, domainTest.ID),
	}
}

func newTestQuestion(q *domainmodel.Question, tID uint64) *testmodel.Question {
	return &testmodel.Question{
		ID:       q.ID,
		Question: q.Question,
		Required: q.Required,
		Answers:  *newTestAnswersArray(&q.Answers, q.ID),
	}
}

func newTestQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uint64) *[]testmodel.Question {
	questions := []testmodel.Question{}
	for _, q := range *qArr {
		questions = append(questions, *newTestQuestion(&q, tID))
	}

	return &questions
}

func newTestAnswer(a *domainmodel.Answer, qID uint64) *testmodel.Answer {
	return &testmodel.Answer{
		ID:      a.ID,
		Answer:  a.Answer,
		Correct: a.Correct,
	}
}

func newTestAnswersArray(aArr *[]domainmodel.Answer, qID uint64) *[]testmodel.Answer {
	answers := []testmodel.Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newTestAnswer(&a, qID))
	}

	return &answers
}
