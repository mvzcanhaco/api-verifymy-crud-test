package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

var Users []User
