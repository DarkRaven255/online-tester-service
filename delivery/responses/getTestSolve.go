package responses

import (
	"online-tester-service/delivery/models/testsolvemodel"
	"online-tester-service/domain/domainmodel"
	"time"

	"github.com/google/uuid"
)

type TestSolveModel struct {
	Test       *testsolvemodel.Test `json:"test"`
	ResultUUID *uuid.UUID           `json:"resultUUID"`
	CreatedAt  *time.Time           `json:"createdAt"`
}

func NewTestSolveModelResp(test *domainmodel.Test, resultUUID *uuid.UUID, createdAt *time.Time) *TestSolveModel {
	return &TestSolveModel{
		Test:       newTestSolveModel(test),
		ResultUUID: resultUUID,
		CreatedAt:  createdAt,
	}
}

func newTestSolveModel(domainTest *domainmodel.Test) *testsolvemodel.Test {
	return &testsolvemodel.Test{
		ID:        domainTest.ID,
		Title:     domainTest.Title,
		TestTime:  domainTest.TestTime,
		Questions: *newTestSolveQuestionsArray(&domainTest.Questions, domainTest.NumOfTestQuestions, domainTest.ID),
	}
}

func newTestSolveQuestion(q *domainmodel.Question, tID uuid.UUID) *testsolvemodel.Question {
	return &testsolvemodel.Question{
		ID:       q.ID,
		Question: q.Question,
		Answers:  *newTestSolveAnswersArray(&q.Answers, q.ID),
	}
}

func newTestSolveQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uuid.UUID) *[]testsolvemodel.Question {
	questions := []testsolvemodel.Question{}
	for _, q := range *qArr {
		questions = append(questions, *newTestSolveQuestion(&q, tID))
	}

	return &questions
}

func newTestSolveAnswer(a *domainmodel.Answer, qID uuid.UUID) *testsolvemodel.Answer {
	return &testsolvemodel.Answer{
		ID:     a.ID,
		Answer: a.Answer,
	}
}

func newTestSolveAnswersArray(aArr *[]domainmodel.Answer, qID uuid.UUID) *[]testsolvemodel.Answer {
	answers := []testsolvemodel.Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newTestSolveAnswer(&a, qID))
	}

	return &answers
}
