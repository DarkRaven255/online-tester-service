package testsolvemodel

import (
	"time"
)

type Result struct {
	ResultUUID string    `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
}
