package domainmodel

import (
	"online-tests/delivery/commands"
	"online-tests/utils"
	"time"

	"gorm.io/gorm"
)

type Test struct {
	ID                 uint64         `json:"id" gorm:"primary_key"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `json:"-" sql:"index"`
	Title              string         `json:"title"`
	NumTestOfQuestions uint           `json:"numOfTestQuestions"`
	TestCode           string         `json:"testCode" qorm:"unique"`
	Questions          []Question     `json:"questions" gorm:"foreignKey:TestID"`
	Results            []Result       `json:"-" gorm:"foreignKey:TestID"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}

func NewTestModel(cmd *commands.TestCmd) Test {
	tCode := utils.RandomCode(8)
	return Test{
		Title:              cmd.Test.Title,
		NumTestOfQuestions: cmd.Test.NumOfTestQuestions,
		TestCode:           tCode,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}

func NewEditTestModel(cmd *commands.TestCmd) Test {
	return Test{
		ID:                 cmd.Test.ID,
		Title:              cmd.Test.Title,
		NumTestOfQuestions: cmd.Test.NumOfTestQuestions,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}
}
