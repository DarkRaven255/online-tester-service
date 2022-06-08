package commands

import "online-tester-service/delivery/models/testsolvemodel"

type StartTestCmd struct {
	Result *testsolvemodel.Result `json:"result" validate:"required"`
}
