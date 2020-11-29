package testsolvemodel

import "online-tests/domain/domainmodel"

type Answer struct {
	ID         uint64 `json:"id"`
	Answer     string `json:"answer"`
	QuestionID uint64 `json:"questionID"`
}

func newTestSolveAnswer(a *domainmodel.Answer, qID uint64) *Answer {
	return &Answer{
		ID:         a.ID,
		Answer:     a.Answer,
		QuestionID: qID,
	}
}

func newTestSolveAnswersArray(aArr *[]domainmodel.Answer, qID uint64) *[]Answer {
	answers := []Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newTestSolveAnswer(&a, qID))
	}

	return &answers
}
