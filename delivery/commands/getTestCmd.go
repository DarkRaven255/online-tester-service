package commands

import "online-tester-service/delivery/models/testmodel"

type GetTestCmd struct {
	Test *testmodel.Test `json:"test" validate:"required"`
}
