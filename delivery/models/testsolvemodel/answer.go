package testsolvemodel

import "github.com/google/uuid"

type Answer struct {
	ID      uuid.UUID `json:"id" validate:"required"`
	Answer  string    `json:"answer"`
	Checked bool      `json:"checked" validate:"required"`
}
