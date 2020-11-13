package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    bool   `json:"email"`
}
