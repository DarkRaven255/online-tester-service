package testsolvemodel

type Answer struct {
	ID      uint64 `json:"id"`
	Answer  string `json:"answer"`
	Checked bool   `json:"checked"`
}
