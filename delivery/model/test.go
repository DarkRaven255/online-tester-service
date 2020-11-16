package model

type Test struct {
	UserID         uint       `json:"user_id"`
	NumOfQuestions int        `json:"num_of_questions"`
	Questions      []Question `json:"questions"`
}
