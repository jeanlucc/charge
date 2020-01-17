package entities

type User struct {
	Id       string `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
