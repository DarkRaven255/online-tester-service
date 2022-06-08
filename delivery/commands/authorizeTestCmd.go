package commands

import "online-tester-service/delivery/models/testmodel"

type AuthorizeTestCmd struct {
	Test *testmodel.TestPassword `json:"test" validate:"required"`
}
