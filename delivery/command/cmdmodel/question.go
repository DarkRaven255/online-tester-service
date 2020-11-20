package cmdmodel

type Question struct {
	Question string   `json:"question"`
	Answer   []Answer `json:"answers"`
	Required bool     `json:"required"`
}
