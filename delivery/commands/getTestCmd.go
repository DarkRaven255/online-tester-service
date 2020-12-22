package commands

import "online-tests/delivery/models/testmodel"

type GetTestCmd struct {
	Test *testmodel.Test `json:"test" validate:"required"`
}
