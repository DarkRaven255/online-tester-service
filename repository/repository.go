package repository

import (
	"online-tests/domain"
	"online-tests/domain/model"

	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Read() ([]model.Question, error) {

	return []model.Question{}, nil
}

func (r *repository) GetByID(id int64) (model.Question, error) {

	return model.Question{}, nil
}

func (r *repository) Update(entry *model.Question) error {

	return nil
}

func (r *repository) Delete(entry *model.Question) error {
	return nil
}

func (r *repository) Create(entry *model.Question) error {
	return nil
}

func NewEntryRepository(dbConn *gorm.DB) domain.TestsRepository {
	return &repository{
		db: dbConn,
	}
}
