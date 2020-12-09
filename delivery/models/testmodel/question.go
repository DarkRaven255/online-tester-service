package testmodel

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question" validate:"required"`
	Answers  []Answer  `json:"answers" validate:"required"`
	Required bool      `json:"required" validate:"required"`
}
