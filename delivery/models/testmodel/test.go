package testmodel

type Test struct {
	ID                 uint64     `json:"id"`
	Title              string     `json:"title" validate:"required"`
	NumOfTestQuestions uint       `json:"numOfTestQuestions" validate:"required"`
	Questions          []Question `json:"questions"`
	Randomize          bool       `json:"randomize"`
}
