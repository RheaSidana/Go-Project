package user

import (
	"database/sql"
	"errors"
	"fmt"
	models "go-project/models"
)

type Repository interface {
	Create(user *models.User) (*models.User, error)
	Get(id int) (*models.User, error)
	Delete(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteAll() ([]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	PatchUser(userUpdates *models.User) (*models.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) PatchUser(userUpdates *models.User) (*models.User, error) {
	if userUpdates.ID == 0 {
		return nil, errors.New("user ID required")
	}

	// Dynamically build the update query
	query := `UPDATE users SET `
	var params []interface{}
	paramCount := 1

	// Dynamically build the update query and parameters
	if userUpdates.Email != "" {
		query += fmt.Sprintf(" email = $%d,", paramCount)
		params = append(params, userUpdates.Email)
		paramCount++
	}
	if userUpdates.Name != "" {
		query += fmt.Sprintf(" name = $%d,", paramCount)
		params = append(params, userUpdates.Name)
		paramCount++
	}
	if userUpdates.Password != "" {
		query += fmt.Sprintf(" password = $%d,", paramCount)
		params = append(params, userUpdates.Password)
	}

	// Remove the trailing comma and add the WHERE clause
	if len(params) == 0 {
		return nil, errors.New("no user fields to update")
	}
	query = query[:len(query)-1] // Remove the last comma
	query += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, userUpdates.ID)

	// Execute the query
	result, err := r.db.Exec(query, params...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user found to patch")
	}

	updateUser, err := r.Get(userUpdates.ID)
	if err != nil {
		return nil, err
	}

	return updateUser, nil	
}

func (r *repository) UpdateUser(user *models.User) (*models.User, error) {
	query := `UPDATE users 
				SET name = $1, email = $2, password = $3 
				WHERE id = $4`

	result, err := r.db.Exec(
		query, 
		user.Name, 
		user.Email, 
		user.Password, 
		user.ID,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user found to update")
	}

	updateUser, err := r.Get(user.ID)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (r *repository) DeleteAll() ([]models.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	query := `DELETE FROM users`

	result, err := r.db.Exec(query)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no users to delete")
	}

	return users, nil
}

func (r *repository) GetAll() ([]models.User, error) {
	query := `SELECT id, name, email, password FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) Delete(id int) (*models.User, error) {
	user, err := r.Get(id)
	if err != nil {
		return nil, err
	}

	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil

}
func (r *repository) Get(id int) (*models.User, error) {
	query := `SELECT id, name, email, password 
				FROM users 
				WHERE id = $1`

	var user models.User

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) Create(user *models.User) (*models.User, error) {
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
