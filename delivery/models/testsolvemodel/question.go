package testsolvemodel

type Question struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}
