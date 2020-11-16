package model

import (
	"online-tests/delivery/command"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// type Test struct {
// 	ModelEntity
// 	gorm.Model
// 	UserID         uint   `json:"user_id"`
// 	NumOfQuestions int    `json:"num_of_questions"`
// 	TestCode       string `json:"test_code"`
// }

type Test struct {
	gorm.Model
	UserID         uint       `json:"user_id"`
	NumOfQuestions int        `json:"num_of_questions"`
	TestUUID       string     `json:"test_uuid"`
	Questions      []Question `json:"questions" gorm:"foreignKey:TestID"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}

func NewTestModel(cmd *command.AddTestCmd) Test {
	tuuid := uuid.New().String()
	return Test{
		UserID:         cmd.Test.UserID,
		NumOfQuestions: cmd.Test.NumOfQuestions,
		TestUUID:       tuuid,
		Questions:      *NewQuestionsArray(&cmd.Test.Questions),
	}
}
