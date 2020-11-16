package repository

import (
	"online-tests/domain"
	"online-tests/domain/model"

	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) GetByID(id int64) (*model.Test, error) {
	var entry model.Test
	errs := r.db.First(&entry, id).GetErrors()
	if len(errs) > 0 {
		return &entry, errs[0]
	}
	return &entry, nil
}

func (r *repository) Delete(entry *model.Test) error {
	errs := r.db.Delete(entry).GetErrors()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (r *repository) Create(entry *model.Test) error {

	errs := r.db.Create(entry).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func NewEntryRepository(dbConn *gorm.DB) domain.TestsRepository {
	return &repository{
		db: dbConn,
	}
}
