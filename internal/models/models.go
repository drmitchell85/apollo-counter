package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Phonenumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"-"` // "-" means this won't be included in JSON
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Active      bool   `json:"is_active"`
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

type GetUserByEmailResponse struct {
	Username    string `json:"username"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Phonenumber string `json:"phone_number"`
	Email       string `json:"email"`
	Active      bool   `json:"is_active"`
}
