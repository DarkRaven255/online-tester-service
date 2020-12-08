package testsolvemodel

import "github.com/google/uuid"

type Test struct {
	ID        uuid.UUID  `json:"id" validate:"required"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions" validate:"required"`
	TestTime  uint       `json:"testTime"`
}
