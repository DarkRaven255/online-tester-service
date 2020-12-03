package responses

import (
	"online-tests/delivery/models/testsolvemodel"
	"online-tests/domain/domainmodel"
	"time"
)

type TestSolveModel struct {
	Test       *testsolvemodel.Test `json:"test"`
	ResultUUID *string              `json:"resultUUID"`
	FinishedAt *time.Time           `json:"finishedAt"`
}

func NewTestSolveModelResp(test *domainmodel.Test, resultUUID *string, finishedAt *time.Time) *TestSolveModel {
	return &TestSolveModel{
		Test:       newTestSolveModel(test),
		ResultUUID: resultUUID,
		FinishedAt: finishedAt,
	}
}

func newTestSolveModel(domainTest *domainmodel.Test) *testsolvemodel.Test {
	return &testsolvemodel.Test{
		ID:        domainTest.ID,
		Title:     domainTest.Title,
		Questions: *newTestSolveQuestionsArray(&domainTest.Questions, domainTest.NumOfTestQuestions, domainTest.ID),
	}
}

func newTestSolveQuestion(q *domainmodel.Question, tID uint64) *testsolvemodel.Question {
	return &testsolvemodel.Question{
		ID:       q.ID,
		Question: q.Question,
		Answers:  *newTestSolveAnswersArray(&q.Answers, q.ID),
	}
}

func newTestSolveQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uint64) *[]testsolvemodel.Question {
	questions := []testsolvemodel.Question{}
	for _, q := range *qArr {
		questions = append(questions, *newTestSolveQuestion(&q, tID))
	}

	return &questions
}

func newTestSolveAnswer(a *domainmodel.Answer, qID uint64) *testsolvemodel.Answer {
	return &testsolvemodel.Answer{
		ID:     a.ID,
		Answer: a.Answer,
	}
}

func newTestSolveAnswersArray(aArr *[]domainmodel.Answer, qID uint64) *[]testsolvemodel.Answer {
	answers := []testsolvemodel.Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newTestSolveAnswer(&a, qID))
	}

	return &answers
}
