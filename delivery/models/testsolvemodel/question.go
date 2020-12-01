package testsolvemodel

type Question struct {
	ID       uint64   `json:"id" validate:"required"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers" validate:"required"`
}
