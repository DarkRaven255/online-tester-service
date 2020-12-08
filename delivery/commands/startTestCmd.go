package commands

import "online-tests/delivery/models/testsolvemodel"

type StartTestCmd struct {
	Result *testsolvemodel.Result `json:"result" validate:"required"`
}
