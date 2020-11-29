package testmodel

type Question struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question" validate:"required"`
	Answers  []Answer `json:"answers"`
	Required bool     `json:"required" validate:"required"`
}

// func newTestQuestion(q *domainmodel.Question, tID uint64) *Question {
// 	return &Question{
// 		ID:       q.ID,
// 		Question: q.Question,
// 		Answers:  *newTestAnswersArray(&q.Answers, q.ID),
// 		// TestID:   tID,
// 	}
// }

// func newTestQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uint64) *[]Question {
// 	questions := []Question{}
// 	for _, q := range *qArr {
// 		questions = append(questions, *newTestQuestion(&q, tID))
// 	}

// 	return &questions
// }
