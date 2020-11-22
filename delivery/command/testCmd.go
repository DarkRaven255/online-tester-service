package command

import "online-tests/delivery/command/cmdmodel"

type TestCmd struct {
	Test cmdmodel.Test `json:"test"`
}
