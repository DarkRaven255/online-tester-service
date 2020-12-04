package testsolvemodel

type Test struct {
	ID        uint64     `json:"id" validate:"required"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions" validate:"required"`
	TestTime  uint       `json:"testTime"`
}
