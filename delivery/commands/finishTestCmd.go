package commands

import "online-tests/delivery/models/testsolvemodel"

type FinishTestCmd struct {
	Test *testsolvemodel.Test `json:"test"`
}
