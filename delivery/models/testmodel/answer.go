package testmodel

import "github.com/google/uuid"

type Answer struct {
	ID      uuid.UUID `json:"id"`
	Answer  string    `json:"answer" validate:"required"`
	Correct bool      `json:"correct" validate:"required"`
}
