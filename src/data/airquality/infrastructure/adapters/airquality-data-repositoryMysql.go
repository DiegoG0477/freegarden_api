package adapters

import (
	database "api-order/src/Database"               // Assuming shared DB connection setup
	"api-order/src/data/airquality/domain/entities" // Adjusted import path
	"database/sql"
	"log"
)

type AirQualityDataRepositoryMysql struct {
	DB *sql.DB
}

func NewAirQualityDataRepositoryMysql() (*AirQualityDataRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &AirQualityDataRepositoryMysql{DB: db}, nil
}

// Create implements ports.IAirQualityData
func (r *AirQualityDataRepositoryMysql) Create(data entities.AirQualityData) (entities.AirQualityData, error) {
	query := "INSERT INTO air_quality_data (kit_id, air_quality_index) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing air quality data insert statement: %v", err)
		return entities.AirQualityData{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.KitID, data.AirQualityIndex)
	if err != nil {
		log.Printf("Error executing air quality data insert statement: %v", err)
		// Check for specific errors like foreign key violation (invalid kit_id)
		return entities.AirQualityData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for air quality data: %v", err)
		return entities.AirQualityData{}, err
	}

	data.DataID = int(id) // Set the generated data_id

	// Timestamp is not fetched here, relies on DB default.

	return data, nil
}
