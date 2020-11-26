package cmdmodel

type Answer struct {
	ID      uint64 `json:"id"`
	Answer  string `json:"answer" validate:"required"`
	Correct bool   `json:"correct" validate:"required"`
}
