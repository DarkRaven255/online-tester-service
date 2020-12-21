package commands

import "online-tests/delivery/models/testmodel"

type AuthorizeTestCmd struct {
	Test testmodel.TestCredentials `json:"test" validate:"required"`
}
