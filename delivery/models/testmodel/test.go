package testmodel

import "github.com/google/uuid"

type Test struct {
	ID                 uuid.UUID  `json:"id"`
	Title              string     `json:"title" validate:"required"`
	NumOfTestQuestions uint       `json:"numOfTestQuestions" validate:"required"`
	NumOfQuestions     uint       `json:"numOfQuestions" validate:"required"`
	Questions          []Question `json:"questions" validate:"required"`
	Randomize          bool       `json:"randomize"`
	TestTime           uint       `json:"testTime" validate:"required"`
	TestCode           string     `json:"testCode"`
	TestPassword
}

type TestPassword struct {
	Password string `json:"password,omitempty" validate:"required"`
}
