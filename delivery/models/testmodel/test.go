package testmodel

type Test struct {
	ID                 uint64     `json:"id"`
	Title              string     `json:"title" validate:"required"`
	NumOfTestQuestions uint       `json:"numOfTestQuestions" validate:"required"`
	Questions          []Question `json:"questions"`
	Randomize          bool       `json:"randomize"`
}

// func newTestModel(domainTest *domainmodel.Test) *Test {
// 	return &Test{
// 		ID:        domainTest.ID,
// 		Title:     domainTest.Title,
// 		Questions: *newTestQuestionsArray(&domainTest.Questions, domainTest.NumTestOfQuestions, domainTest.ID),
// 	}
// }
