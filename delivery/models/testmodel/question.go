package testmodel

type Question struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question" validate:"required"`
	Answers  []Answer `json:"answers" validate:"required"`
	Required bool     `json:"required" validate:"required"`
}
