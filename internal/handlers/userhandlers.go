package handlers

import (
	"apollo-counter/internal/controllers"
	"apollo-counter/internal/models"
	"apollo-counter/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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
	req := models.NewUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, utils.ErrInvalidJsonFormat)
		return
	}

	// validate if fields are empty
	if req.Username == "" || req.Firstname == "" || req.Lastname == "" || req.Phonenumber == "" || req.Email == "" || req.Password == "" {
		respondFailure(w, http.StatusBadRequest, utils.ErrMissingFields)
		return
	}

	err := u.userController.CreateUser(req)
	if err != nil {
		respondFailure(w, http.StatusConflict, err)
		return
	}

	respondSuccess(w, http.StatusOK, "User created successfully")
}

func (u *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if email == "" {
		respondFailure(w, http.StatusBadRequest, utils.ErrMissingFields)
		return
	}

	user, err := u.userController.GetUserByEmail(email)
	if err != nil {
		respondFailure(w, http.StatusNotFound, err)
		return
	}

	respondSuccess(w, http.StatusOK, user)
}
