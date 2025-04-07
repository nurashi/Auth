package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Job      string `json:"job"`
	Country  string `json:"country"`
	Role     string `json:"role"`
}
