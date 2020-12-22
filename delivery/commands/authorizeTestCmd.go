package commands

import "online-tests/delivery/models/testmodel"

type AuthorizeTestCmd struct {
	Test testmodel.TestPassword `json:"test" validate:"required"`
}
