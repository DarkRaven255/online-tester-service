package command

import "online-tests/delivery/command/cmdmodel"

type AddTestCmd struct {
	Test cmdmodel.Test `json:"test"`
}
