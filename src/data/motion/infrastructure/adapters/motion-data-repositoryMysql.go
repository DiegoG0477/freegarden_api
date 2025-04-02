package adapters

import (
	database "api-order/src/Database"           // Assuming shared DB connection setup
	"api-order/src/data/motion/domain/entities" // Adjusted import path
	"database/sql"
	"log"
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

// Create implements ports.IMotionData
func (r *MotionDataRepositoryMysql) Create(data entities.MotionData) (entities.MotionData, error) {
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
		// Check for specific errors like foreign key violation
		return entities.MotionData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for motion data: %v", err)
		return entities.MotionData{}, err
	}

	data.DataID = int(id) // Set the generated data_id

	// Timestamp is not fetched here.

	return data, nil
}
