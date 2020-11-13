package model

type Test struct {
	NumOfQuestions int        `json:"num_of_questions"`
	Questions      []Question `json:"questions"`
}
