package testmodel

type Answer struct {
	ID      uint64 `json:"id"`
	Answer  string `json:"answer" validate:"required"`
	Correct bool   `json:"correct" validate:"required"`
}

// func newTestAnswer(a *domainmodel.Answer, qID uint64) *Answer {
// 	return &Answer{
// 		ID:     a.ID,
// 		Answer: a.Answer,
// 		// QuestionID: qID,
// 	}
// }

// func newTestAnswersArray(aArr *[]domainmodel.Answer, qID uint64) *[]Answer {
// 	answers := []Answer{}
// 	for _, a := range *aArr {
// 		answers = append(answers, *newTestAnswer(&a, qID))
// 	}

// 	return &answers
// }
