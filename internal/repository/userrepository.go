package repository

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/utils"
	"database/sql"

	"github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.GetUserByEmailResponse, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user models.User) error {
	q := `
		INSERT INTO users (
			first_name, 
			last_name, 
			phone_number, 
			email, 
			password_hash, 
			username
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.db.Exec(
		q,
		user.Firstname,
		user.Lastname,
		user.Phonenumber,
		user.Email,
		user.Password,
		user.Username,
	)
	if err != nil {

		pErr, _ := err.(*pq.Error)
		if pErr.Code == "23505" && pErr.Constraint == "users_email_key" {
			return utils.ErrDuplicateEmail
		}

		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(email string) (models.GetUserByEmailResponse, error) {
	user := models.GetUserByEmailResponse{}
	q := (`SELECT username, first_name, last_name, phone_number, email, is_active FROM users WHERE email = $1`)
	err := r.db.QueryRow(q, email).Scan(&user.Username, &user.Firstname, &user.Lastname, &user.Phonenumber, &user.Email, &user.Active)

	if err != nil {
		return user, utils.ErrUserNotFound
	}

	return user, nil
}
