package domainmodel

import (
	"online-tests/delivery/commands"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Test struct {
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `sql:"index"`
	Title              string
	Password           string
	TestCode           string `qorm:"unique"`
	NumOfTestQuestions uint
	NumOfQuestions     uint
	TestTime           uint       `gorm:"default:20"`
	Questions          []Question `gorm:"foreignKey:TestID"`
	Results            []Result   `gorm:"foreignKey:TestID"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}

func (test *Test) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if test.ID == nullUUID {
		test.ID = uuid.New()
	}
	return
}

func NewTestModel(cmd *commands.AddEditTestCmd) Test {
	return Test{
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Password:           cmd.Test.Password,
		TestCode:           cmd.Test.TestCode,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}

func NewEditTestModel(cmd *commands.AddEditTestCmd) Test {
	return Test{
		ID:                 cmd.Test.ID,
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}
