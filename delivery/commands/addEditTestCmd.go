package commands

import (
	"online-tests/delivery/models/testmodel"
)

type AddEditTestCmd struct {
	Test testmodel.Test `json:"test" validate:"required"`
}
