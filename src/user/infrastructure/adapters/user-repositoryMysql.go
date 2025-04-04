package adapters

import (
	database "api-order/src/Database" // Assuming Database package is at this path
	"api-order/src/user/domain/entities"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type UserRepositoryMysql struct {
	DB *sql.DB
}

// Assuming a shared DB connection setup like in client
func NewUserRepositoryMysql() (*UserRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// It's generally better to pass the *sql.DB instance in rather than connecting here
	// e.g., func NewUserRepositoryMysql(db *sql.DB) *UserRepositoryMysql { return &UserRepositoryMysql{DB: db} }
	// But following the client example pattern:
	return &UserRepositoryMysql{DB: db}, nil
}

func (r *UserRepositoryMysql) Create(user entities.User) (entities.User, error) {
	query := "INSERT INTO users (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to prepare user insert statement: %w", err)
	}
	defer stmt.Close()

	now := time.Now()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, now)
	if err != nil {
		// Check for duplicate email error (adjust 'Error 1062' and 'users.email' if your DB differs)
		if strings.Contains(err.Error(), "Error 1062") && strings.Contains(err.Error(), "users.email") {
			return entities.User{}, err // Use the specific error
		}
		return entities.User{}, fmt.Errorf("failed to execute user insert: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to get last insert ID for user: %w", err)
	}

	user.ID = id
	user.CreatedAt = now
	// user.Password = "" // Keep password hash internal, clear before sending response in controller/usecase if needed

	return user, nil
}

func (r *UserRepositoryMysql) GetByEmail(email string) (entities.User, error) {
	query := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?"
	row := r.DB.QueryRow(query, email)

	var user entities.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, fmt.Errorf("user with email %s not found: %w", email, err)
		}
		return entities.User{}, fmt.Errorf("failed to scan user row by email: %w", err)
	}

	return user, nil
}

func (r *UserRepositoryMysql) GetById(id int64) (entities.User, error) {
	query := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, fmt.Errorf("user with id %d not found: %w", id, err)
		}
		return entities.User{}, fmt.Errorf("failed to scan user row by id: %w", err)
	}
	return user, nil
}

func (r *UserRepositoryMysql) Update(id int64, user entities.User) (entities.User, error) {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to prepare user update statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, id)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to execute user update for ID %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to get rows affected for user update ID %d: %w", id, err)
	}

	if rowsAffected == 0 {
		// This means the user ID didn't exist
		return entities.User{}, fmt.Errorf("user with id %d not found for update", id)
	}

	// Fetch the updated user data to return
	updatedUser, err := r.GetById(id)
	if err != nil {
		// This shouldn't ideally happen if update succeeded, but handle it
		return entities.User{}, fmt.Errorf("failed to fetch user data after update for ID %d: %w", id, err)
	}

	return updatedUser, nil
}

func (r *UserRepositoryMysql) CheckEmailExists(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	var exists bool
	err := r.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence for %s: %w", email, err)
	}
	return exists, nil
}
