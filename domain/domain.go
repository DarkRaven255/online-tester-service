package domain

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain/model"
)

type TestsService interface {
	GetTest(testID uint) (*response.GetTestResp, error)
	AddTest(addTestCmd command.AddTestCmd) error
}

type TestsRepository interface {
	Create(entity *model.ModelEntity) error
	Delete(entity *model.ModelEntity) error
	GetByID(id int64) (*model.ModelEntity, error)
}
