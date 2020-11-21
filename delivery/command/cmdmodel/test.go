package cmdmodel

type Test struct {
	Title          string     `json:"title"`
	UserID         uint64     `json:"userID"`
	NumOfQuestions uint       `json:"numOfQuestions"`
	Questions      []Question `json:"questions"`
	TestCode       string     `json:"testCode,omitempty"`
}
