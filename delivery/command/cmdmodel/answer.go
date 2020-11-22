package cmdmodel

type Answer struct {
	ID      uint64 `json:"id"`
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}
