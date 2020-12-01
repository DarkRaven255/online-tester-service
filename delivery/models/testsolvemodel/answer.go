package testsolvemodel

type Answer struct {
	ID      uint64 `json:"id" validate:"required"`
	Answer  string `json:"answer"`
	Checked bool   `json:"checked" validate:"required"`
}
