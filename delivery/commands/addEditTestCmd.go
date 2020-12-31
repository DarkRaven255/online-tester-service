package commands

import (
	"online-tester-service/delivery/models/testmodel"
)

type AddEditTestCmd struct {
	Test *testmodel.Test `json:"test" validate:"required"`
}
