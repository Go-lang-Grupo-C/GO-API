package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdmin() *User {
	return &User{
		Username: "admin",
		Password: "supersenha",
	}
}

type Token struct {
	Token string `json:"token"`
}

const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"
