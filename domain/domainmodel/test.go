package domainmodel

import (
	"errors"
	"math/rand"
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
	TestTime           uint `gorm:"default:15"`
	Randomize          bool
	Questions          []Question `gorm:"foreignKey:TestID"`
	Results            []Result   `gorm:"foreignKey:TestID"`
}

func (Test) TableName() string {
	return "online_tester_service.tests"
}

func (test *Test) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if test.ID == nullUUID {
		test.ID = uuid.New()
	}
	return
}

func NewTestModel(cmd *commands.AddEditTestCmd) (*Test, error) {

	if err := validateCommand(cmd); err != nil {
		return nil, err
	}

	return &Test{
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Password:           cmd.Test.Password,
		TestCode:           cmd.Test.TestCode,
		Randomize:          cmd.Test.Randomize,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}, nil
}

func NewEditTestModel(cmd *commands.AddEditTestCmd) (*Test, error) {

	if err := validateCommand(cmd); err != nil {
		return nil, err
	}

	return &Test{
		ID:                 cmd.Test.ID,
		Title:              cmd.Test.Title,
		NumOfTestQuestions: cmd.Test.NumOfTestQuestions,
		NumOfQuestions:     cmd.Test.NumOfQuestions,
		TestTime:           cmd.Test.TestTime,
		Randomize:          cmd.Test.Randomize,
		Questions:          *newQuestionsArray(&cmd.Test.Questions, cmd.Test.ID),
	}, nil
}

func validateCommand(cmd *commands.AddEditTestCmd) error {
	if &cmd.Test.Title == nil {
		return errors.New("Title can not be empty!")
	}

	if &cmd.Test.NumOfTestQuestions == nil {
		return errors.New("NumOfTestQuestions can not be empty!")
	}

	if &cmd.Test.NumOfQuestions == nil {
		return errors.New("NumOfQuestions can not be empty!")
	}

	if &cmd.Test.TestTime == nil {
		return errors.New("TestTime can not be empty!")
	}

	if &cmd.Test.Password == nil {
		return errors.New("Password can not be empty!")
	}

	if &cmd.Test.Randomize == nil {
		return errors.New("Randomize can not be empty!")
	}

	if &cmd.Test.Questions == nil {
		return errors.New("Questions can not be empty!")
	}

	if cmd.Test.NumOfTestQuestions > cmd.Test.NumOfQuestions {
		return errors.New("Amount of NumOfTestQuestions can not be larger than NumOfQuestions!")
	}

	if cmd.Test.NumOfQuestions != uint(len(cmd.Test.Questions)) {
		return errors.New("NumOfQuestions does not match real amount of questions!")
	}

	return nil
}

func (test *Test) ShuffleTest() {

	if !test.Randomize {
		return
	}

	for i := 1; i < len(test.Questions); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			test.Questions[r], test.Questions[i] = test.Questions[i], test.Questions[r]
		}
	}

	for _, question := range test.Questions {
		for i := 1; i < len(question.Answers); i++ {
			r := rand.Intn(i + 1)
			if i != r {
				question.Answers[r], question.Answers[i] = question.Answers[i], question.Answers[r]
			}
		}
	}
}

func (test *Test) PrepareTest() {

	if test.NumOfQuestions == test.NumOfTestQuestions {
		return
	}

	for counter := test.NumOfQuestions; counter > test.NumOfTestQuestions; {
		random := rand.Intn(int(counter))
		if test.Questions[random].Required == false {
			test.Questions[random] = test.Questions[len(test.Questions)-1]
			test.Questions = test.Questions[:len(test.Questions)-1]
			counter--
		}
	}
}
