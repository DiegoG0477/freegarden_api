package adapters

import (
	database "api-order/src/Database"          // Adjust path if needed
	"api-order/src/gardendata/domain/entities" // Corrected path
	"database/sql"
	"fmt"
	"log"
)

type GardenDataRepositoryMysql struct {
	DB *sql.DB
}

func NewGardenDataRepositoryMysql() (*GardenDataRepositoryMysql, error) {
	// Assuming database.Connect() provides a configured *sql.DB
	db, err := database.Connect()
	if err != nil {
		log.Printf("Error connecting to database for GardenDataRepository: %v", err)
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	return &GardenDataRepositoryMysql{DB: db}, nil
}

// Create implements ports.IGardenData
func (r *GardenDataRepositoryMysql) Create(data entities.GardenData) (entities.GardenData, error) {
	query := `
        INSERT INTO garden_data
        (kit_id, temperature, ground_humidity, enviroment_humidity, ph_level, time)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing garden data insert statement: %v", err)
		return entities.GardenData{}, fmt.Errorf("database prepare error: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		data.KitID,
		data.Temperature,
		data.GroundHumidity,
		data.EnvironmentHumidity,
		data.PhLevel,
		data.Time,
	)
	if err != nil {
		// Log the specific error for debugging
		log.Printf("Error executing garden data insert for kit %d: %v", data.KitID, err)
		// Check for foreign key constraints, etc., if needed
		return entities.GardenData{}, fmt.Errorf("database execution error: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for garden data: %v", err)
		// This usually indicates a more serious problem
		return entities.GardenData{}, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	data.DataID = id
	// We might need to fetch the timestamp generated by the DB if we want it immediately
	// For simplicity now, we assume the controller doesn't need the exact DB timestamp right after creation

	return data, nil
}

// GetRecordsByKitIDAndTime implements ports.IGardenData
func (r *GardenDataRepositoryMysql) GetRecordsByKitIDAndTime(kitID int64, minutesAgo int) ([]entities.GardenData, error) {
	// Use MySQL's NOW() and INTERVAL functions for filtering
	query := `
        SELECT data_id, kit_id, temperature, ground_humidity, enviroment_humidity, ph_level, time, timestamp
        FROM garden_data
        WHERE kit_id = ?
          AND timestamp >= NOW() - INTERVAL ? MINUTE
        ORDER BY timestamp DESC
    `
	rows, err := r.DB.Query(query, kitID, minutesAgo)
	if err != nil {
		log.Printf("Error querying garden data by kit %d and time (%d min): %v", kitID, minutesAgo, err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	var records []entities.GardenData
	for rows.Next() {
		var record entities.GardenData
		// Ensure Scan order matches the SELECT statement columns precisely
		if err := rows.Scan(
			&record.DataID,
			&record.KitID,
			&record.Temperature,
			&record.GroundHumidity,
			&record.EnvironmentHumidity,
			&record.PhLevel,
			&record.Time,
			&record.Timestamp, // Scan the DB timestamp
		); err != nil {
			log.Printf("Error scanning garden data row: %v", err)
			// Return potentially partial results or fail entirely? Failing is safer.
			return nil, fmt.Errorf("database scan error: %w", err)
		}
		records = append(records, record)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating garden data rows: %v", err)
		return nil, fmt.Errorf("database row iteration error: %w", err)
	}

	// Return empty slice if no records were found (not an error)
	if len(records) == 0 {
		return []entities.GardenData{}, nil
	}

	return records, nil
}
