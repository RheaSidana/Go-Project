package user

import (
	"database/sql"
	"errors"
	models "go-project/models"
)

type Repository interface {
	Create(user *models.User) (*models.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func(r *repository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users(name, email, password) 
			VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(
		query, user.Name, user.Email, user.Password,
		).Scan(&user.ID)
	
	if err != nil {
		return nil, errors.New("error inserting user into database")
	}

	return user, nil
}