package adapters

import (
	database "api-order/src/Database"          // Assuming shared DB connection setup
	"api-order/src/data/light/domain/entities" // Adjusted import path
	"database/sql"
	"log"
)

type LightDataRepositoryMysql struct {
	DB *sql.DB
}

func NewLightDataRepositoryMysql() (*LightDataRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &LightDataRepositoryMysql{DB: db}, nil
}

// Create implements ports.ILightData
func (r *LightDataRepositoryMysql) Create(data entities.LightData) (entities.LightData, error) {
	query := "INSERT INTO light_data (kit_id, light_level) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing light data insert statement: %v", err)
		return entities.LightData{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.KitID, data.LightLevel)
	if err != nil {
		log.Printf("Error executing light data insert statement: %v", err)
		// Check for specific errors like foreign key violation
		return entities.LightData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for light data: %v", err)
		return entities.LightData{}, err
	}

	data.DataID = int(id) // Set the generated data_id

	// Timestamp is not fetched here.

	return data, nil
}
