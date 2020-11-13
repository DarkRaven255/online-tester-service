package command

import "online-tests/delivery/model"

type AddTestCmd struct {
	UserID uint       `json:"user_id"`
	Test   model.Test `json:"test"`
}
