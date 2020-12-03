package commands

import "online-tests/delivery/models/testmodel"

type TestCmd struct {
	Test testmodel.Test `json:"test" validate:"required"`
}
