package testsolvemodel

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	ResultUUID uuid.UUID `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	FirstName  string    `json:"firstName" validate:"required email"`
	LastName   string    `json:"lastName" validate:"required"`
	Email      string    `json:"email" validate:"required"`
}
