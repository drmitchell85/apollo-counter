package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Password    string `json:"-"` // "-" means this won't be included in JSON
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Active      bool   `json:"active"`
}
type NewUserRequest struct {
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type NewUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
