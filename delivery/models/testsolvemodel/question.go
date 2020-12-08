package testsolvemodel

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Question string    `json:"question"`
	Answers  []Answer  `json:"answers" validate:"required"`
}
