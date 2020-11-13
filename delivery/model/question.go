package model

type Question struct {
	Question string   `json:"question"`
	Answer   []Answer `json:"answers"`
}
