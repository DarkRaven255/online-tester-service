package commands

import "online-tester-service/delivery/models/testsolvemodel"

type FinishTestCmd struct {
	Test *testsolvemodel.Test `json:"test" validate:"required"`
}
