package utils

import (
	"math/rand"
	"online-tests/domain/domainmodel"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomCode(length int) string {
	return stringWithCharset(length, charset)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ShuffleTest(test *domainmodel.Test) {

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

func PrepareTest(test *domainmodel.Test) {

	if test.NumOfQuestions == test.NumOfTestQuestions {
		return
	}

	for counter := test.NumOfQuestions; counter > test.NumOfTestQuestions; {
		random := rand.Intn(int(counter))
		if test.Questions[random].Required == false {
			remove(test, &random)
			counter--
		}
	}
}

func remove(q *domainmodel.Test, i *int) {
	q.Questions[*i] = q.Questions[len(q.Questions)-1]
	q.Questions = q.Questions[:len(q.Questions)-1]
}
