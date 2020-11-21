package domainmodel

import (
	"online-tests/delivery/command"
	"online-tests/utils"
	"time"
)

type Test struct {
	ID             uint64     `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"-" sql:"index"`
	Title          string     `json:"title"`
	UserID         uint64     `json:"userID"`
	NumOfQuestions uint       `json:"numOfQuestions"`
	TestCode       string     `json:"testCode" qorm:"unique"`
	Questions      []Question `json:"questions" gorm:"foreignKey:TestID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}

func NewTestModel(cmd *command.AddTestCmd) Test {
	tCode := utils.RandomCode(8)
	return Test{
		Title:          cmd.Test.Title,
		UserID:         cmd.Test.UserID,
		NumOfQuestions: cmd.Test.NumOfQuestions,
		TestCode:       tCode,
		Questions:      *NewQuestionsArray(&cmd.Test.Questions),
	}
}
