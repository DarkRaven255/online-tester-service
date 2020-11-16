package command

import "online-tests/delivery/model"

type AddTestCmd struct {
	Test model.Test `json:"test"`
}
