package adapters

import (
	database "api-order/src/Database"
	"api-order/src/data/motion/domain/entities" // Ajustado
	"database/sql"
	"log"
	// time package might not be strictly needed if using DB functions
)

type MotionDataRepositoryMysql struct {
	DB *sql.DB
}

func NewMotionDataRepositoryMysql() (*MotionDataRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &MotionDataRepositoryMysql{DB: db}, nil
}

// Create implements ports.IMotionData (unchanged)
func (r *MotionDataRepositoryMysql) Create(data entities.MotionData) (entities.MotionData, error) {
	// ... (código existente sin cambios) ...
	query := "INSERT INTO motion_data (kit_id, motion_detected) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing motion data insert statement: %v", err)
		return entities.MotionData{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.KitID, data.MotionDetected)
	if err != nil {
		log.Printf("Error executing motion data insert statement: %v", err)
		return entities.MotionData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for motion data: %v", err)
		return entities.MotionData{}, err
	}

	data.DataID = int(id)
	return data, nil
}

// GetRecordsByKitIDAndTime implements ports.IMotionData <- NUEVO METODO
func (r *MotionDataRepositoryMysql) GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.MotionData, error) {
	// Use MySQL's NOW() and INTERVAL functions
	query := `
        SELECT data_id, kit_id, motion_detected, timestamp
        FROM motion_data                             -- Ajustado nombre de tabla
        WHERE kit_id = ?
          AND timestamp >= NOW() - INTERVAL ? MINUTE
        ORDER BY timestamp DESC
    `
	rows, err := r.DB.Query(query, kitID, minutesAgo)
	if err != nil {
		log.Printf("Error querying motion data by kit %d and time (%d min): %v", kitID, minutesAgo, err)
		return nil, err
	}
	defer rows.Close()

	var records []entities.MotionData // Ajustado tipo
	for rows.Next() {
		var record entities.MotionData // Ajustado tipo
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&record.DataID, &record.KitID, &record.MotionDetected, &record.Timestamp); err != nil { // Ajustado Scan
			log.Printf("Error scanning motion data row: %v", err)
			return nil, err
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating motion data rows: %v", err)
		return nil, err
	}

	// Return empty slice if no records found
	if len(records) == 0 {
		return []entities.MotionData{}, nil // Ajustado tipo
	}

	return records, nil
}
