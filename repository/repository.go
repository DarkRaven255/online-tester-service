package repository

import (
	"online-tests/domain"
	"online-tests/domain/domainmodel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(entry *domainmodel.Test) error {

	err := r.db.Create(&entry).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByTestCode(testCode *string) (*domainmodel.Test, error) {

	var entry domainmodel.Test
	err := r.db.Preload("Questions.Answers").Preload("Questions").Where("test_code = ?", testCode).First(&entry).Error

	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *repository) EditTestByTestCode(entry *domainmodel.Test, testCode *string) error {
	//TODO: Check if id can be set before save.
	for _, question := range entry.Questions {
		err := r.db.Model(&question).Association("Answers").Replace(&question.Answers)

		if err != nil {
			return err
		}
	}

	err := r.db.Model(&entry).Association("Questions").Replace(&entry.Questions)

	if err != nil {
		return err
	}

	err = r.db.Session(&gorm.Session{FullSaveAssociations: true}).Where("test_code = ?", testCode).Updates(&entry).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(testCode *string) error {

	deleteEntry, err := r.GetByTestCode(testCode)

	if err != nil {
		return err
	}

	err = r.db.Select("Answers").Delete(&deleteEntry.Questions).Error

	if err != nil {
		return err
	}

	err = r.db.Delete(&deleteEntry).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) AddResult(entry *domainmodel.Test, result *domainmodel.Result) error {

	err := r.db.Create(&result).Error

	if err != nil {
		return err
	}

	err = r.db.Save(&result).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateResult(resultUUID *string, finalScore *float32) error {

	err := r.db.Where("id = ?", resultUUID).Updates(domainmodel.Result{Result: *finalScore}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetTestPasswordHashByTestCode(testCode *string) (*string, error) {
	var entry domainmodel.Test
	err := r.db.Where("test_code = ?", testCode).First(&entry).Error

	if err != nil {
		return nil, err
	}
	return &entry.Password, nil
}

func NewEntryRepository(dbConn *gorm.DB) domain.TestsRepository {

	return &repository{
		db: dbConn,
	}
}
