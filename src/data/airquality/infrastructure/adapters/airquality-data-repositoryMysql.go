package adapters

import (
	database "api-order/src/Database"
	"api-order/src/data/airquality/domain/entities" // Ajustado
	"database/sql"
	"log"
	// time package might not be strictly needed
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

// Create implements ports.IAirQualityData (unchanged)
func (r *AirQualityDataRepositoryMysql) Create(data entities.AirQualityData) (entities.AirQualityData, error) {
	// ... (código existente sin cambios) ...
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
		return entities.AirQualityData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for air quality data: %v", err)
		return entities.AirQualityData{}, err
	}

	data.DataID = int(id)
	return data, nil
}

// GetRecordsByKitIDAndTime implements ports.IAirQualityData <- NUEVO METODO
func (r *AirQualityDataRepositoryMysql) GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.AirQualityData, error) {
	// Use MySQL's NOW() and INTERVAL functions
	query := `
        SELECT data_id, kit_id, air_quality_index, timestamp
        FROM air_quality_data                       -- Ajustado nombre de tabla
        WHERE kit_id = ?
          AND timestamp >= NOW() - INTERVAL ? MINUTE
        ORDER BY timestamp DESC
    `
	rows, err := r.DB.Query(query, kitID, minutesAgo)
	if err != nil {
		log.Printf("Error querying air quality data by kit %d and time (%d min): %v", kitID, minutesAgo, err)
		return nil, err
	}
	defer rows.Close()

	var records []entities.AirQualityData // Ajustado tipo
	for rows.Next() {
		var record entities.AirQualityData // Ajustado tipo
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&record.DataID, &record.KitID, &record.AirQualityIndex, &record.Timestamp); err != nil { // Ajustado Scan
			log.Printf("Error scanning air quality data row: %v", err)
			return nil, err
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating air quality data rows: %v", err)
		return nil, err
	}

	// Return empty slice if no records found
	if len(records) == 0 {
		return []entities.AirQualityData{}, nil // Ajustado tipo
	}

	return records, nil
}
