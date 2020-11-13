package model

type Answer struct {
	Answer   string `json:"question"`
	Correct  bool   `json:"correct"`
	Required bool   `json:"required"`
}
