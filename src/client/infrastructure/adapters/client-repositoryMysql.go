package adapters

import (
	database "api-order/src/Database"
	"api-order/src/client/domain/entities"
	"database/sql"
)

type ClientRepositoryMysql struct {
	DB *sql.DB
}

func NewClientRepositoryMysql() (*ClientRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return &ClientRepositoryMysql{DB: db}, nil
}

func (r *ClientRepositoryMysql) Create(client entities.Client) (entities.Client, error) {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Client{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(client.Name, client.LastName, client.Email, client.Password)

	if err != nil {
		return entities.Client{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Client{}, err
	}

	client.ID = int(id)
	client.Password = ""

	return client, nil
}

func (r *ClientRepositoryMysql) GetByEmail(email string) (entities.Client, error) {
	query := "SELECT user_id, first_name, last_name, email, password FROM users WHERE email = ?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Client{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var client entities.Client
	err = row.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil
}

func (r *ClientRepositoryMysql) GetById(id int64) (entities.Client, error) {
	query := "SELECT user_id, first_name, last_name, email, password FROM users WHERE id = ?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Client{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var client entities.Client
	err = row.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil
}
