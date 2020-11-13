package model

type Results struct {
	TestID uint    `json:"test_id"`
	UserID uint    `json:"user_id"`
	Result float32 `json:"result"`
}
