package domainmodel

import (
	"online-tests/delivery/commands"
	"online-tests/utils"
	"time"

	"gorm.io/gorm"
)

type Test struct {
	ID                 uint64 `gorm:"primary_key"`
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

func NewTestModel(cmd *commands.TestCmd) Test {
	tCode := utils.RandomCode(8)
	return Test{
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Password:           cmd.Test.Password,
		TestCode:           tCode,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}

func NewEditTestModel(cmd *commands.TestCmd) Test {
	return Test{
		ID:                 cmd.Test.ID,
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}
