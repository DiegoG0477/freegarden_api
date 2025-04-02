package adapters

import (
	database "api-order/src/Database" // Assuming shared DB connection setup
	"api-order/src/data/temperature/domain/entities"
	"database/sql"
	"log"
)

type TemperatureDataRepositoryMysql struct {
	DB *sql.DB
}

func NewTemperatureDataRepositoryMysql() (*TemperatureDataRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &TemperatureDataRepositoryMysql{DB: db}, nil
}

// Create implements ports.ITemperatureData
func (r *TemperatureDataRepositoryMysql) Create(data entities.TemperatureData) (entities.TemperatureData, error) {
	query := "INSERT INTO temperature_humidity_data (kit_id, temperature, humidity) VALUES (?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing temperature data insert statement: %v", err)
		return entities.TemperatureData{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.KitID, data.Temperature, data.Humidity)
	if err != nil {
		log.Printf("Error executing temperature data insert statement: %v", err)
		// Check for specific errors like foreign key violation (invalid kit_id)
		return entities.TemperatureData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for temperature data: %v", err)
		return entities.TemperatureData{}, err
	}

	data.DataID = int(id) // Set the generated data_id

	// Similar to the client create, we don't fetch the timestamp here.
	// The entity has the field, but it remains zero/nil in this returned object.
	// GET operations would populate it correctly from the DB.
	// Or, you could add a SELECT query here to fetch the timestamp if immediately needed.

	return data, nil
}
