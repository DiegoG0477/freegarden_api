package adapters

import (
	database "api-order/src/Database"
	"api-order/src/data/light/domain/entities" // Ajustado
	"database/sql"
	"log"
	// time package might not be strictly needed
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

// Create implements ports.ILightData (unchanged)
func (r *LightDataRepositoryMysql) Create(data entities.LightData) (entities.LightData, error) {
	// ... (código existente sin cambios) ...
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
		return entities.LightData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for light data: %v", err)
		return entities.LightData{}, err
	}

	data.DataID = int(id)
	return data, nil
}

// GetRecordsByKitIDAndTime implements ports.ILightData <- NUEVO METODO
func (r *LightDataRepositoryMysql) GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.LightData, error) {
	// Use MySQL's NOW() and INTERVAL functions
	query := `
        SELECT data_id, kit_id, light_level, timestamp
        FROM light_data                             -- Ajustado nombre de tabla
        WHERE kit_id = ?
          AND timestamp >= NOW() - INTERVAL ? MINUTE
        ORDER BY timestamp DESC
    `
	rows, err := r.DB.Query(query, kitID, minutesAgo)
	if err != nil {
		log.Printf("Error querying light data by kit %d and time (%d min): %v", kitID, minutesAgo, err)
		return nil, err
	}
	defer rows.Close()

	var records []entities.LightData // Ajustado tipo
	for rows.Next() {
		var record entities.LightData // Ajustado tipo
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&record.DataID, &record.KitID, &record.LightLevel, &record.Timestamp); err != nil { // Ajustado Scan
			log.Printf("Error scanning light data row: %v", err)
			return nil, err
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating light data rows: %v", err)
		return nil, err
	}

	// Return empty slice if no records found
	if len(records) == 0 {
		return []entities.LightData{}, nil // Ajustado tipo
	}

	return records, nil
}
