package cmdmodel

type Results struct {
	TestID uint64  `json:"testID"`
	UserID uint64  `json:"userID"`
	Result float32 `json:"result"`
}
