package testsolvemodel

import (
	"time"
)

type Result struct {
	ResultUUID string    `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	FirstName  string    `json:"firstName" validate:"required email"`
	LastName   string    `json:"lastName" validate:"required"`
	Email      string    `json:"email" validate:"required"`
}
