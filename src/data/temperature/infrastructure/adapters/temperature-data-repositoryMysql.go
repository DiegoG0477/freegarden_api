package adapters

import (
	database "api-order/src/Database" // Assuming shared DB connection setup
	"api-order/src/data/temperature/domain/entities"
	"database/sql"
	"log"
	// Needed for time calculations if done in Go, but prefer DB functions
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

// Create implements ports.ITemperatureData (unchanged)
func (r *TemperatureDataRepositoryMysql) Create(data entities.TemperatureData) (entities.TemperatureData, error) {
	// ... (código existente sin cambios) ...
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
		return entities.TemperatureData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for temperature data: %v", err)
		return entities.TemperatureData{}, err
	}

	data.DataID = int(id)
	return data, nil
}

// GetRecordsByKitIDAndTime implements ports.ITemperatureData <- NUEVO METODO
func (r *TemperatureDataRepositoryMysql) GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.TemperatureData, error) {
	// Use MySQL's NOW() and INTERVAL functions for efficient time filtering
	query := `
        SELECT data_id, kit_id, temperature, humidity, timestamp
        FROM temperature_humidity_data
        WHERE kit_id = ?
          AND timestamp >= NOW() - INTERVAL ? MINUTE
        ORDER BY timestamp DESC
    `
	rows, err := r.DB.Query(query, kitID, minutesAgo)
	if err != nil {
		log.Printf("Error querying temperature data by kit %d and time (%d min): %v", kitID, minutesAgo, err)
		return nil, err // Return error for DB issues
	}
	defer rows.Close()

	var records []entities.TemperatureData
	for rows.Next() {
		var record entities.TemperatureData
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&record.DataID, &record.KitID, &record.Temperature, &record.Humidity, &record.Timestamp); err != nil {
			log.Printf("Error scanning temperature data row: %v", err)
			return nil, err // Fail fast on scan error
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating temperature data rows: %v", err)
		return nil, err
	}

	// Return empty slice if no records found
	if len(records) == 0 {
		return []entities.TemperatureData{}, nil
	}

	return records, nil
}
