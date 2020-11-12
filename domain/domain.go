package domain

import "online-tests/domain/model"

type TestsService interface {
	GetTest(testID int64) error
	Create(question *model.Question) error
	Read() ([]model.Question, error)
	Delete(id int64) error
}

type TestsRepository interface {
	Create(question *model.Question) error
	Read() ([]model.Question, error)
	Update(question *model.Question) error
	Delete(question *model.Question) error
	GetByID(id int64) (model.Question, error)
}
