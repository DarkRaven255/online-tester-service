package testsolvemodel

type Test struct {
	ID        uint64     `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}
