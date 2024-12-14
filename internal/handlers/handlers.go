package handlers

import (
	"apollo-counter/internal/controllers"
	"apollo-counter/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userController controllers.UserController
}

// here is our handler constructor
func NewUserHandler(controller controllers.UserController) *UserHandler {
	return &UserHandler{
		userController: controller,
	}
}

func (u *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	newuserrequest := models.NewUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&newuserrequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("check: %s", newuserrequest)))
}
