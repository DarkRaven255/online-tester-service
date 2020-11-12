package service

import (
	"online-tests/domain"
	"online-tests/domain/model"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) Read() ([]model.Question, error) {
	return es.testsRepo.Read()
}

func (es *testsService) Update(entry *model.Question) error {

	return nil
}

func (es *testsService) Delete(id int64) error {

	return nil
}

func (es *testsService) Create(entry *model.Question) error {

	return nil
}

func (es *testsService) GetTest(testID int64) error {
	return nil
}

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
