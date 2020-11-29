package testsolvemodel

import "online-tests/domain/domainmodel"

type Test struct {
	ID        uint64     `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

func NewTestSolveModel(domainTest *domainmodel.Test) *Test {
	return &Test{
		ID:        domainTest.ID,
		Title:     domainTest.Title,
		Questions: *newTestSolveQuestionsArray(&domainTest.Questions, domainTest.NumTestOfQuestions, domainTest.ID),
	}
}
