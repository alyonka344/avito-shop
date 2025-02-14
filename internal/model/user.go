package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Balance  int    `json:"balance"`
}
