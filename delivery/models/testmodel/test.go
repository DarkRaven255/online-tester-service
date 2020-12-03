package testmodel

type Test struct {
	ID                 uint64     `json:"id"`
	Title              string     `json:"title" validate:"required"`
	NumOfTestQuestions uint       `json:"numOfTestQuestions" validate:"required"`
	TestCode           string     `json:"testCode"`
	Questions          []Question `json:"questions" validate:"required"`
	Randomize          bool       `json:"randomize" validate:"required"`
	TestTime           int        `json:"testTime" validate:"required"`
}
