package testsolvemodel

import (
	"online-tests/domain/domainmodel"
)

type Question struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
	TestID   uint64   `json:"testID"`
}

func newTestSolveQuestion(q *domainmodel.Question, tID uint64) *Question {
	return &Question{
		ID:       q.ID,
		Question: q.Question,
		Answers:  *newTestSolveAnswersArray(&q.Answers, q.ID),
		TestID:   tID,
	}
}

func newTestSolveQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uint64) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *newTestSolveQuestion(&q, tID))
	}

	return &questions
}
