package model

type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Gmail    string `json:"gmail"`
	Password string `json:"password"`
}
