package repository

import (
	"online-tests/domain"
	"online-tests/domain/domainmodel"

	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) GetByID(testUUID string) (*domainmodel.Test, error) {
	var entry domainmodel.Test
	var errs []error

	errs = r.db.Preload("Questions").Preload("Questions.Answers").Where("test_uuid = ?", testUUID).First(&entry).GetErrors()

	if len(errs) > 0 {
		return &entry, errs[0]
	}
	return &entry, nil
}

func (r *repository) Delete(entry *domainmodel.Test) error {
	errs := r.db.Delete(entry).GetErrors()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (r *repository) Create(entry *domainmodel.Test) error {

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
