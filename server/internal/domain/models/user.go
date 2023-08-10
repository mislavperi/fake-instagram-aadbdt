package models

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Plan      Plan   `json:"plan"`
	Role      Role   `json:"role"`
}
