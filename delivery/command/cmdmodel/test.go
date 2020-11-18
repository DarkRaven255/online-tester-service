package cmdmodel

type Test struct {
	Title          string     `json:"title"`
	UserID         uint       `json:"user_id"`
	NumOfQuestions int        `json:"num_of_questions"`
	Questions      []Question `json:"questions"`
}
